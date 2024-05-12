package repositories

import (
	"context"
	"orders/internal/models"

	"gorm.io/gorm"
)

func NewOrderProductRepository(connection *gorm.DB) *OrderProductRepository {
	return &OrderProductRepository{connection}
}

type OrderProductRepository struct {
	connection *gorm.DB
}

func (o OrderProductRepository) Create(ctx context.Context, product *models.OrderProduct) error {
	return o.connection.WithContext(ctx).Create(&product).Error
}

func (o OrderProductRepository) GetAllByOrderId(ctx context.Context, orderID int) ([]models.OrderProduct, error) {
	var products []models.OrderProduct
	result := o.connection.WithContext(ctx).Find(&products, "order_id = ?", orderID)
	return products, result.Error
}
