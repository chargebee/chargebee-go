package token

import (
	"github.com/chargebee/chargebee-go/enum"
	tokenEnum "github.com/chargebee/chargebee-go/models/token/enum"
)

type Token struct {
	Id                string                 `json:"id"`
	Gateway           enum.Gateway           `json:"gateway"`
	GatewayAccountId  string                 `json:"gateway_account_id"`
	PaymentMethodType enum.PaymentMethodType `json:"payment_method_type"`
	Status            tokenEnum.Status       `json:"status"`
	IdAtVault         string                 `json:"id_at_vault"`
	Vault             tokenEnum.Vault        `json:"vault"`
	IpAddress         string                 `json:"ip_address"`
	ResourceVersion   int64                  `json:"resource_version"`
	UpdatedAt         int64                  `json:"updated_at"`
	CreatedAt         int64                  `json:"created_at"`
	ExpiredAt         int64                  `json:"expired_at"`
	Object            string                 `json:"object"`
}
