package virtualbankaccount

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
)

type VirtualBankAccount struct {
	Id               string       `json:"id"`
	CustomerId       string       `json:"customer_id"`
	Email            string       `json:"email"`
	BankName         string       `json:"bank_name"`
	AccountNumber    string       `json:"account_number"`
	RoutingNumber    string       `json:"routing_number"`
	SwiftCode        string       `json:"swift_code"`
	Gateway          enum.Gateway `json:"gateway"`
	GatewayAccountId string       `json:"gateway_account_id"`
	ReferenceId      string       `json:"reference_id"`
	Deleted          bool         `json:"deleted"`
	Object           string       `json:"object"`
}
type CreateUsingPermanentTokenRequestParams struct {
	CustomerId  string `json:"customer_id"`
	ReferenceId string `json:"reference_id"`
}
type CreateRequestParams struct {
	CustomerId string `json:"customer_id"`
	Email      string `json:"email,omitempty"`
}
type ListRequestParams struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
