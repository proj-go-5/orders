package repositories

import (
	"context"
	"orders/internal/models"

	"gorm.io/gorm"
)

func NewOrderHistoryRepository(connection *gorm.DB) *OrderHistoryRepository {
	return &OrderHistoryRepository{connection}
}

type OrderHistoryRepository struct {
	connection *gorm.DB
}

func (h OrderHistoryRepository) Create(ctx context.Context, record *models.OrderHistory) error {
	return h.connection.WithContext(ctx).Create(&record).Error
}

func (h OrderHistoryRepository) List(ctx context.Context, orderID int) ([]models.OrderHistory, error) {
	var history []models.OrderHistory
	result := h.connection.WithContext(ctx).Find(&history, "order_id = ?", orderID)
	return history, result.Error
}
