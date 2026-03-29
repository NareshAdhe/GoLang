package main

import "fmt"

type Product struct {
	id string
	title string
	price float64
}

func main()  {
	hobbies := []string{"Cricket","Singing","Technology"}
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	// slice1 := hobbies[0:2] 
	slice1 := hobbies[:2]
	fmt.Println(slice1)
	slice1 = hobbies[1:3]
	fmt.Println(slice1)

	dynamic_array := []string{"Movies","Singing"}
	dynamic_array[1] = "Coding"
	dynamic_array = append(dynamic_array, "Watching Matches")
	fmt.Println(dynamic_array)

	products := []Product{{"1","Title1",200.0},{"2","Title2",500.0}}
	fmt.Println(products)

	products = append(products, Product{"3","Title3",700})

	fmt.Println(products)
}