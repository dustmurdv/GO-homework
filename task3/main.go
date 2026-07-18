package main

import "fmt"

func isZeroInt(n int) bool {
	return n == 0
}

func isZeroString(s string) bool {
	return s == ""
}

func isZeroFloat(f float64) bool {
	return f == 0
}

func main() {

	fmt.Println(isZeroInt(0))
	fmt.Println(isZeroInt(5))

	fmt.Println(isZeroString(""))
	fmt.Println(isZeroString("Hello"))

	fmt.Println(isZeroFloat(0.0))
	fmt.Println(isZeroFloat(3.14))
}