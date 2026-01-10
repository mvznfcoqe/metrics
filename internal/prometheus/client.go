package prometheus

import (
	"fmt"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type Client struct {
	api v1.API
}

func NewClient(prometheusURL string) (*Client, error) {
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create prometheus client: %w", err)
	}

	return &Client{
		api: v1.NewAPI(client),
	}, nil
}

func (c *Client) API() v1.API {
	return c.api
}
