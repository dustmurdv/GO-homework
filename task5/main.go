package main

import "fmt"

type BundleProduct struct {
	Product
	items    []*Product
	discount float64
}

func NewBundleProduct(id, name string, discount float64) *BundleProduct {
	return &BundleProduct{
		Product: Product{
			id:   id,
			name: name,
		},
		items: []*Product{},
		discount: discount,
	}
}

func (b *BundleProduct) AddItem(p *Product) {
	b.items = append(b.items, p)
}

func (b *BundleProduct) Price() float64 {
	var total float64

	for _, item := range b.items {
		total += item.Price()
	}

	total = total - total*b.discount/100

	return total
}

func (b *BundleProduct) Reserve(qty int) error {
	for _, item := range b.items {
		err := item.Reserve(qty)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	keyboard := NewProduct("P1", "Keyboard", 250000, 10, "PC")
	mouse := NewProduct("P2", "Mouse", 150000, 20, "PC")

	bundle := NewBundleProduct("B1", "Gaming Set", 10)

	bundle.AddItem(keyboard)
	bundle.AddItem(mouse)

	fmt.Println("Bundle:", bundle.Name())
	fmt.Println("Price:", bundle.Price())

	err := bundle.Reserve(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Reserved successfully")
	}
}