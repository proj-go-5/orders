package repositories

import (
	"gorm.io/gorm"
	"orders/internal/models"
)

func NewOrderRepository(connection *gorm.DB) *OrderRepository {
	return &OrderRepository{connection}
}

type OrderRepository struct {
	connection *gorm.DB
}

func (r OrderRepository) Create(order *models.Order) {
}

func (r OrderRepository) List() []*models.Order {
	return []*models.Order{}
}
