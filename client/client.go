package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var (
	// ErrUnexpectedStatusCode returned on unexpected status received from API
	ErrUnexpectedStatusCode = errors.New("unexpected status code received")
)

// Client type
type Client struct {
	endpoint string
}

// New returns new Client instance
func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

// Ping method pings endpoint if it's available
func (c *Client) Ping() error {
	return c.request("GET", "/ping")
}

// Low method sets low bit to the pin with ID provided
func (c *Client) Low(id int) error {
	return c.request("POST", fmt.Sprintf("/gpio/%d/low", id))
}

// High sets high bit to the pin with ID provided
func (c *Client) High(id int) error {
	return c.request("POST", fmt.Sprintf("/gpio/%d/high", id))
}

// Output sets output mode for the pin with ID provided
func (c *Client) Output(id int) error {
	return c.request("POST", fmt.Sprintf("/gpio/%d/output", id))
}

func (c *Client) request(method, uri string) error {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	r, err := http.NewRequest(method, c.endpoint+uri, nil)
	if err != nil {
		return err
	}

	response, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.Wrapf(ErrUnexpectedStatusCode, "status_code=%d", response.StatusCode)
	}
	return nil
}
