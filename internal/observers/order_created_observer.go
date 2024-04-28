package observers

import (
	"fmt"
	"orders/internal/models"
)

type EmailObserver struct {
}

func (o *EmailObserver) Notify(order *models.Order) {
	fmt.Printf("Email was sent for user #%d", order.CustomerID)
}

func NewEmailObserver() *EmailObserver {
	return &EmailObserver{}
}
