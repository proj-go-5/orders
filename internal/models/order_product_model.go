package models

type OrderProduct struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Price     int `json:"price"`
	Quantity  int `json:"quantity"`
}
