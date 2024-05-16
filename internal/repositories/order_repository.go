package repositories

import (
	"context"
	"orders/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB) error
	RollbackTransaction(tx *gorm.DB) error
	List(ctx context.Context) ([]models.Order, error)
	// Create(ctx context.Context, tx *gorm.DB, order *models.Order) error
	Create(ctx context.Context, order *models.Order) error
}

func NewOrderRepository(connection *gorm.DB) OrderRepository {
	return &orderRepository{connection}
}

type orderRepository struct {
	connection *gorm.DB
}

func (r *orderRepository) BeginTransaction() *gorm.DB {
	return r.connection.Begin()
}

func (r *orderRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *orderRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

// func (r orderRepository) Create(ctx context.Context, tx *gorm.DB, order *models.Order) error {
// 	return tx.WithContext(ctx).Create(order).Error
// }

func (r orderRepository) Create(ctx context.Context, order *models.Order) error {
	return r.connection.WithContext(ctx).Create(&order).Error
	// r.connection.Save(&order)
	// return nil
}

func (r orderRepository) List(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.connection.WithContext(ctx).Model(&models.Order{}).Preload("Products").Find(&orders).Error
	return orders, err
}
