package virtualbankaccount

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	virtualBankAccountEnum "github.com/chargebee/chargebee-go/models/virtualbankaccount/enum"
)

type VirtualBankAccount struct {
	Id               string                        `json:"id"`
	CustomerId       string                        `json:"customer_id"`
	Email            string                        `json:"email"`
	Scheme           virtualBankAccountEnum.Scheme `json:"scheme"`
	BankName         string                        `json:"bank_name"`
	AccountNumber    string                        `json:"account_number"`
	RoutingNumber    string                        `json:"routing_number"`
	SwiftCode        string                        `json:"swift_code"`
	Gateway          enum.Gateway                  `json:"gateway"`
	GatewayAccountId string                        `json:"gateway_account_id"`
	ResourceVersion  int64                         `json:"resource_version"`
	UpdatedAt        int64                         `json:"updated_at"`
	CreatedAt        int64                         `json:"created_at"`
	ReferenceId      string                        `json:"reference_id"`
	Deleted          bool                          `json:"deleted"`
	Object           string                        `json:"object"`
}
type CreateUsingPermanentTokenRequestParams struct {
	CustomerId  string                        `json:"customer_id"`
	ReferenceId string                        `json:"reference_id"`
	Scheme      virtualBankAccountEnum.Scheme `json:"scheme,omitempty"`
}
type CreateRequestParams struct {
	CustomerId string                        `json:"customer_id"`
	Email      string                        `json:"email,omitempty"`
	Scheme     virtualBankAccountEnum.Scheme `json:"scheme,omitempty"`
}
type ListRequestParams struct {
	Limit      *int32                  `json:"limit,omitempty"`
	Offset     string                  `json:"offset,omitempty"`
	CustomerId *filter.StringFilter    `json:"customer_id,omitempty"`
	UpdatedAt  *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt  *filter.TimestampFilter `json:"created_at,omitempty"`
}
