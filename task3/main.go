package main

import "fmt"

func main() {

	cart := NewCart()

	product1 := NewProduct(1, "Keyboard", 250000, 10)
	product2 := NewProduct(2, "Mouse", 120000, 20)
	product3 := NewProduct(3, "Monitor", 1800000, 5)

	cart.AddProduct(product1)
	cart.AddProduct(product2)
	cart.AddProduct(product3)

	fmt.Println("Products in cart:")

	for _, product := range cart.Products() {
		fmt.Println(product.Name(), "-", product.Price())
	}

	fmt.Println()

	fmt.Println("Total items:", cart.TotalItems())

	fmt.Println()

	err := cart.RemoveProduct(2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("After removing Mouse:")

	for _, product := range cart.Products() {
		fmt.Println(product.Name())
	}

	fmt.Println()

	fmt.Println("Total items:", cart.TotalItems())

	fmt.Println()

	cart.Clear()

	fmt.Println("After clear:", cart.TotalItems())
}