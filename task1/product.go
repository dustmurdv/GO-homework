package main

import "errors"

type Product struct {
	id    int
	name  string
	price float64
	stock int
}

func NewProduct(id int, name string, price float64, stock int) Product {
	return Product{
		id:    id,
		name:  name,
		price: price,
		stock: stock,
	}
}

func (p Product) ID() int {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() float64 {
	return p.price
}

func (p Product) Stock() int {
	return p.stock
}

func (p *Product) SetPrice(price float64) error {

	if price <= 0 {
		return errors.New("price must be greater than 0")
	}

	p.price = price
	return nil
}

func (p *Product) AddStock(count int) {

	if count > 0 {
		p.stock += count
	}
}

func (p *Product) Reserve(count int) error {

	if count <= 0 {
		return errors.New("invalid quantity")
	}

	if count > p.stock {
		return errors.New("not enough stock")
	}

	p.stock -= count
	return nil
}