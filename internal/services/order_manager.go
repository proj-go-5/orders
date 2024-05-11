package services

import (
	"context"
	"errors"
	"math/rand/v2"
	"orders/internal/enums/status"
	"orders/internal/models"
)

var (
	ErrEntityNotValid = errors.New("order validation failed")
)

type OrderRepository interface {
	List(ctx context.Context) ([]models.Order, error)
	Create(ctx context.Context, order *models.Order) error
}

type OrderProductRepository interface {
	GetAllByOrderId(ctx context.Context, orderID int) ([]models.OrderProduct, error)
	Create(ctx context.Context, orderProduct *models.OrderProduct) error
}

func NewOrderManager(orderRepo OrderRepository, orderProductRepo OrderProductRepository) *OrderManager {
	return &OrderManager{orderRepo, orderProductRepo}
}

type OrderManager struct {
	orderRepo         OrderRepository
	orderProductsRepo OrderProductRepository
}

func (m *OrderManager) List(ctx context.Context) ([]models.Order, error) {
	orders, err := m.orderRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	for i := range orders {
		products, err := m.orderProductsRepo.GetAllByOrderId(ctx, orders[i].ID)
		if err != nil {
			return nil, err
		}

		orders[i].Products = products
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

	err = m.orderRepo.Create(ctx, order)
	if err != nil {
		return err
	}

	for _, product := range order.Products {
		product.OrderID = order.ID
		productPrice, err := m.getProductPrice()
		if err != nil {
			return err
		}
		product.Price = productPrice
		err = m.orderProductsRepo.Create(ctx, &product)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *OrderManager) getTotalPrice() (int, error) {
	return rand.IntN(300) + 200, nil
}

func (m *OrderManager) getProductPrice() (int, error) {
	return rand.IntN(50) + 50, nil
}
