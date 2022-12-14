package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Interval string

func (i Interval) Minutes() int {
	m, ok := SupportedIntervals[i]
	if !ok {
		return ParseInterval(i)
	}
	return m
}

func (i Interval) Duration() time.Duration {
	return time.Duration(i.Minutes()) * time.Minute
}

func (i *Interval) UnmarshalJSON(b []byte) (err error) {
	var a string
	err = json.Unmarshal(b, &a)
	if err != nil {
		return err
	}

	*i = Interval(a)
	return
}

func (i Interval) String() string {
	return string(i)
}

type IntervalSlice []Interval

func (s IntervalSlice) StringSlice() (slice []string) {
	for _, interval := range s {
		slice = append(slice, `"`+interval.String()+`"`)
	}
	return slice
}

var Interval1m = Interval("1m")
var Interval3m = Interval("3m")
var Interval5m = Interval("5m")
var Interval15m = Interval("15m")
var Interval30m = Interval("30m")
var Interval1h = Interval("1h")
var Interval2h = Interval("2h")
var Interval4h = Interval("4h")
var Interval6h = Interval("6h")
var Interval12h = Interval("12h")
var Interval1d = Interval("1d")
var Interval3d = Interval("3d")
var Interval1w = Interval("1w")
var Interval2w = Interval("2w")
var Interval1mo = Interval("1mo")

func ParseInterval(input Interval) int {
	t := 0
	index := 0
	for i, rn := range string(input) {
		if rn >= '0' && rn <= '9' {
			t = t*10 + int(rn-'0')
		} else {
			index = i
			break
		}
	}
	switch strings.ToLower(string(input[index:])) {
	case "m":
		return t
	case "h":
		t *= 60
	case "d":
		t *= 60 * 24
	case "w":
		t *= 60 * 24 * 7
	case "mo":
		t *= 60 * 24 * 30
	default:
		panic("unknown input: " + input)
	}
	return t
}

var SupportedIntervals = map[Interval]int{
	Interval1m:  1,
	Interval3m:  3,
	Interval5m:  5,
	Interval15m: 15,
	Interval30m: 30,
	Interval1h:  60,
	Interval2h:  60 * 2,
	Interval4h:  60 * 4,
	Interval6h:  60 * 6,
	Interval12h: 60 * 12,
	Interval1d:  60 * 24,
	Interval3d:  60 * 24 * 3,
	Interval1w:  60 * 24 * 7,
	Interval2w:  60 * 24 * 14,
	Interval1mo: 60 * 24 * 30,
}

// IntervalWindow is used by the indicators
type IntervalWindow struct {
	// The interval of kline
	Interval Interval `json:"interval"`

	// The windows size of the indicator (for example, EWMA and SMA)
	Window int `json:"window"`

	// RightWindow is used by the pivot indicator
	RightWindow int `json:"rightWindow"`
}

type IntervalWindowBandWidth struct {
	IntervalWindow
	BandWidth float64 `json:"bandWidth"`
}

func (iw IntervalWindow) String() string {
	return fmt.Sprintf("%s (%d)", iw.Interval, iw.Window)
}
