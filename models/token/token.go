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
	CreatedAt         int64                  `json:"created_at"`
	ExpiredAt         int64                  `json:"expired_at"`
	Object            string                 `json:"object"`
}
type CreateForCardCardParams struct {
	GatewayAccountId string `json:"gateway_account_id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Number           string `json:"number"`
	ExpiryMonth      *int32 `json:"expiry_month"`
	ExpiryYear       *int32 `json:"expiry_year"`
	Cvv              string `json:"cvv,omitempty"`
	BillingAddr1     string `json:"billing_addr1,omitempty"`
	BillingAddr2     string `json:"billing_addr2,omitempty"`
	BillingCity      string `json:"billing_city,omitempty"`
	BillingStateCode string `json:"billing_state_code,omitempty"`
	BillingState     string `json:"billing_state,omitempty"`
	BillingZip       string `json:"billing_zip,omitempty"`
	BillingCountry   string `json:"billing_country,omitempty"`
}
type CreateUsingTempTokenTokenBillingAddressParams struct {
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	StateCode   string `json:"state_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Zip         string `json:"zip,omitempty"`
}
