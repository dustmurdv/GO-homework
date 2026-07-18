package main

import "fmt"

func simulateConnections(maxConnections int) (int, int) {

	accepted := 0
	rejected := 0

	for i := 1; i <= 10; i++ {

		if accepted < maxConnections {
			accepted++
		} else {
			rejected++
		}
	}

	return accepted, rejected
}

func main() {

	a, r := simulateConnections(5)

	fmt.Println("Accepted:", a)
	fmt.Println("Rejected:", r)
}