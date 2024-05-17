package dto

import "errors"

type Order struct {
	CustomerInfo CustomerInfo   `json:"customer_info"`
	Products     []OrderProduct `json:"products"`
}

type CustomerInfo struct {
	Name            string `json:"name"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
	Email           string `json:"email,omitempty"`
}

func (o *Order) Validate() error {
	if len(o.Products) == 0 {
		return errors.New("products list is empty")
	}
	return nil
}
