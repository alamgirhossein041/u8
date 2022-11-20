let abc=ta.sma(close,5)
//let cc=close.Clone()
function a(){
	var array = [1];
	var other = _.concat(array, 2, [3], [[4]]);

	//		console.log(other);
	//console.log(typeof close);

	console.log(close.addScalar(1).values())
	const map1 = close.values().map(x => x -3);
//console.log(map1)
//console.log(close.values())
//_.reduce(close.values(), function(sum, n) {
// console.log(sum)
// return sum;
//}, 0);
	for(let i=0;i<close.len();i++){
		//console.log(close.indexAt(i))
		//console.log()
	}
	return close.indexAt(2)
}
console.log(abc.indexAt(1))
console.log(a())