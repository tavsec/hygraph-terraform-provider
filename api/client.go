package hygraph

import (
	"net/http"
	"time"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	AuthToken  string
}

func NewClient(host, auth_token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
	}

	c.AuthToken = *auth_token

	return &c, nil
}

type webhook struct {
	ID   string
	Name string
}

type webhooks []webhook

func (c *Client) GetWebhooks() (webhooks, error) {
	// TODO: Implement
	return append(make(webhooks, 0), webhook{ID: "abc", Name: "test"}), nil
}
