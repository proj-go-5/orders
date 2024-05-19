package product

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"orders/internal/config"
)

type Client struct {
	http *http.Client
	addr string
}

func (c *Client) Get(ctx context.Context, path string, params url.Values) ([]byte, error) {
	absoluteURL := c.getAbsoluteURL(path, params)
	request, err := http.NewRequestWithContext(ctx, "GET", absoluteURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) getAbsoluteURL(path string, params url.Values) string {
	return fmt.Sprintf("%s%s?%s", c.addr, path, params.Encode())
}

func NewClient(httpClient *http.Client) *Client {
	addr := fmt.Sprintf(
		"%s://%s",
		config.Env("PRODUCT_CATALOG_ADDR_SCHEME"),
		config.Env("PRODUCT_CATALOG_SERVICE_ADDR"),
	)

	return &Client{httpClient, addr}
}
