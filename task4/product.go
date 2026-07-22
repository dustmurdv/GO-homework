package main

import "errors"

type Product struct {
	id       string
	name     string
	price    float64
	stock    int
	category string
}

func (p Product) ID() string {
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

func (p Product) Category() string {
	return p.category
}

func (p *Product) SetPrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}

	p.price = price
	return nil
}

func (p *Product) AddStock(qty int) error {
	if qty <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	p.stock += qty
	return nil
}

func (p *Product) Reserve(qty int) error {
	if qty <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	if qty > p.stock {
		return errors.New("not enough stock")
	}

	p.stock -= qty
	return nil
}