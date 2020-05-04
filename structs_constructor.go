package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Price  float64
	Volume int
	Buy    bool
}

func NewTrade(symbol string, price float64, volume int, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol cannot be null")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("Volume cannot be less than 0")
	}
	if price <= 0.0 {
		return nil, fmt.Errorf("Traded price cannot be less than 0")
	}

	trade := &Trade{
		Symbol: symbol,
		Price:  price,
		Volume: volume,
		Buy:    buy,
	}

	return trade, nil
}

func main() {
	trade, _ := NewTrade("Reliance", 1000, 10, true)

	fmt.Println(*trade)
}
