package repositories

import (
	"context"
	"orders/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	List(ctx context.Context) ([]models.Order, error)
	Create(ctx context.Context, order *models.Order) error
}

func NewOrderRepository(connection *gorm.DB) OrderRepository {
	return &orderRepository{connection}
}

type orderRepository struct {
	connection *gorm.DB
}

func (r orderRepository) Create(ctx context.Context, order *models.Order) error {
	return r.connection.WithContext(ctx).Create(&order).Error
}

func (r orderRepository) List(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.connection.WithContext(ctx).Model(&models.Order{}).Preload("Products").Find(&orders).Error
	return orders, err
}
