import ta from "k6/ta";

import {Nats} from 'k6/x/nats';

const natsConfig = {
	servers: ['nats://54.160.229.90:80'],
	unsafe: true,
};

const publisher = new Nats(natsConfig);
const subscriber = new Nats(natsConfig);

export default function (data) {
	subscriber.subscribe('OrderResult', (msg) => {
		console.log(msg)
	});

	subscriber.subscribe('Order.*', (msg) => {
		console.log(msg)
	});

	let jma = ta.jma(close, 7, 50, 1)

	let dwma = ta.dwma(close, 10)

	console.log("close", close.tail(3).reverse())


	console.log("jma", jma.tail(3).reverse())
	console.log("dwma", dwma.tail(3).reverse())
	if (jma.crossOver(dwma)){
		console.log( {
			symbol:"ETHUSDT",
			side:"BUY",
			price:close.last(),
			quantity:1
		})
		publisher.publish('Order.Open.Long', {
			symbol:"ETHUSDT",
			side:"BUY",
			price:close.last(),
			quantity:1
		});
	}

	if (jma.crossUnder(dwma)){
		console.log( {
			symbol:"ETHUSDT",
			side:"SELL",
			price:close.last(),
			quantity:1
		})
		publisher.publish('Order.Close.Long', {
			symbol:"ETHUSDT",
			side:"SELL",
			price:close.last(),
			quantity:1
		});
	}

	console.log("=========\n")

}

