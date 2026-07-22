package main

import "fmt"

func main() {

	book := NewDigitalProduct(
		"D1",
		"Go Programming",
		89000,
		"Books",
		"https://download.com/go-book",
		25.5,
	)

	fmt.Println("ID:", book.ID())
	fmt.Println("Name:", book.Name())
	fmt.Println("Price:", book.Price())
	fmt.Println("Category:", book.Category())

	err := book.Reserve(100)

	if err == nil {
		fmt.Println("Reserve successful")
	}

	fmt.Println(book.DownloadLink("AKMAL001"))
}