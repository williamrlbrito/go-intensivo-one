package entity

import "errors"

type OrderRepositoryInterface interface {
	Save(order *Order) error
}

type Order struct {
	ID         int
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id int, price, tax float64) (*Order, error) {
	order := Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}

func (o Order) IsValid() error {
	if o.ID == 0 {
		return errors.New("invalid ID")
	}
	if o.Price == 0 {
		return errors.New("invalid Price")
	}
	if o.Tax == 0 {
		return errors.New("invalid Tax")
	}
	return nil
}
