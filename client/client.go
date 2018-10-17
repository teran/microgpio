package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/teran/microgpio/models"
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
	return c.request("GET", "/ping", nil)
}

// On turnes the pin on
func (c *Client) On(name string) error {
	return c.request("POST", fmt.Sprintf("/pin/%s/on", name), nil)
}

// Off turnes the pin off
func (c *Client) Off(name string) error {
	return c.request("POST", fmt.Sprintf("/pin/%s/off", name), nil)
}

// Status returns the pin status
func (c *Client) Status(name string) (models.ToggleStatus, error) {
	var st models.ToggleStatus
	err := c.request("GET", fmt.Sprintf("/pin/%s", name), &st)
	return st, err
}

func (c *Client) request(method, uri string, data interface{}) error {
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

	if response != nil {
		err := json.NewDecoder(response.Body).Decode(data)
		return err
	}

	return nil
}
