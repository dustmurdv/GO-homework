package main

import "fmt"

type DigitalProduct struct {
	Product

	downloadURL string
	fileSizeMB  float64
}

func NewDigitalProduct(
	id string,
	name string,
	price float64,
	category string,
	url string,
	size float64,
) *DigitalProduct {

	p := &DigitalProduct{
		Product: Product{
			id:       id,
			name:     name,
			price:    price,
			stock:    0,
			category: category,
		},
		downloadURL: url,
		fileSizeMB:  size,
	}

	return p
}

// Digital product is always available.
func (d *DigitalProduct) Reserve(qty int) error {
	return nil
}

func (d DigitalProduct) DownloadLink(customerID string) string {
	return fmt.Sprintf("%s?customer=%s", d.downloadURL, customerID)
}