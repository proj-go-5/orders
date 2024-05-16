package dto

import "errors"

type Order struct {
	CustomerInfo CustomerInfo   `json:"customer_info"`
	Products     []OrderProduct `json:"products"`
}

func (o *Order) Validate() error {
	if len(o.Products) == 0 {
		return errors.New("products list is empty")
	}
	return nil
}
