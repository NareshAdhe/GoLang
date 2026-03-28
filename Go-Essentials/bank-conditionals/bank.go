package main

import (
	"fmt"
	"example.com/bank/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := fileOps.ReadFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Something went wrong!")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic("Can't continue. Sorry!")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach out to us 24/7: ",randomdata.PhoneNumber())
Loop:
	for {
		presentOptions()

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
			err := fileOps.WriteFloatToFile(accountBalanceFile,accountBalance)
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
			err := fileOps.WriteFloatToFile(accountBalanceFile,accountBalance)
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
