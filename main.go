package main

import (
	"fmt"
)

const EQUITIES = 4

type Stock struct {
	ticker   string
	price    int64
	low52    int64
	dividend int64
	yield    float64
}

/*
Array que almacena la información necesaria para determinar las compras que se deben realizar.
*/
var stocks [4]Stock

/*
CurrencyToString(c) dada cierta cantidad c de centavos de dólar (formato para cálculos interno) devuelve una cadena
equivalente en el formato "$xx.xx"
*/
func CurrencyToString(c int64) string {
	return fmt.Sprintf("$%d.%02d", c/100, c%100)
}

/*
StringToCurrency(s) dada una cadena de texto s que representa cierta cantidad de dólares en el formato $xx.xx
devuelve la reprensentación de los centavos en formato entero de 64 bits
*/
func StringToCurrency(s string) int64 {
	var dec, cent int64
	_, err := fmt.Sscanf(s, "$%d.%02d", &dec, &cent)
	if err != nil {
		return 0
	}
	return dec*100 + cent
}

/*
PrintStocks() Muestra la información de las acciones seleccionadas
*/
func PrintStocks() {
	fmt.Printf("Ticker\tPrice\tLow52\tDividend\tYield\n")
	fmt.Printf("------------------------------------------\n")
	for _, stock := range stocks {
		fmt.Printf("%s\t%s\t%s\t%s\t\t%.2f%%\n", stock.ticker, CurrencyToString(stock.price), CurrencyToString(stock.low52), CurrencyToString(stock.dividend), stock.yield)
	}
}

/*
Inicializa los datos necesarios y calcula la rentabilidad por dividendo de cada activo
*/
func Init() {
	stocks[0] = Stock{"BPYPP", StringToCurrency("$14.66"), StringToCurrency("$12.95"), StringToCurrency("$1.63"), 0}
	stocks[1] = Stock{"BPYPO", StringToCurrency("$14.35"), StringToCurrency("$12.74"), StringToCurrency("$1.59"), 0}
	stocks[2] = Stock{"BPYPN", StringToCurrency("$13.91"), StringToCurrency("$11.51"), StringToCurrency("$1.44"), 0}
	stocks[3] = Stock{"BPYPM", StringToCurrency("$15.20"), StringToCurrency("$12.82"), StringToCurrency("$1.56"), 0}
	for i := 0; i < EQUITIES; i++ {
		stocks[i].yield = (float64(stocks[i].dividend) / float64(stocks[i].price)) * 100.00
	}
}

/*
Distributive() Determina la cantidad de acciones a comprar dándole igual peso a cada una de las posibilidades
*/
func Distributive(capital int64) {
	capitalPerEquity := capital / EQUITIES
	fmt.Println("Lot size: ", CurrencyToString(capitalPerEquity))
	for i := 0; i < EQUITIES; i++ {
		fmt.Println(capitalPerEquity/stocks[i].price, stocks[i].ticker, "shares [", CurrencyToString(capitalPerEquity/stocks[i].price*stocks[i].price), "]")
	}
}

/*
ByYield() Determina el activo con mayor rentabilidad por dividendo y calcula la posible compra de acciones de ese activo
*/
func ByYield(capital int64) {
	fmt.Println("Lot size: ", CurrencyToString(capital))
	maxYield := -1.0
	maxPos := -1
	for i := 0; i < EQUITIES; i++ {
		if stocks[i].yield > maxYield {
			maxPos = i
			maxYield = stocks[i].yield
		}
	}
	fmt.Println(capital/stocks[maxPos].price, stocks[maxPos].ticker, "shares [", CurrencyToString(capital/stocks[maxPos].price*stocks[maxPos].price), "]")
}

/*
By52Low() Determina el activo más próximo al mínimo de 52 semanas y calcula la posible compra de acciones de ese activo
*/
func By52Low(capital int64) {
	fmt.Println("Lot size: ", CurrencyToString(capital))
	var percentage float64
	maxPercentage := -1.0
	maxPos := -1
	for i := 0; i < EQUITIES; i++ {
		percentage = (float64(stocks[i].low52) / float64(stocks[i].price)) * 100.0
		if percentage > maxPercentage {
			maxPos = i
			maxPercentage = percentage
		}
	}
	fmt.Println(capital/stocks[maxPos].price, stocks[maxPos].ticker, "shares [", CurrencyToString(capital/stocks[maxPos].price*stocks[maxPos].price), "]")
}

/*
weighted() Calcula una compra ponderada (A% por rentabilidad, B% por ceranía a mínimos, C% balanceado)

	Modificar Pesos
*/
func weighted(capital int64) {
	A := 0.3
	B := 0.3
	C := 0.4
	var capPerShare [EQUITIES]int64
	var percentage float64
	maxPercentage := -1.0
	maxPos := -1

	for i := 0; i < EQUITIES; i++ {
		capPerShare[i] = int64(float64(capital) * C / EQUITIES)
	}

	for i := 0; i < EQUITIES; i++ {
		percentage = (float64(stocks[i].low52) / float64(stocks[i].price)) * 100.0
		if percentage > maxPercentage {
			maxPos = i
			maxPercentage = percentage
		}
	}
	capPerShare[maxPos] += int64(float64(capital) * B)
	maxYield := -1.0
	maxPos = -1
	for i := 0; i < EQUITIES; i++ {
		if stocks[i].yield > maxYield {
			maxPos = i
			maxYield = stocks[i].yield
		}
	}
	capPerShare[maxPos] += int64(float64(capital) * A)
	for i := 0; i < EQUITIES; i++ {
		fmt.Println(capPerShare[i]/stocks[i].price, stocks[i].ticker, "shares [", CurrencyToString(capPerShare[i]/stocks[i].price*stocks[i].price), "]")
	}
}

func main() {
	var capital string = "$116.00"

	Init()
	PrintStocks()

	fmt.Println("\nBalanced purchase")
	Distributive(StringToCurrency(capital))
	fmt.Println("\nBuy for higher dividend yield")
	ByYield(StringToCurrency(capital))
	fmt.Println("\nBuy the one closest to the 52 week low")
	By52Low(StringToCurrency(capital))
	fmt.Println("\nWeighted purchase (30% by yield, 30% by closeness to minimum 52 weeks, 40% balanced)")
	weighted(StringToCurrency(capital))
}
