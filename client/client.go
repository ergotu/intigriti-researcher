package intigriti

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	AuthToken  string
}

func New(authToken string) *Client {
	return &Client{
		BaseURL:    "https://api.intigriti.com/external/researcher",
		HTTPClient: &http.Client{},
		AuthToken:  authToken,
	}
}
