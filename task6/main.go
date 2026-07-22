package main

import "fmt"

func main() {

	inventory := NewInventory()

	p1 := NewProduct("1", "Keyboard", 250000, 5, "PC")
	p2 := NewProduct("2", "Mouse", 120000, 10, "PC")
	p3 := NewProduct("3", "Book", 89000, 2, "Books")

	inventory.AddProduct(p1)
	inventory.AddProduct(p2)
	inventory.AddProduct(p3)

	fmt.Println("Total value:", inventory.TotalValue())

	fmt.Println()

	fmt.Println("Books:")

	books := inventory.FindByCategory("Books")

	for _, b := range books {
		fmt.Println(b.Name())
	}

	fmt.Println()

	fmt.Println("Low stock:")

	low := inventory.LowStock(3)

	for _, p := range low {
		fmt.Println(p.Name(), p.Stock())
	}
}