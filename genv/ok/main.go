package main

import (
	"context"
	"fmt"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/uvite/u8/js"
	"github.com/uvite/u8/lib"
	"github.com/uvite/u8/lib/consts"
	"github.com/uvite/u8/lib/testutils"
	"github.com/uvite/u8/loader"
	"github.com/uvite/u8/metrics"
	"gopkg.in/guregu/null.v3"
	"net/url"
	"testing"
	"time"
)

func getTestPreInitState(tb testing.TB, logger *logrus.Logger, rtOpts *lib.RuntimeOptions) *lib.TestPreInitState {
	if logger == nil {
		logger = testutils.NewLogger(tb)
	}
	if rtOpts == nil {
		rtOpts = &lib.RuntimeOptions{}
	}
	reg := metrics.NewRegistry()
	return &lib.TestPreInitState{
		Logger:         logger,
		RuntimeOptions: *rtOpts,
		Registry:       reg,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(reg),
	}
}

func getSimpleBundle(tb testing.TB, filename, data string, opts ...interface{}) (*js.Bundle, error) {
	fs := afero.NewMemMapFs()
	var rtOpts *lib.RuntimeOptions
	var logger *logrus.Logger
	for _, o := range opts {
		switch opt := o.(type) {
		case afero.Fs:
			fs = opt
		case lib.RuntimeOptions:
			rtOpts = &opt
		case *logrus.Logger:
			logger = opt
		default:
			tb.Fatalf("unknown test option %q", opt)
		}
	}

	return js.NewBundle(
		getTestPreInitState(tb, logger, rtOpts),
		&loader.SourceData{
			URL:  &url.URL{Path: filename, Scheme: "file"},
			Data: []byte(data),
		},
		map[string]afero.Fs{"file": fs, "https": afero.NewMemMapFs()},
	)
}
func getSimpleRunner1(tb testing.TB, filename, data string, opts ...interface{}) (*js.Runner, error) {
	var (
		rtOpts      = lib.RuntimeOptions{CompatibilityMode: null.NewString("base", true)}
		logger      = testutils.NewLogger(tb)
		fsResolvers = map[string]afero.Fs{"file": afero.NewMemMapFs(), "https": afero.NewMemMapFs()}
	)
	for _, o := range opts {
		switch opt := o.(type) {
		case afero.Fs:
			fsResolvers["file"] = opt
		case map[string]afero.Fs:
			fsResolvers = opt
		case lib.RuntimeOptions:
			rtOpts = opt
		case *logrus.Logger:
			logger = opt
		default:
			tb.Fatalf("unknown test option %q", opt)
		}
	}
	registry := metrics.NewRegistry()
	builtinMetrics := metrics.RegisterBuiltinMetrics(registry)
	return js.New(
		&lib.TestPreInitState{
			Logger:         logger,
			RuntimeOptions: rtOpts,
			BuiltinMetrics: builtinMetrics,
			Registry:       registry,
		},
		&loader.SourceData{
			URL:  &url.URL{Path: filename, Scheme: "file"},
			Data: []byte(data),
		},
		fsResolvers,
	)
}

func getSimpleRunner(tb testing.TB, filename, data string, opts ...interface{}) (*js.Runner, error) {
	var (
		rtOpts      = lib.RuntimeOptions{CompatibilityMode: null.NewString("base", true)}
		logger      = testutils.NewLogger(tb)
		fsResolvers = map[string]afero.Fs{"file": afero.NewMemMapFs(), "https": afero.NewMemMapFs()}
	)
	for _, o := range opts {
		switch opt := o.(type) {
		case afero.Fs:
			fsResolvers["file"] = opt
		case map[string]afero.Fs:
			fsResolvers = opt
		case lib.RuntimeOptions:
			rtOpts = opt
		case *logrus.Logger:
			logger = opt
		default:
			tb.Fatalf("unknown test option %q", opt)
		}
	}
	registry := metrics.NewRegistry()
	builtinMetrics := metrics.RegisterBuiltinMetrics(registry)
	return js.New(
		&lib.TestPreInitState{
			Logger:         logger,
			RuntimeOptions: rtOpts,
			BuiltinMetrics: builtinMetrics,
			Registry:       registry,
		},
		&loader.SourceData{
			URL:  &url.URL{Path: filename, Scheme: "file"},
			Data: []byte(data),
		},
		fsResolvers,
	)
}

func main1() {
	//code := `export let options = { vus: 12345 }; export default function() { return options.vus; };`
	//arc := &lib.Archive{
	//	Type:        "js",
	//	FilenameURL: &url.URL{Scheme: "file", Path: "/script"},
	//	K6Version:   consts.Version,
	//	Data:        []byte(code),
	//	Options:     lib.Options{VUs: null.IntFrom(999)},
	//	PwdURL:      &url.URL{Scheme: "file", Path: "/"},
	//	Filesystems: nil,
	//}
	var t testing.TB
	r, err := getSimpleRunner(t, "/script.js", `
			var parseHTML = require("k6/html").parseHTML;

			exports.options = { iterations: 1, vus: 1 };

			exports.default = function() {
				var doc = parseHTML("<html><div class='something'><h1 id='top'>Lorem ipsum</h1></div></html>");

				var o = doc.find("div").get(0).attributes()

				console.log(o)
			};
		`)
	require.NoError(t, err)

	ch := make(chan metrics.SampleContainer, 1000)
	initVU, err := r.NewVU(1, 1, ch)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	vu := initVU.Activate(&lib.VUActivationParams{RunContext: ctx})
	errC := make(chan error)
	go func() { errC <- vu.RunOnce() }()
	select {
	case <-time.After(15 * time.Second):
		cancel()
		t.Fatal("Test timed out")
	case err := <-errC:
		cancel()
		require.NoError(t, err)
	}
	fmt.Println(r, err)
}

func main() {
	var t *testing.T
	b, err := getSimpleBundle(t, "/script.js", `
					import k6 from "k6";
					export let _k6 = k6;
					export let dummy = "abc123";
					export default function() {
					console.log(3333)
				}
			`)
	logger := testutils.NewLogger(t)

	bi, err := b.Instantiate(logger, 0)
	exports := bi.Pgm.Exports
	//require.NotNil(t, exports)
	_, defaultOk := goja.AssertFunction(exports.Get("default"))
	//assert.True(t, defaultOk, "default export is not a function")

	k6 := exports.Get("_k6").ToObject(bi.Runtime)
	_, groupOk := goja.AssertFunction(k6.Get("group"))

	fmt.Println(defaultOk, k6, groupOk, "abc123", exports.Get("dummy").String())
	_, err = bi.GetCallableExport(consts.DefaultFn)(goja.Undefined())

	//assert.True(t, groupOk, "k6.group is not a function")
	fmt.Println(err)
}
