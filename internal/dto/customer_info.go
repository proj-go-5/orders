package dto

type CustomerInfo struct {
	Name            string `json:"name"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
	Email           string `json:"email,omitempty"`
}
