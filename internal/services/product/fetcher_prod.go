//go:build !dev

package product

import (
	"context"
	"encoding/json"
	"net/url"
	"orders/internal/dto"
	"strconv"
	"strings"
)

type CatalogFetcher struct {
	client *Client
}

func (f *CatalogFetcher) GetProducts(ctx context.Context, filter *Filter) ([]*dto.Product, error) {
	params := url.Values{}
	var stringIDs []string
	for _, id := range filter.IDs {
		stringIDs = append(stringIDs, strconv.Itoa(id))
	}
	params.Add("ids", strings.Join(stringIDs, ","))

	responseData, err := f.client.Get(ctx, "/products", params)
	if err != nil {
		return nil, err
	}

	var response *dto.ListProductsResponse
	err = json.Unmarshal(responseData, &response)
	return response.Products, err
}

func NewFetcher(client *Client) *CatalogFetcher {
	return &CatalogFetcher{client}
}
