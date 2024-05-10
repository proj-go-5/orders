package models

import (
	"orders/internal/enums/status"
	"time"
)

type Order struct {
	ID         int           `json:"id"`
	Status     status.Status `json:"status"`
	TotalPrice int           `json:"total_price"`
	CustomerID int           `json:"customer_id"`
	CreatedAt  time.Time     `json:"created_at"`
}
