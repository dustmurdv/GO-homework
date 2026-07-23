package main

import "fmt"

func main() {

	keyboard := NewProduct("P1", "Keyboard", 250000, 5, "PC")
	mouse := NewProduct("P2", "Mouse", 150000, 10, "PC")

	cart := NewCart()

	cart.Add(keyboard)
	cart.Add(mouse)

	card := CardProcessor{}

	err := ProcessOrder(cart, "AKMAL001", card)

	if err != nil {
		fmt.Println(err)
	}
}