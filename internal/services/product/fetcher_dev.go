//go:build dev

package product

import (
	"context"
	"math/rand/v2"
	"orders/internal/dto"
)

type MockFetcher struct {
}

func (f *MockFetcher) GetProducts(ctx context.Context, filter *Filter) ([]*dto.Product, error) {
	var products []*dto.Product

	for _, ID := range filter.IDs {
		price, err := f.randomPrice()
		if err != nil {
			return nil, err
		}

		products = append(products, &dto.Product{
			ID:    ID,
			Price: price,
		})
	}

	return products, nil
}

func (f *MockFetcher) randomPrice() (int, error) {
	return rand.IntN(50) + 50, nil
}

func NewFetcher(client *Client) *MockFetcher {
	return &MockFetcher{}
}
