
 let abc=ta.rsi(close,5)
// let aa=ta.series()
// aa.push("hold")
// aa.push("buy")
// aa.push("buy")
// aa.push("sell")
// aa.push("sell")
// let changed=ta.change(aa,2)
 console.log("rsi",abc.tail(5) )
console.log("close",close.tail(5))

//console.log(high.last())
let c=ta.atr(high,low,close,10)
// let c=ta.sma(fuck,10)
  console.log("atr",c.tail(5))

console.log("==jsend===")
