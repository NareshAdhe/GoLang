package main

import "fmt"

func main(){
	var revenue,expense,tax_rate float64

	takeInput(&revenue,&expense,&tax_rate)
	profit_before_tax,profit_after_tax,ratio := calculateProfitValues(revenue,expense,tax_rate)

	fmt.Printf("Profit before tax: %.2f\n",profit_before_tax)
	fmt.Printf("Profit after tax: %.2f\n",profit_after_tax)
	fmt.Printf("Profit ratio: %.2f\n",ratio)
}

func takeInput(revenue *float64,expense *float64,tax_rate *float64){
	fmt.Print("Revenue: ")
	fmt.Scan(revenue)

	fmt.Print("Expense: ")
	fmt.Scan(expense)

	fmt.Print("Tax Rate: ")
	fmt.Scan(tax_rate)
}

func calculateProfitValues(revenue,expense,tax_rate float64) (float64,float64,float64) {
	profit_before_tax := revenue-expense
	profit_after_tax := profit_before_tax - (revenue*tax_rate)/100
	ratio := profit_before_tax/profit_after_tax

	return profit_before_tax,profit_after_tax,ratio
}