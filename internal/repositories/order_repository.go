package repositories

import (
	"orders/internal/models"
)

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

type OrderRepository struct {
}

func (r OrderRepository) Create(order *models.Order) {
}

func (r OrderRepository) List() []*models.Order {
	return []*models.Order{}
}
