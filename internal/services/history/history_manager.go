package history

import (
	"context"
	"orders/internal/models"
)

type HistoryRepository interface {
	Create(ctx context.Context, record *models.OrderHistory) error
	List(ctx context.Context, orderID int) ([]models.OrderHistory, error)
}

func NewOrderHistoryManager(historyRepo HistoryRepository) *HistoryManager {
	return &HistoryManager{historyRepo}
}

type HistoryManager struct {
	historyRepo HistoryRepository
}

func (h *HistoryManager) Create(ctx context.Context, historyRecord *models.OrderHistory) error {
	return h.historyRepo.Create(ctx, historyRecord)
}

func (h *HistoryManager) List(ctx context.Context, orderID int) ([]models.OrderHistory, error) {
	return h.historyRepo.List(ctx, orderID)
}
