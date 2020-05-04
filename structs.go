package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price
	if trade.Buy {
		value = -value
	}

	return value
}

func main() {
	t := Trade{
		Symbol: "Reliance",
		Volume: 100,
		Price:  1200,
		Buy:    true,
	}

	fmt.Printf("%+v\n", t)

	value := t.Value()

	fmt.Printf("Value= %v\n", value)
}
