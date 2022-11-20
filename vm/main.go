package main

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/uvite/u8/js"
	"github.com/uvite/u8/lib"
	"github.com/uvite/u8/loader"
	"github.com/uvite/u8/metrics"
	"github.com/uvite/u8/series"
	fta "github.com/uvite/u8/ta"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	apiKey    = "your api key"
	secretKey = "your secret key"
	interval  = "1m"
	symbol    = "ETHUSDT"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func newReader(file string) *csv.Reader {
	f, err := os.Open(file)
	checkErr(err)

	gzReader, err := gzip.NewReader(f)
	checkErr(err)

	csvReader := csv.NewReader(gzReader)
	csvReader.Comma = ','
	csvReader.ReuseRecord = true

	return csvReader
}

func readOHLCV(file string) fta.OHLCV {
	csvReader := newReader(file)

	ohlcv, err := fta.ReadCSV(csvReader, int64(time.Minute), fta.Seconds)
	checkErr(err)

	ohlcv = ohlcv.Resample(int64(time.Hour))

	return ohlcv
}

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

	logger := logrus.New()

	const (
		period        = 20
		stdMultiplier = 2.1
	)

	//ohlcv := readOHLCV("./BTCUSDT.csv.gz")

	// Calculate indicators.
	//sma := fta.SMA(ohlcv.Close, period)
	//fmt.Println(sma.IndexAt(1))

	rtOpts := lib.RuntimeOptions{Genv: map[string]any{

		"okk": 4444,
	}}
	fs := afero.NewOsFs()
	pwd, err := os.Getwd()
	fmt.Println(pwd)
	sourceData, err := loader.ReadSource(logger, "./1.js", pwd, map[string]afero.Fs{"file": fs}, nil)

	//fmt.Println(ohlcv.Close)
	b, err := getSimpleBundle("/path/to/script.js",
		fmt.Sprintf(`
					import k6 from "k6";
					export let _k6 = k6;
					export let dummy = "abc123";
 					import _ from "https://cdnjs.cloudflare.com/ajax/libs/lodash.js/4.17.21/lodash.min.js";
					export sma=tart.NewSma(5)
					export default function() {
						%s
					}
			`, sourceData.Data), fs, rtOpts)

	bi, err := b.Instantiate(logger, 0)
	//bi.Runtime.Set("ok", "3333")
	fmt.Println(err)
	exports := bi.Pgm.Exports
	//bb := exports.Get("_k6").ToObject(bi.Runtime)
	//for _, k := range bb.Keys() {
	//	fmt.Println(k)
	//}
	//fmt.Println(exports.Get("_k6"))

	var client = binance.NewClient(apiKey, secretKey)

	klines, err := client.NewKlinesService().Symbol(symbol).
		Interval(interval).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	kk := []binance.Kline{}

	for _, event := range klines {
		k := binance.Kline{

			Open:     event.Open,
			High:     event.High,
			Low:      event.Low,
			Close:    event.Close,
			Volume:   event.Volume,
			OpenTime: event.OpenTime,
		}
		kk = append(kk, k)

	}

	//cci1 := ok.NewCci(5)
	wsKlineHandler := func(event *binance.WsKlineEvent) {

		if event.Kline.IsFinal {
			k := binance.Kline{

				Open:     event.Kline.Open,
				High:     event.Kline.High,
				Low:      event.Kline.Low,
				Close:    event.Kline.Close,
				Volume:   event.Kline.Volume,
				OpenTime: event.Kline.StartTime,

				//OpenTime: types.NewTimeFromUnix(0, event.Kline.EndTime*int64(time.Millisecond)),
			}
			kk = append(kk, k)

			aline, err := getData(kk, int64(time.Minute), Seconds)
			//fmt.Println(k, err)

			close := (aline.Close)
			b.Set("close", close)
			b.Set("bbb", event.Kline.Close)
			//fmt.Println(close.Len(), close.Values())
			sma := fta.EMA(close, 30, true)
			fmt.Println(sma.At(-1), "11111")

			//fmt.Println(close.Rolling(period).Mean().IndexAt(0), "22222")

			if call, ok := goja.AssertFunction(exports.Get("default")); ok {
				if _, err = call(goja.Undefined()); err != nil {

				}
			}
		}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsKlineServe(symbol, interval, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

type UnixTime int

const (
	Seconds UnixTime = iota
	Milliseconds
)

type OHLCV struct{ Open, High, Low, Close, Volume, Time series.Data }
type DType = series.DType

func getData(kline []binance.Kline, freq int64, unixTime UnixTime) (ohlcv OHLCV, err error) {
	const (
		Time = iota
		Open
		High
		Low
		Close
		Volume
	)

	var (
		T []int64
		O,
		H,
		L,
		C,
		TT,
		V []series.DType
	)
	for _, record := range kline {

		ts := record.OpenTime

		o, err := strconv.ParseFloat(record.Open, 64)
		if err != nil {
			return ohlcv, fmt.Errorf("parse float: field 'Open': %w", err)
		}

		h, err := strconv.ParseFloat(record.High, 64)
		if err != nil {
			return ohlcv, fmt.Errorf("parse float: field 'High': %w", err)
		}

		l, err := strconv.ParseFloat(record.Low, 64)
		if err != nil {
			return ohlcv, fmt.Errorf("parse float: field 'Low': %w", err)
		}

		c, err := strconv.ParseFloat(record.Close, 64)
		if err != nil {
			return ohlcv, fmt.Errorf("parse float: field 'Close': %w", err)
		}

		v, err := strconv.ParseFloat(record.Volume, 64)
		if err != nil {
			return ohlcv, fmt.Errorf("parse float: field 'Volume': %w", err)
		}

		switch unixTime {
		case Seconds:
			ts *= int64(time.Second)
		case Milliseconds:
			ts *= int64(time.Millisecond)
		}

		T = append(T, ts)
		TT = append(O, DType(ts))
		O = append(O, DType(o))
		H = append(H, DType(h))
		L = append(L, DType(l))
		C = append(C, DType(c))
		V = append(V, DType(v))

	}
	//fmt.Println(T)
	ohlcv = OHLCV{
		Open:   series.MakeData(freq, T, O),
		High:   series.MakeData(freq, T, H),
		Low:    series.MakeData(freq, T, L),
		Close:  series.MakeData(freq, T, C),
		Volume: series.MakeData(freq, T, V),
		Time:   series.MakeData(freq, T, TT),
	}

	return ohlcv, nil
}
