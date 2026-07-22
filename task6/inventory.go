package main

type Inventory struct {
	products map[string]*Product
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]*Product),
	}
}

func (i *Inventory) AddProduct(p *Product) {
	i.products[p.ID()] = p
}

func (i *Inventory) FindByID(id string) (*Product, error) {

	product, ok := i.products[id]

	if !ok {
		return nil, nil
	}

	return product, nil
}

func (i *Inventory) FindByCategory(cat string) []*Product {

	var result []*Product

	for _, product := range i.products {

		if product.Category() == cat {
			result = append(result, product)
		}
	}

	return result
}

func (i *Inventory) LowStock(threshold int) []*Product {

	var result []*Product

	for _, product := range i.products {

		if product.Stock() <= threshold {
			result = append(result, product)
		}
	}

	return result
}

func (i *Inventory) TotalValue() float64 {

	total := 0.0

	for _, product := range i.products {
		total += product.Price() * float64(product.Stock())
	}

	return total
}