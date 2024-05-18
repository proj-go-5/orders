package models

import (
	"orders/internal/enums/status"
	"time"
)

type OrderHistory struct {
	OrderID   int           `json:"order_id"`
	Status    status.Status `json:"status"`
	Comment   string        `json:"comment"`
	CreatedAt time.Time     `json:"created_at"`
}

func (OrderHistory) TableName() string {
	return "order_history"
}
