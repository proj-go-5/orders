package repositories

import (
	"context"
	"orders/internal/enums/status"
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
	return r.connection.WithContext(ctx).Create(&order).Error
}

func (r OrderRepository) List(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	result := r.connection.WithContext(ctx).Find(&orders)
	return orders, result.Error
}

func (r OrderRepository) UpdateStatus(ctx context.Context, orderID int, newStatus status.Status) error {
	return r.connection.WithContext(ctx).Where("order_id = ?", orderID).Update("status", newStatus).Error
}
