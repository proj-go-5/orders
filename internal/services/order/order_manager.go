package order

import (
	"context"
	"fmt"
	"orders/internal/dto"
	"orders/internal/enums/status"
	"orders/internal/mail"
	"orders/internal/models"
	"orders/internal/services/history"
	"orders/internal/services/product"
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
	GetProducts(ctx context.Context, filter *product.Filter) ([]*dto.Product, error)
}

func NewOrderManager(orderRepo Repository, historyRepo history.HistoryRepository, productFetcher ProductFetcher, emailSender *mail.EmailSender) *Manager {
	return &Manager{orderRepo, historyRepo, productFetcher, emailSender}
}

type Manager struct {
	orderRepo      Repository
	historyRepo    history.HistoryRepository
	productFetcher ProductFetcher
	emailSender    *mail.EmailSender
}

func (m *Manager) List(ctx context.Context) ([]models.Order, error) {
	orders, err := m.orderRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (m *Manager) Create(ctx context.Context, order *models.Order) error {
	productIDs := m.GetProductIDs(order)

	catalogProducts, err := m.productFetcher.GetProducts(ctx, &product.Filter{IDs: productIDs})
	if err != nil {
		return err
	}

	totalPrice, priceProducts := m.getPrices(catalogProducts)

	order.TotalPrice = totalPrice
	order.Status = status.Active

	for i, p := range order.Products {
		order.Products[i].OrderID = order.ID
		order.Products[i].Price = priceProducts[p.ProductID]
	}

	err = m.orderRepo.Create(ctx, order)
	if err != nil {
		return err
	}

	historyRecord := models.OrderHistory{
		OrderID:   order.ID,
		Status:    order.Status,
		Comment:   fmt.Sprintf("Created order %d", order.ID),
		CreatedAt: order.CreatedAt,
	}
	err = m.historyRepo.Create(ctx, &historyRecord)
	if err != nil {
		fmt.Println(err)
	}

	err = m.emailSender.SendEmail(order)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (m *Manager) UpdateStatus(ctx context.Context, orderID int, newStatus status.Status, comment string) error {
	err := m.orderRepo.UpdateStatus(ctx, orderID, newStatus)
	if err != nil {
		return err
	}

	historyComment := fmt.Sprintf("New status %s in order %d", newStatus, orderID)
	if comment != "" {
		historyComment = comment
	}

	historyRecord := models.OrderHistory{
		OrderID: orderID,
		Status:  newStatus,
		Comment: historyComment,
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

func (m *Manager) getPrices(catalogProducts []*dto.Product) (int, map[int]int) {
	priceProducts := make(map[int]int, len(catalogProducts))
	var totalPrice int

	for _, catalogProduct := range catalogProducts {
		priceProducts[catalogProduct.ID] = catalogProduct.Price
		totalPrice += catalogProduct.Price
	}

	return totalPrice, priceProducts
}
