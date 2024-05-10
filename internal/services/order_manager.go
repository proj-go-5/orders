package services

import (
	"context"
	"errors"
	"orders/internal/models"
)

var (
	ErrEntityNotValid = errors.New("order validation failed")
)

type OrderRepository interface {
	List(ctx context.Context) ([]*models.Order, error)
	Create(ctx context.Context, order *models.Order) error
}

func NewOrderManager(repository OrderRepository) *OrderManager {
	return &OrderManager{repository}
}

type OrderManager struct {
	repository OrderRepository
}

func (m *OrderManager) List(ctx context.Context) ([]*models.Order, error) {
	return m.repository.List(ctx)
}

func (m *OrderManager) Create(ctx context.Context, order *models.Order) error {
	if order.Status != 0 || order.TotalPrice < 1 {
		return ErrEntityNotValid
	}
	return m.repository.Create(ctx, order)
}
