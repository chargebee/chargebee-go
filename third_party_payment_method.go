package chargebee

type ThirdPartyPaymentMethod struct {
	Type             enum.Type    `json:"type"`
	Gateway          enum.Gateway `json:"gateway"`
	GatewayAccountId string       `json:"gateway_account_id"`
	ReferenceId      string       `json:"reference_id"`
	Object           string       `json:"object"`
}
