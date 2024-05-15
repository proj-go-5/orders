package product

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	http *http.Client
	host string
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) getAbsoluteURL(path string, params url.Values) string {
	return fmt.Sprintf("%s/%s?%s", c.host, path, params.Encode())
}

func NewClient(httpClient *http.Client, host string) *Client {
	return &Client{httpClient, host}
}
