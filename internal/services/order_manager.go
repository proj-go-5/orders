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

func NewOrderManager(orderRepo repositories.OrderRepository, orderProductRepo repositories.OrderProductRepository) *OrderManager {
	return &OrderManager{orderRepo, orderProductRepo}
}

type OrderManager struct {
	orderRepo         repositories.OrderRepository
	orderProductsRepo repositories.OrderProductRepository
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

	for i := range order.OrderProducts {
		productPrice, err := m.getProductPrice()
		if err != nil {
			return err
		}
		order.OrderProducts[i].Price = productPrice
	}

	err = m.orderRepo.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

// func (m *OrderManager) Create(ctx context.Context, order *models.Order) error {
// 	totalPrice, err := m.getTotalPrice()
// 	if err != nil {
// 		return err
// 	}
// 	order.TotalPrice = totalPrice
// 	order.Status = status.Active

// 	tx := m.orderRepo.BeginTransaction()
// 	err = m.orderRepo.Create(ctx, tx, order)
// 	if err != nil {
// 		m.orderRepo.RollbackTransaction(tx)
// 		return err
// 	}

// 	for _, product := range order.Products {
// 		product.OrderID = order.ID
// 		productPrice, err := m.getProductPrice()
// 		if err != nil {
// 			return err
// 		}
// 		product.Price = productPrice
// 		err = m.orderProductsRepo.Create(ctx, tx, &product)
// 		if err != nil {
// 			m.orderRepo.RollbackTransaction(tx)
// 			return err
// 		}
// 	}

// 	return m.orderRepo.CommitTransaction(tx)
// }

func (m *OrderManager) getTotalPrice() (int, error) {
	return rand.IntN(300) + 200, nil
}

func (m *OrderManager) getProductPrice() (int, error) {
	return rand.IntN(50) + 50, nil
}
