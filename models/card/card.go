package card

import (
	"github.com/chargebee/chargebee-go/enum"
	cardEnum "github.com/chargebee/chargebee-go/models/card/enum"
)

type Card struct {
	PaymentSourceId  string               `json:"payment_source_id"`
	Status           cardEnum.Status      `json:"status"`
	Gateway          enum.Gateway         `json:"gateway"`
	GatewayAccountId string               `json:"gateway_account_id"`
	RefTxId          string               `json:"ref_tx_id"`
	FirstName        string               `json:"first_name"`
	LastName         string               `json:"last_name"`
	Iin              string               `json:"iin"`
	Last4            string               `json:"last4"`
	CardType         cardEnum.CardType    `json:"card_type"`
	FundingType      cardEnum.FundingType `json:"funding_type"`
	ExpiryMonth      int32                `json:"expiry_month"`
	ExpiryYear       int32                `json:"expiry_year"`
	IssuingCountry   string               `json:"issuing_country"`
	BillingAddr1     string               `json:"billing_addr1"`
	BillingAddr2     string               `json:"billing_addr2"`
	BillingCity      string               `json:"billing_city"`
	BillingStateCode string               `json:"billing_state_code"`
	BillingState     string               `json:"billing_state"`
	BillingCountry   string               `json:"billing_country"`
	BillingZip       string               `json:"billing_zip"`
	CreatedAt        int64                `json:"created_at"`
	ResourceVersion  int64                `json:"resource_version"`
	UpdatedAt        int64                `json:"updated_at"`
	IpAddress        string               `json:"ip_address"`
	PoweredBy        cardEnum.PoweredBy   `json:"powered_by"`
	CustomerId       string               `json:"customer_id"`
	MaskedNumber     string               `json:"masked_number"`
	Object           string               `json:"object"`
}
type UpdateCardForCustomerRequestParams struct {
	Gateway          enum.Gateway                         `json:"gateway,omitempty"`
	GatewayAccountId string                               `json:"gateway_account_id,omitempty"`
	TmpToken         string                               `json:"tmp_token,omitempty"`
	FirstName        string                               `json:"first_name,omitempty"`
	LastName         string                               `json:"last_name,omitempty"`
	Number           string                               `json:"number"`
	ExpiryMonth      *int32                               `json:"expiry_month"`
	ExpiryYear       *int32                               `json:"expiry_year"`
	Cvv              string                               `json:"cvv,omitempty"`
	BillingAddr1     string                               `json:"billing_addr1,omitempty"`
	BillingAddr2     string                               `json:"billing_addr2,omitempty"`
	BillingCity      string                               `json:"billing_city,omitempty"`
	BillingStateCode string                               `json:"billing_state_code,omitempty"`
	BillingState     string                               `json:"billing_state,omitempty"`
	BillingZip       string                               `json:"billing_zip,omitempty"`
	BillingCountry   string                               `json:"billing_country,omitempty"`
	IpAddress        string                               `json:"ip_address,omitempty"`
	Customer         *UpdateCardForCustomerCustomerParams `json:"customer,omitempty"`
}
type UpdateCardForCustomerCustomerParams struct {
	VatNumber string `json:"vat_number,omitempty"`
}
type SwitchGatewayForCustomerRequestParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id"`
}
type CopyCardForCustomerRequestParams struct {
	GatewayAccountId string `json:"gateway_account_id"`
}
