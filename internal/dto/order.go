package dto

type Order struct {
	CustomerID int            `json:"customer_id"`
	Products   []OrderProduct `json:"products"`
}
