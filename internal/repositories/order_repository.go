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
	if err == nil {
		_, err = r.UpdateOrderStatusByOrderId(ctx, order.ID, status.Active)
	}
	return err
}

func (r OrderRepository) List(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	result := r.connection.WithContext(ctx).Find(&orders)
	return orders, result.Error
}

func (r OrderRepository) updateOrderFieldByOrderID(ctx context.Context, orderID int, field string, newData interface{}) ([]models.Order, error) {
	updateErr := r.connection.WithContext(ctx).Where("order_id = ?", orderID).Update(field, newData).Error
	if updateErr != nil {
		return nil, updateErr
	}

	var orders []models.Order
	result := r.connection.WithContext(ctx).Find(&orders, "order_id = ?", orderID)
	return orders, result.Error
}

func (r OrderRepository) UpdateOrderStatusByOrderId(ctx context.Context, orderID int, newStatus status.Status) (models.Order, error) {
	result, error := r.updateOrderFieldByOrderID(ctx, orderID, "status", newStatus)
	return result[0], error
}
