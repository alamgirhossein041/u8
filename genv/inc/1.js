export function setup() {
	let a=4444
	let bb=5555
	console.log(a)
	return {
		a,
		bb
	};
}
export default function (){
	let c=setup()
	console.log(c.bb)
}
