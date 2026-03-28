package main

import (
	"example.com/structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserInput("Please enter your first name: ")
	userLastName := getUserInput("Please enter your last name: ")
	userBirthDate := getUserInput("Please enter your birthdate(MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserInput(outputString string) string {
	var userInput string
	fmt.Print(outputString)
	fmt.Scanln(&userInput)
	return userInput
}
