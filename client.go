package chargebee

type Client struct {
	config *ClientConfig

	Addon   *AddonService
	Address *AddressService
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		config:  config,
		Addon:   &AddonService{config: config},
		Address: &AddressService{config: config},
	}
}
