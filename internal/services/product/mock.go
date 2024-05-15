package product

import (
	"context"
	"math/rand/v2"
	"orders/internal/dto"
)

type MockFetcher struct {
}

func (f *MockFetcher) GetCatalogProducts(ctx context.Context, filter *Filter) ([]*dto.CatalogProduct, error) {
	var products []*dto.CatalogProduct

	for _, ID := range filter.IDs {
		price, err := f.randomPrice()
		if err != nil {
			return nil, err
		}

		products = append(products, &dto.CatalogProduct{
			ID:    ID,
			Price: price,
		})
	}

	return products, nil
}

func (f *MockFetcher) randomPrice() (int, error) {
	return rand.IntN(50) + 50, nil
}

func NewMockFetcher() *MockFetcher {
	return &MockFetcher{}
}
