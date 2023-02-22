package hygraph

import (
	"github.com/machinebox/graphql",
	"context"
)

type Client struct {
	HostURL       string
	GraphQlClient *graphql.Client
	AuthToken     string
}

func NewClient(host, auth_token *string) (*Client, error) {
	c := Client{
		GraphQlClient: graphql.NewClient(*host),
		HostURL:       *host,
	}

	c.AuthToken = *auth_token

	return &c, nil
}

func (c *Client) MakeRequest(ctx context.Context, query string, variables map[string]string, responseData interface{}) error {
	req := graphql.NewRequest(query)
	req.Header.Add("Authorization", "Bearer "+c.AuthToken)
	for key, value := range variables {
		req.Var(key, value)
	}

	return c.GraphQlClient.Run(ctx, req, &responseData)
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
