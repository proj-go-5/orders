package services

import (
	"orders/internal/models"
)

type OrderRepository interface {
	List() []*models.Order
	Create(order *models.Order)
}

func NewOrderManager(repository OrderRepository) *OrderManager {
	return &OrderManager{repository}
}

type OrderManager struct {
	repository OrderRepository
}

func (m *OrderManager) List() []*models.Order {
	return m.repository.List()
}

func (m *OrderManager) Create(order *models.Order) error {
	return nil
}
