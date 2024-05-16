package repositories

import (
	"context"
	"orders/internal/models"

	"gorm.io/gorm"
)

type OrderProductRepository interface {
	Create(ctx context.Context, tx *gorm.DB, orderProduct *models.OrderProduct) error
}

func NewOrderProductRepository(connection *gorm.DB) OrderProductRepository {
	return &orderProductRepository{connection}
}

type orderProductRepository struct {
	connection *gorm.DB
}

func (o orderProductRepository) Create(ctx context.Context, tx *gorm.DB, product *models.OrderProduct) error {
	return tx.WithContext(ctx).Create(&product).Error
}
