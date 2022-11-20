package main

import (
	"fmt"
	"github.com/uvite/indicator"
	"github.com/uvite/u8/series"
	"github.com/uvite/u8/vm/ta/fta"
)

func main() {
	sma := tart.NewSma(5)
	//close:=[]fta.DType{}
	var close []series.DType
	var T []int64
	for i := 0; i < 20; i++ {
		T = append(T, int64(i))
		close = append(close, float64(i%7))
		val := sma.Update(float64(i % 7))
		if sma.Valid() {
			fmt.Printf("sma[%v] = %.4f\n", i, val)
		} else {
			fmt.Printf("sma[%v] = unavail.\n", i)
		}
	}
	closes := series.MakeData(10, T, close)

	aa := fta.SMA(closes, 5)
	fmt.Println(aa)

}
