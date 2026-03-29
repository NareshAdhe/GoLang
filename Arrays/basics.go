package main

import "fmt"

func main()  {
	var productNames = [4]string{"A book"}
	prices := []float64{10.5,12.4,15.3}
	fmt.Println(prices)

	productNames[2] = "A carpet"
	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 20.50
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices))
	fmt.Println(len(highlightedPrices),cap(highlightedPrices))
}