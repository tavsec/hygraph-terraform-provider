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
