package main

import (
	"errors"
	"strings"
)

type Customer struct {
	id      int
	name    string
	email   string
	orders  []string
}

func NewCustomer(id int, name string, email string) (*Customer, error) {

	if !strings.Contains(email, "@") {
		return nil, errors.New("invalid email")
	}

	customer := &Customer{
		id:     id,
		name:   name,
		email:  email,
		orders: []string{},
	}

	return customer, nil
}

func (c Customer) ID() int {
	return c.id
}

func (c Customer) Name() string {
	return c.name
}

func (c Customer) Email() string {
	return c.email
}

func (c Customer) Orders() []string {
	return c.orders
}

func (c *Customer) AddOrder(order string) {

	c.orders = append(c.orders, order)
}