package chargebee

import (
	"net/http"
)

type Transport struct {
	apiKey     string
	apiUrl     string
	httpClient *http.Client
}

func NewTransport(apiKey string) *Transport {
	return &Transport{
		apiKey:     apiKey,
		apiUrl:     "https://api.chargebee.com",
		httpClient: NewDefaultHTTPClient(),
	}
}

type Client struct {
	transport *Transport

	Addon *AddonService
}

func NewClient(apiKey string) *Client {
	transport := NewTransport(apiKey)
	return &Client{
		transport: transport,
		Addon:     &AddonService{transport: transport},
	}
}
