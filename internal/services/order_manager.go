package services

import (
	"context"
	"errors"
	"math/rand/v2"
	"orders/internal/enums/status"
	"orders/internal/models"
	"orders/internal/repositories"
)

var (
	ErrEntityNotValid = errors.New("order validation failed")
)

func NewOrderManager(orderRepo repositories.OrderRepository) *OrderManager {
	return &OrderManager{orderRepo}
}

type OrderManager struct {
	orderRepo repositories.OrderRepository
}

func (m *OrderManager) List(ctx context.Context) ([]models.Order, error) {
	orders, err := m.orderRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (m *OrderManager) Create(ctx context.Context, order *models.Order) error {
	totalPrice, err := m.getTotalPrice()
	if err != nil {
		return err
	}
	order.TotalPrice = totalPrice
	order.Status = status.Active

	for i := range order.Products {
		productPrice, err := m.getProductPrice()
		if err != nil {
			return err
		}
		order.Products[i].Price = productPrice
	}

	return m.orderRepo.Create(ctx, order)
}

func (m *OrderManager) getTotalPrice() (int, error) {
	return rand.IntN(300) + 200, nil
}

func (m *OrderManager) getProductPrice() (int, error) {
	return rand.IntN(50) + 50, nil
}
