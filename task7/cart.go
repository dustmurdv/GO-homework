package main

type Cart struct {
	items []Orderable
}

func NewCart() *Cart {
	return &Cart{}
}

func (c *Cart) Add(item Orderable) {
	c.items = append(c.items, item)
}