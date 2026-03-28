package main

import (
	"errors"
	"fmt"
	"os"
)

const fileName = "portfolio.txt"

func writeValuesToFile(pbt float64, pat float64, ratio float64) error {
	content := fmt.Sprintf("Profit before tax: %.2f\nProfit after tax: %.2f\nRatio: %.3f", pbt, pat, ratio)
	return os.WriteFile(fileName, []byte(content), 0644)
}

func main() {
	var revenue, expense, tax_rate float64

	err1 := takeInput(&revenue, "Revenue: ")
	err2 := takeInput(&expense, "Expense: ")
	err3 := takeInput(&tax_rate, "Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		fmt.Println("Can't continue. Sorry!!")
		return
	}

	profit_before_tax, profit_after_tax, ratio := calculateProfitValues(revenue, expense, tax_rate)

	fmt.Printf("Profit before tax: %.2f\n", profit_before_tax)
	fmt.Printf("Profit after tax: %.2f\n", profit_after_tax)
	fmt.Printf("Profit ratio: %.3f\n", ratio)

	err := writeValuesToFile(profit_before_tax, profit_after_tax, ratio)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Can't continue. Sorry!!")
		return
	}
}

func takeInput(variable *float64, outputString string) error {
	
	fmt.Print(outputString)
	fmt.Scan(variable)

	if(*variable<=0) {
		return errors.New("The value should be a positive number.")
	}

	return nil
}

func calculateProfitValues(revenue, expense, tax_rate float64) (float64, float64, float64) {
	profit_before_tax := revenue - expense
	profit_after_tax := profit_before_tax - (revenue*tax_rate)/100
	ratio := profit_before_tax / profit_after_tax

	return profit_before_tax, profit_after_tax, ratio
}
