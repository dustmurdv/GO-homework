package main

import (
	"errors"
	"fmt"
)

func fibonacci(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("number cannot be negative")
	}

	if n == 0 {
		return 0, nil
	}

	if n == 1 {
		return 1, nil
	}

	a := 0
	b := 1

	for i := 2; i <= n; i++ {

		temp := a + b
		a = b
		b = temp
	}

	return b, nil
}

func main() {

	nums := []int{0, 1, 2, 3, 4, 5, 10}

	for _, n := range nums {

		answer, _ := fibonacci(n)

		fmt.Println(n, "=", answer)
	}
}