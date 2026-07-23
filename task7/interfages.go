package main

type Purchasable interface {
	ID() string
	Name() string
	Price() float64
	Reserve(qty int) error
}

type Discountable interface {
	ApplyDiscount(percent float64) float64
}

type Orderable interface {
	Purchasable
	Category() string
}