package chargebee

import "net/http"

type Transport struct {
	apiKey     string
	apiUrl     string
	httpClient *http.Client
}

func NewTransport(apiKey string) *Transport {
	return &Transport{
		apiKey: apiKey,
		apiUrl: "https://api.chargebee.com",
	}
}

type Client struct {
	transport *Transport
}

func NewClient(apiKey string) *Client {
	return &Client{
		transport: NewTransport(apiKey),
	}
}
