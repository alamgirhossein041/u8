package main

import (
	"context"
	"fmt"
	"github.com/dop251/goja"
	"github.com/oxtoacart/bpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/uvite/v9/js"
	"github.com/uvite/v9/lib"
	"github.com/uvite/v9/lib/consts"
	"github.com/uvite/v9/lib/netext"
	"github.com/uvite/v9/lib/types"
	"github.com/uvite/v9/loader"
	"github.com/uvite/v9/metrics"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

func getPreInitState(logger *logrus.Logger, rtOpts *lib.RuntimeOptions) *lib.InitState {

	if rtOpts == nil {
		rtOpts = &lib.RuntimeOptions{}
	}
	reg := metrics.NewRegistry()
	return &lib.InitState{
		Logger:         logger,
		RuntimeOptions: *rtOpts,
		Registry:       reg,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(reg),
	}
}
func getSimpleBundle(filename, data string, opts ...interface{}) (*js.Bundle, error) {
	fs := afero.NewMemMapFs()
	var rtOpts *lib.RuntimeOptions
	//var logger *logrus.Logger
	for _, o := range opts {
		switch opt := o.(type) {
		case afero.Fs:
			fs = opt
		case lib.RuntimeOptions:
			rtOpts = &opt
		//case *logrus.Logger:
		//	logger = opt
		default:
			fmt.Println("unknown test option %q", opt)
		}
	}
	logger := logrus.New()

	return js.NewBundle(
		getPreInitState(logger, rtOpts),
		&loader.SourceData{
			URL:  &url.URL{Path: filename, Scheme: "file"},
			Data: []byte(data),
		},
		map[string]afero.Fs{"file": fs, "https": afero.NewMemMapFs()},
	)
}
func main() {

	ch := make(chan bool, 1)
	logger := logrus.New()

	h := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			ch <- true
		}()

		file, _, err := r.FormFile("file")

		bytes := make([]byte, 3)
		_, err = file.Read(bytes)
		fmt.Println(err)
	}
	srv := httptest.NewServer(http.HandlerFunc(h))
	defer srv.Close()

	fs := afero.NewMemMapFs()
	fs.MkdirAll("/path/to", 0o755)
	afero.WriteFile(fs, "/path/to/file.bin", []byte("hi!"), 0o644)

	b, err := getSimpleBundle("/path/to/script.js",
		fmt.Sprintf(`
			import http from "k6/http";
			//let binFile = open("/path/to/file.bin", "b");
			export default function() {
				  const res = http.get('https://httpbin.test.k6.io/');
				  
				console.log(res)
				return true;
			}
			`), fs)

	bi, err := b.Instantiate(logger, 0)

	registry := metrics.NewRegistry()
	builtinMetrics := metrics.RegisterBuiltinMetrics(registry)

	root, err := lib.NewGroup("", nil)

	bi.ModuleVUImpl.Status = &lib.State{
		Options: lib.Options{},
		Logger:  logger,
		Group:   root,
		Transport: &http.Transport{
			DialContext: (netext.NewDialer(
				net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 60 * time.Second,
					DualStack: true,
				},
				netext.NewResolver(net.LookupIP, 0, types.DNSfirst, types.DNSpreferIPv4),
			)).DialContext,
		},
		BPool:          bpool.NewBufferPool(1),
		Samples:        make(chan metrics.SampleContainer, 500),
		BuiltinMetrics: builtinMetrics,
		Tags:           lib.NewVUStateTags(registry.RootTagSet()),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bi.ModuleVUImpl.Ctx = ctx
	v, err := bi.GetCallableExport(consts.DefaultFn)(goja.Undefined())

	fmt.Println(v, err)
}
