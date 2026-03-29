package main

import "fmt"

func main()  {
	prices := []float64{10.20,20}
	prices = append(prices,15.20)
	fmt.Println(prices)
}