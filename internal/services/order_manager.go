package services

import (
	"orders/internal/models"
)

type OrderRepository interface {
	List() []*models.Order
	Create(order *models.Order)
}

type OrderObserver interface {
	Notify(order *models.Order)
}

func NewOrderManager(repository OrderRepository, observers []OrderObserver) *OrderManager {
	return &OrderManager{repository, observers}
}

type OrderManager struct {
	repository OrderRepository
	observers  []OrderObserver
}

func (m *OrderManager) List() []*models.Order {
	return m.repository.List()
}

func (m *OrderManager) Create(order *models.Order) error {
	m.notifyObservers(order)
	return nil
}

func (m *OrderManager) notifyObservers(order *models.Order) {
	for _, observer := range m.observers {
		observer.Notify(order)
	}
}
