package main

func Checkout(items []Orderable) (float64, error) {

	var total float64

	for _, item := range items {

		err := item.Reserve(1)
		if err != nil {
			return 0, err
		}

		total += item.Price()
	}

	return total, nil
}

func ProcessOrder(cart *Cart, customerID string, processor PaymentProcessor) error {

	total, err := Checkout(cart.items)
	if err != nil {
		return err
	}

	return processor.Charge(total, customerID)
}