package thirdpartypaymentmethod

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/v3/enum"
)

type ThirdPartyPaymentMethod struct {
	Type                        enum.Type       `json:"type"`
	Gateway                     enum.Gateway    `json:"gateway"`
	GatewayAccountId            string          `json:"gateway_account_id"`
	ReferenceId                 string          `json:"reference_id"`
	NetworkTransactionReference json.RawMessage `json:"network_transaction_reference"`
	Object                      string          `json:"object"`
}
