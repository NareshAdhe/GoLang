package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func readBalanceFromFile() (float64,error) {
	data, err := os.ReadFile(accountBalanceFile)
	fmt.Println(data)
	if err != nil {
		return 1000,errors.New("Failed to read from the file!")
	}
	balanceText := string(data)
	fmt.Println(balanceText)
	balance, err := strconv.ParseFloat(balanceText,64)

	if err != nil {
		return 1000, errors.New("Failed to parse to float value!")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) error {
	balanceText := fmt.Sprint(balance)
	return os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Something went wrong!")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("Can't continue. Sorry!")
	}

	fmt.Println("Welcome to Go Bank!")
Loop:
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount < 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
			} else {
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount: ", accountBalance)
			}
			err := writeBalanceToFile(accountBalance)
			if err != nil {
				fmt.Println("Could not save balance:", err)
			}
		case 3:
			var withdrawAmount float64
			fmt.Print("Your withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount < 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Must be smaller than accountBalance: ", accountBalance)
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Balance updated! New amount: ", accountBalance)
			}
			err := writeBalanceToFile(accountBalance)
			if err != nil {
				fmt.Println("Could not save balance:", err)
			}
		case 4:
			fmt.Println("Goodbye")
			break Loop
		default:
			fmt.Println("Invalid choice")
		}
	}
	fmt.Println("Thanks for choosing our bank!")
}
