package main

import "fmt"

func main() {

	customer, err := NewCustomer(
		1,
		"Akmal",
		"akmal@gmail.com",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	customer.AddOrder("Keyboard")
	customer.AddOrder("Mouse")
	customer.AddOrder("Monitor")

	fmt.Println("ID:", customer.ID())
	fmt.Println("Name:", customer.Name())
	fmt.Println("Email:", customer.Email())

	fmt.Println()

	fmt.Println("Orders:")

	for _, order := range customer.Orders() {
		fmt.Println("-", order)
	}
}