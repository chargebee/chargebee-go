package gatewaypaymentmethodtoken

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	gatewayPaymentMethodTokenEnum "github.com/chargebee/chargebee-go/v3/models/gatewaypaymentmethodtoken/enum"
)

type GatewayPaymentMethodToken struct {
	Id                string                               `json:"id"`
	GatewayAccountId  string                               `json:"gateway_account_id"`
	GatewayName       enum.GatewayName                     `json:"gateway_name"`
	GatewayCustomerId string                               `json:"gateway_customer_id"`
	GatewayToken      string                               `json:"gateway_token"`
	Status            gatewayPaymentMethodTokenEnum.Status `json:"status"`
	CreatedAt         int64                                `json:"created_at"`
	UpdatedAt         int64                                `json:"updated_at"`
	Object            string                               `json:"object"`
}
