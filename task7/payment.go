package main

import "fmt"

type PaymentProcessor interface {
	Charge(amount float64, customerID string) error
}

type CardProcessor struct{}

func (c CardProcessor) Charge(amount float64, customerID string) error {
	fmt.Println("Payment successful")
	fmt.Println("Customer:", customerID)
	fmt.Println("Amount:", amount)
	return nil
}	