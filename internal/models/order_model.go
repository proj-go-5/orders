package models

import (
	"orders/internal/enums/status"
	"time"
)

type Order struct {
	ID           int            `json:"id"`
	Status       status.Status  `json:"status"`
	TotalPrice   int            `json:"total_price"`
	CustomerInfo CustomerInfo   `json:"customer_info" gorm:"serializer:json"`
	CreatedAt    time.Time      `json:"created_at"`
	Products     []OrderProduct `json:"products"`
}

type CustomerInfo struct {
	Name            string `json:"name"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
	Email           string `json:"email,omitempty"`
}
