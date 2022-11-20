package main

import (
	"fmt"
	"github.com/c9s/bbgo/pkg/indicator"
	"github.com/c9s/bbgo/pkg/types"
	tart "github.com/uvite/indicator"

	//fixedpoint "github.com/c9s/bbgo/pkg/fixedpoint"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/uvite/u8/js"
	"github.com/uvite/u8/lib"
	"github.com/uvite/u8/loader"
	"github.com/uvite/u8/metrics"
	"net/url"
	"os"
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
func atr(s int) *indicator.SMA {
	a := &indicator.SMA{IntervalWindow: types.IntervalWindow{Window: s}}
	return a
}
func main() {

	logger := logrus.New()

	const (
		period        = 20
		stdMultiplier = 2.1
	)

	rtOpts := lib.RuntimeOptions{Genv: map[string]any{
		//"close": close,

		"okk": 4444,
	}}
	fs := afero.NewOsFs()
	pwd, err := os.Getwd()

	sourceData, err := loader.ReadSource(logger, "./1.js", pwd, map[string]afero.Fs{"file": fs}, nil)

	//fmt.Println(ohlcv.Close)
	b, err := getSimpleBundle("/path/to/script.js",
		fmt.Sprintf(`
					import k6 from "k6";
					import ta from "k6/ta";
					//export let close=ta.slice();
					//export let high=ta.slice();
					//export let low=ta.slice();
					//export let abc="23423"; 
					export default function() {
						//close.push(33);
						console.log(b.last())

						%s
					}
			`, sourceData.Data), fs, rtOpts)
	//b.Set("close", close)
	bi, err := b.Instantiate(logger, 0)

	res := &tart.Series{}
	b.Set("b", res)

	res.Update(33)
	res.Update(44)
	res.Update(55)

	//b.Set("b", b)

	//bi.Runtime.Set("ok", "3333")
	//fmt.Println(err)
	exports := bi.Pgm.Exports

	//b.Set("close", close)
	//fmt.Println(exports.Get("okk"))

	//fmt.Println(close.Rolling(period).Mean().IndexAt(0), "22222")
	//fmt.Println(Close.Last(), 3333)
	if call, ok := goja.AssertFunction(exports.Get("default")); ok {
		if _, err = call(goja.Undefined()); err != nil {

		}
	}

	//vm := goja.New()
	//
	//new(require.Registry).Enable(vm)
	//console.Enable(vm)
	//vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	//
	//bb := atr(3)
	////bb.Update(33)
	////bb.Update(44)
	////bb.Update(55)
	//vm.Set("b", bb)
	//
	//script := `
	//
	//	 function a(){
	//
	//	 return b.last()
	//	}
	//
	//
	//
	//
	//`
	//
	//vm.RunString(script) // without the mapper it would have been s.Field
	////fmt.Println(res.Export())
	//
	//var fn func() float64
	//err = vm.ExportTo(vm.Get("a"), &fn)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(fn(), err)
	//
	//bb.Update(33)
	//bb.Update(44)
	//bb.Update(55)
	////vm.Set("atr", a)
	//fmt.Println(fn())

}
