package order

import (
	"context"
	"fmt"
	"orders/internal/dto"
	"orders/internal/enums/status"
	"orders/internal/models"
	"orders/internal/services/history"
	"orders/internal/services/product"
	"time"
)

type Repository interface {
	List(ctx context.Context) ([]models.Order, error)
	Create(ctx context.Context, order *models.Order) error
	UpdateStatus(ctx context.Context, orderID int, newStatus status.Status) error
}

type ProductRepository interface {
	GetAllByOrderId(ctx context.Context, orderID int) ([]models.OrderProduct, error)
	Create(ctx context.Context, orderProduct *models.OrderProduct) error
}

type ProductFetcher interface {
	GetCatalogProducts(ctx context.Context, filter *product.Filter) ([]*dto.CatalogProduct, error)
}

func NewOrderManager(orderRepo Repository, orderProductRepo ProductRepository, historyRepo history.HistoryRepository, productFetcher ProductFetcher) *Manager {
	return &Manager{orderRepo, orderProductRepo, historyRepo, productFetcher}
}

type Manager struct {
	orderRepo         Repository
	orderProductsRepo ProductRepository
	historyRepo       history.HistoryRepository
	productFetcher    ProductFetcher
}

func (m *Manager) List(ctx context.Context) ([]models.Order, error) {
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

func (m *Manager) Create(ctx context.Context, order *models.Order) error {
	productIDs := m.GetProductIDs(order)

	catalogProducts, err := m.productFetcher.GetCatalogProducts(ctx, &product.Filter{IDs: productIDs})
	if err != nil {
		return err
	}

	totalPrice, priceProducts := m.getPrices(catalogProducts)

	order.TotalPrice = totalPrice
	order.Status = status.Active

	err = m.orderRepo.Create(ctx, order)
	if err != nil {
		return err
	}

	for _, orderProduct := range order.Products {
		orderProduct.OrderID = order.ID
		orderProduct.Price = priceProducts[orderProduct.ProductID]

		err = m.orderProductsRepo.Create(ctx, &orderProduct)
		if err != nil {
			return err
		}
	}

	historyRecord := models.OrderHistory{
		OrderID:   order.ID,
		Status:    order.Status,
		Comment:   fmt.Sprintf("Created order %d", order.ID),
		CreatedAt: order.CreatedAt,
	}
	err = m.historyRepo.Create(ctx, &historyRecord)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) UpdateStatus(ctx context.Context, orderID int, newStatus status.Status) error {
	err := m.orderRepo.UpdateStatus(ctx, orderID, newStatus)
	if err != nil {
		return err
	}

	historyRecord := models.OrderHistory{
		OrderID:   orderID,
		Status:    newStatus,
		Comment:   fmt.Sprintf("New status %s in order %d", newStatus, orderID),
		CreatedAt: time.Now(),
	}
	err = m.historyRepo.Create(ctx, &historyRecord)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetProductIDs(order *models.Order) []int {
	var productIDs []int
	for _, orderProduct := range order.Products {
		productIDs = append(productIDs, orderProduct.ProductID)
	}
	return productIDs
}

func (m *Manager) getPrices(catalogProducts []*dto.CatalogProduct) (int, map[int]int) {
	priceProducts := make(map[int]int, len(catalogProducts))
	var totalPrice int

	for _, catalogProduct := range catalogProducts {
		priceProducts[catalogProduct.ID] = catalogProduct.Price
		totalPrice += catalogProduct.Price
	}

	return totalPrice, priceProducts
}
