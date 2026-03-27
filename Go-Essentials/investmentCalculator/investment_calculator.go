package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main(){
	var investmentAmount float64
	var years float64
	interestRate := 5.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	fmt.Print("Enter interest rate: ")
	fmt.Scan(&interestRate)

	futureAmount,futureRealAmount := calculateAmounts(investmentAmount,years,interestRate) 

	futureAmountString := fmt.Sprintf("Future Value: %.2f\n",futureAmount)
	futureRealAmountString := fmt.Sprintf("Future Real Value: %.2f\n",futureRealAmount)

	fmt.Print(futureAmountString,futureRealAmountString);
}

func calculateAmounts(investmentAmount, years, interestRate float64) (float64,float64) {
	futureAmount := investmentAmount * math.Pow(1 + interestRate / 100,years)
	futureRealAmount := futureAmount / math.Pow(1+inflationRate/100,years)
	
	return futureAmount,futureRealAmount
}