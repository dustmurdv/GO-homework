package main

import (
	"errors"
	"fmt"
)

func power(base int, exp int) (int, error) {

	if exp < 0 {
		return 0, errors.New("negative exponent")
	}

	result := 1

	for i := 0; i < exp; i++ {
		result = result * base
	}

	return result, nil
}

func main() {

	a, _ := power(2, 10)
	fmt.Println(a)

	b, _ := power(3, 0)
	fmt.Println(b)

	c, err := power(2, -1)
	fmt.Println(c, err)
}