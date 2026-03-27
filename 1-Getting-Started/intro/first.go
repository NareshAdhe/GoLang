package main

import (
	"fmt"
	"first-module/testpackage"
)

func main()  {
	var result int
	result = testpackage.Add(5,10)
	fmt.Println("The total is:", result)
}