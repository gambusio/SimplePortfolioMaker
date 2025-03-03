package main

import (
	"fmt"
)

type Stock struct {
	ticker   string
	price    int64
	low52    int64
	dividend int64
	yield    float64
}

var stocks [4]Stock

func CurrencyToString(c int64) string {
	return fmt.Sprintf("$%d.%02d", c/100, c%100)
}

func StringToCurrency(s string) int64 {
	var dec, cent int64
	fmt.Sscanf(s, "$%d.%02d", &dec, &cent)
	return dec*100 + cent
}

func PrintStocks() {
	fmt.Printf("Ticker\tPrice\tLow52\tDividend\tYield\n")
	fmt.Printf("------------------------------------------\n")
	for _, stock := range stocks {
		fmt.Printf("%s\t%s\t%s\t%s\t\t%f\n", stock.ticker, CurrencyToString(stock.price), CurrencyToString(stock.low52), CurrencyToString(stock.dividend), stock.yield)
	}
}

func main() {

	stocks[0] = Stock{"BPYPP", StringToCurrency("$16.10"), StringToCurrency("$12.95"), StringToCurrency("$1.63"), 0}
	stocks[1] = Stock{"BPYPO", StringToCurrency("$15.37"), StringToCurrency("$12.74"), StringToCurrency("$1.59"), 0}
	stocks[2] = Stock{"BPYPN", StringToCurrency("$14.22"), StringToCurrency("$11.51"), StringToCurrency("$1.44"), 0}
	stocks[3] = Stock{"BPYPM", StringToCurrency("$16.15"), StringToCurrency("$12.82"), StringToCurrency("$1.56"), 0}
	PrintStocks()
}
