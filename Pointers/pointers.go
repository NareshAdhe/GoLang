package main

import "fmt"

func main()  {
	age := 32

	fmt.Println("Before Age: ",age)

	var agePointer *int
	agePointer = &age

	editAdultAge(agePointer)
	fmt.Println("After Age: ",age)
}

func editAdultAge(age *int) {
	*age -= 18
}