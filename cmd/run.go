package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uvite/u8/core"
	"github.com/uvite/u8/core/local"
	"github.com/uvite/u8/errext"
	"github.com/uvite/u8/errext/exitcodes"
	"github.com/uvite/u8/js/common"
	"github.com/uvite/u8/lib/consts"
	"io"
	"net/http"
	"os"
	"runtime"
)

// cmdRun handles the `k6 run` sub-command
type cmdRun struct {
	gs *globalState
}

// TODO: split apart some more
//
//nolint:funlen,gocognit,gocyclo,cyclop
func (c *cmdRun) run(cmd *cobra.Command, args []string) error {
	//printBanner(c.gs)

	test, err := loadAndConfigureTest(c.gs, cmd, args, getConfig)
	if err != nil {
		return err
	}

	// Write the full consolidated *and derived* options back to the Runner.
	conf := test.derivedConfig
	testRunState, err := test.buildTestRunState(conf.Options)
	if err != nil {
		return err
	}

	// We prepare a bunch of contexts:
	//  - The runCtx is cancelled as soon as the Engine's run() lambda finishes,
	//    and can trigger things like the usage report and end of test summary.
	//    Crucially, metrics processing by the Engine will still work after this
	//    context is cancelled!
	//  - The lingerCtx is cancelled by Ctrl+C, and is used to wait for that
	//    event when k6 was ran with the --linger option.
	//  - The globalCtx is cancelled only after we're completely done with the
	//    test execution and any --linger has been cleared, so that the Engine
	//    can start winding down its metrics processing.
	globalCtx, globalCancel := context.WithCancel(c.gs.ctx)
	defer globalCancel()
	lingerCtx, lingerCancel := context.WithCancel(globalCtx)
	defer lingerCancel()
	runCtx, runCancel := context.WithCancel(lingerCtx)
	defer runCancel()

	logger := testRunState.Logger
	// Create a local execution scheduler wrapping the runner.
	logger.Debug("Initializing the execution scheduler...")
	execScheduler, err := local.NewExecutionScheduler(testRunState)
	if err != nil {
		return err
	}

	// This is manually triggered after the Engine's Run() has completed,
	// and things like a single Ctrl+C don't affect it. We use it to make
	// sure that the progressbars finish updating with the latest execution
	// state one last time, after the test run has finished.
	//progressCtx, progressCancel := context.WithCancel(globalCtx)
	//defer progressCancel()
	//initBar := execScheduler.GetInitProgressBar()
	//progressBarWG := &sync.WaitGroup{}
	//progressBarWG.Add(1)
	//go func() {
	//	//pbs := []*pb.ProgressBar{execScheduler.GetInitProgressBar()}
	//	//for _, s := range execScheduler.GetExecutors() {
	//	//	pbs = append(pbs, s.GetProgress())
	//	//}
	//	//showProgress(progressCtx, c.gs, pbs, logger)
	//	progressBarWG.Done()
	//}()

	// Create all outputs.
	executionPlan := execScheduler.GetExecutionPlan()
	outputs, err := createOutputs(c.gs, test, executionPlan)
	if err != nil {
		return err
	}

	// TODO: create a MetricsEngine here and add its ingester to the list of
	// outputs (unless both NoThresholds and NoSummary were enabled)

	// TODO: remove this completely
	// Create the engine.
	//initBar.Modify(pb.WithConstProgress(0, "Init engine"))

	engine, err := core.NewEngine(testRunState, execScheduler, outputs)
	if err != nil {
		return err
	}

	engineRun, err := engine.Init(globalCtx, runCtx)

	var interrupt error
	err = engineRun()
	if err != nil {
		err = common.UnwrapGojaInterruptedError(err)
		if errext.IsInterruptError(err) {
			// Don't return here since we need to work with --linger,
			// show the end-of-test summary and exit cleanly.
			interrupt = err
		}
		if !conf.Linger.Bool && interrupt == nil {
			return errext.WithExitCodeIfNone(err, exitcodes.GenericEngine)
		}
	}
	runCancel()
	logger.Debug("Engine run terminated cleanly")

	globalCancel() // signal the Engine that it should wind down
	logger.Debug("Waiting for engine processes to finish...")
	//engineWait()

	return nil
}

func (c *cmdRun) flagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	flags.SortFlags = false
	flags.AddFlagSet(optionFlagSet())
	flags.AddFlagSet(runtimeOptionFlagSet(true))
	flags.AddFlagSet(configFlagSet())
	return flags
}

func getCmdRun(gs *globalState) *cobra.Command {
	c := &cmdRun{
		gs: gs,
	}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Start a load test",
		Long: `Start a load test.

This also exposes a REST API to interact with it. Various k6 subcommands offer
a commandline interface for interacting with it.`,
		Example: `
  # Run a single VU, once.
  k6 run script.js

  # Run a single VU, 10 times.
  k6 run -i 10 script.js

  # Run 5 VUs, splitting 10 iterations between them.
  k6 run -u 5 -i 10 script.js

  # Run 5 VUs for 10s.
  k6 run -u 5 -d 10s script.js

  # Ramp VUs from 0 to 100 over 10s, stay there for 60s, then 10s down to 0.
  k6 run -u 0 -s 10s:100 -s 60s -s 10s:0

  # Send metrics to an influxdb server
  k6 run -o influxdb=http://1.2.3.4:8086/k6`[1:],
		Args: exactArgsWithMsg(1, "arg should either be \"-\", if reading script from stdin, or a path to a script file"),
		RunE: c.run,
	}

	runCmd.Flags().SortFlags = false
	runCmd.Flags().AddFlagSet(c.flagSet())

	return runCmd
}

func reportUsage(execScheduler *local.ExecutionScheduler) error {
	execState := execScheduler.GetState()
	executorConfigs := execScheduler.GetExecutorConfigs()

	executors := make(map[string]int)
	for _, ec := range executorConfigs {
		executors[ec.GetType()]++
	}

	body, err := json.Marshal(map[string]interface{}{
		"k6_version": consts.Version,
		"executors":  executors,
		"vus_max":    execState.GetInitializedVUsCount(),
		"iterations": execState.GetFullIterationCount(),
		"duration":   execState.GetCurrentTestRunDuration().String(),
		"goos":       runtime.GOOS,
		"goarch":     runtime.GOARCH,
	})
	if err != nil {
		return err
	}
	res, err := http.Post("https://reports.k6.io/", "application/json", bytes.NewBuffer(body)) //nolint:noctx
	defer func() {
		if err == nil {
			_ = res.Body.Close()
		}
	}()

	return err
}

func handleSummaryResult(fs afero.Fs, stdOut, stdErr io.Writer, result map[string]io.Reader) error {
	var errs []error

	getWriter := func(path string) (io.Writer, error) {
		switch path {
		case "stdout":
			return stdOut, nil
		case "stderr":
			return stdErr, nil
		default:
			return fs.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		}
	}

	for path, value := range result {
		if writer, err := getWriter(path); err != nil {
			errs = append(errs, fmt.Errorf("could not open '%s': %w", path, err))
		} else if n, err := io.Copy(writer, value); err != nil {
			errs = append(errs, fmt.Errorf("error saving summary to '%s' after %d bytes: %w", path, n, err))
		}
	}

	return consolidateErrorMessage(errs, "Could not save some summary information:")
}
