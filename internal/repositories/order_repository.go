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
	err := r.connection.WithContext(ctx).Create(&order).Error
	return err
}

func (r OrderRepository) List(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	result := r.connection.WithContext(ctx).Find(&orders)
	return orders, result.Error
}

func (r OrderRepository) updateFieldByOrderID(ctx context.Context, orderID int, field string, newValue interface{}) error {
	return r.connection.WithContext(ctx).Where("order_id = ?", orderID).Update(field, newValue).Error
}

func (r OrderRepository) UpdateStatusByOrderId(ctx context.Context, orderID int, newStatus status.Status) error {
	return r.updateFieldByOrderID(ctx, orderID, "status", newStatus)
}
