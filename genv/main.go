package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/uvite/u8/js"
	"github.com/uvite/u8/lib"
	"github.com/uvite/u8/loader"
	"github.com/uvite/u8/metrics"
	"gopkg.in/guregu/null.v3"
	"net/url"
	"time"
)

func getSimpleRunner(filename, data string, opts ...interface{}) (*js.Runner, error) {
	var (
		rtOpts      = lib.RuntimeOptions{CompatibilityMode: null.NewString("base", true)}
		logger      = logrus.New()
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
			logger.Fatalf("unknown test option %q", opt)
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
func main() {

	r, err := getSimpleRunner("/script.js", `
			var parseHTML = require("k6/html").parseHTML;

			exports.options = { iterations: 1, vus: 1 };

			exports.default = function() {
				var doc = parseHTML("<html><div class='something'><h1 id='top'>Lorem ipsum</h1></div></html>");

				var o = doc.find("div").get(0).attributes()

				console.log(o)
			};
		`)

	ch := make(chan metrics.SampleContainer, 1000)
	initVU, err := r.NewVU(1, 1, ch)

	ctx, cancel := context.WithCancel(context.Background())
	vu := initVU.Activate(&lib.VUActivationParams{RunContext: ctx})
	errC := make(chan error)
	go func() { errC <- vu.RunOnce() }()
	select {
	case <-time.After(15 * time.Second):
		cancel()

	case err := <-errC:
		cancel()
		require.NoError(t, err)
	}
	fmt.Println(r, err)
}
