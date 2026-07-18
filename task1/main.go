package main

import "fmt"

func main() {

	a := uint8(0b10101010)
	b := uint8(0b11110000)

	andResult := a & b
	orResult := a | b
	xorResult := a ^ b
	andNotResult := a &^ b
	shiftResult := a >> 2

	fmt.Printf("AND     : %08b\n", andResult)
	fmt.Printf("OR      : %08b\n", orResult)
	fmt.Printf("XOR     : %08b\n", xorResult)
	fmt.Printf("AND NOT : %08b\n", andNotResult)
	fmt.Printf("SHIFT   : %08b\n", shiftResult)
}