package chargebee

// just struct
type ThirdPartyPaymentMethod struct {
	Type             Type    `json:"type"`
	Gateway          Gateway `json:"gateway"`
	GatewayAccountId string  `json:"gateway_account_id"`
	ReferenceId      string  `json:"reference_id"`
	Object           string  `json:"object"`
}

// sub resources
// operations
// input params
