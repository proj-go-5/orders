package repositories

import (
	"context"
	"orders/internal/models"

	"gorm.io/gorm"
)

func NewOrderRepository(connection *gorm.DB) *OrderRepository {
	return &OrderRepository{connection}
}

type OrderRepository struct {
	connection *gorm.DB
}

func (r OrderRepository) Create(ctx context.Context, order *models.Order) error {
	result := r.connection.WithContext(ctx).Create(&order)
	return result.Error
}

func (r OrderRepository) List(ctx context.Context) ([]*models.Order, error) {
	var orders []*models.Order
	result := r.connection.WithContext(ctx).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
