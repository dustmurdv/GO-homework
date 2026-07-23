package main

import "fmt"

type Product struct {
	id       string
	name     string
	price    float64
	stock    int
	category string
}

func NewProduct(id, name string, price float64, stock int, category string) *Product {
	return &Product{
		id:       id,
		name:     name,
		price:    price,
		stock:    stock,
		category: category,
	}
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) Category() string {
	return p.category
}

func (p *Product) Reserve(qty int) error {
	if qty > p.stock {
		return fmt.Errorf("not enough stock")
	}

	p.stock -= qty
	return nil
}