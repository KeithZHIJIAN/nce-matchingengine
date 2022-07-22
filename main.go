package main

import (
	orderbook "github.com/KeithZHIJIAN/nce-matchingengine/internal/pkg/orderbook"
)

var symbols = [...]string{"BTCUSD", "ETHUSD", "XRPUSD"}

func main() {
	var forever chan struct{}
	for _, symbol := range symbols {
		ob := orderbook.NewOrderBook(symbol)
		go ob.Start()
	}
	<-forever
}
