package models

import "time"

type Order struct {
	ID         int `json:"id"`
	Status     int `json:"status"`
	TotalPrice int `json:"total_price"`
	CustomerID int `json:"customer_id"`
	CreatedAt  time.Time
}
