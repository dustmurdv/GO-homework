package main

import "errors"

type Cart struct {
	products []Product
}

func NewCart() *Cart {
	return &Cart{
		products: []Product{},
	}
}

func (c *Cart) AddProduct(product Product) {
	c.products = append(c.products, product)
}

func (c *Cart) RemoveProduct(id int) error {

	for i := 0; i < len(c.products); i++ {

		if c.products[i].ID() == id {

			c.products = append(c.products[:i], c.products[i+1:]...)
			return nil
		}
	}

	return errors.New("product not found")
}

func (c *Cart) Clear() {
	c.products = []Product{}
}

func (c Cart) TotalItems() int {
	return len(c.products)
}

func (c Cart) Products() []Product {
	return c.products
}