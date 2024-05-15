package product

import (
	"context"
	"encoding/json"
	"net/url"
	"orders/internal/dto"
	"strconv"
)

type Fetcher struct {
	client *Client
}

func (f *Fetcher) GetCatalogProducts(ctx context.Context, filter *Filter) ([]*dto.CatalogProduct, error) {
	var products []*dto.CatalogProduct

	params := url.Values{}
	for _, ID := range filter.IDs {
		params.Add("ids", strconv.Itoa(ID))
	}

	responseData, err := f.client.Get(ctx, "/products", params)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseData, &products)
	return products, err
}

func NewFetcher(client *Client) *Fetcher {
	return &Fetcher{client}
}
