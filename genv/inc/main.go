package main

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/uvite/u8/js"
	"github.com/uvite/u8/lib"
	"github.com/uvite/u8/loader"
	"github.com/uvite/u8/metrics"
	"net/http"
	"net/http/httptest"
	"net/url"
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
			import k6 from "k6";
					export let _k6 = k6;
					export let dummy = "abc123";
					export default function() {
						console.log(dummy)
					}
			`), fs)

	bi, err := b.Instantiate(logger, 0)
	fmt.Println(err)
	exports := bi.Pgm.Exports
	bb := exports.Get("_k6").ToObject(bi.Runtime)
	for _, k := range bb.Keys() {
		fmt.Println(k)
	}
	//fmt.Println(exports.Get("_k6"))
	_, defaultOk := goja.AssertFunction(exports.Get("default"))
	fmt.Println(defaultOk, exports.Get("dummy").String())

	//k6 := exports.Get("_k6").ToObject(bi.Runtime)
	//fmt.Println(k6)
	//_, groupOk := goja.AssertFunction(k6.Get("group"))
	////assert.True(t, groupOk, "k6.group is not a function")
	//fmt.Println(groupOk)
}
