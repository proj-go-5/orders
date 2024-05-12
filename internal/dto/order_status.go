package dto

import "orders/internal/enums/status"

type OrderStatus struct {
	Status  status.Status `json:"status"`
	Comment string        `json:"comment"`
}
