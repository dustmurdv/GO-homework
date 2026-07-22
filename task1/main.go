package main

import "fmt"

func main() {

	product := NewProduct(1, "Keyboard", 250000, 10)

	fmt.Println("ID:", product.ID())
	fmt.Println("Name:", product.Name())
	fmt.Println("Price:", product.Price())
	fmt.Println("Stock:", product.Stock())

	fmt.Println()

	err := product.SetPrice(300000)

	if err != nil {
		fmt.Println(err)
	}

	product.AddStock(5)

	err = product.Reserve(3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("After changes")
	fmt.Println("Price:", product.Price())
	fmt.Println("Stock:", product.Stock())
}