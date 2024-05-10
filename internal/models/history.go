package models

import (
	"orders/internal/enums/status"
	"time"
)

type History struct {
	OrderID   int
	Status    status.Status
	Comment   string
	CreatedAt time.Time
}
