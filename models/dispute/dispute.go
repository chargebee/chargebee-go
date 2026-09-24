package dispute

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	disputeEnum "github.com/chargebee/chargebee-go/v3/models/dispute/enum"
)

type Dispute struct {
	Id               string             `json:"id"`
	CustomerId       string             `json:"customer_id"`
	TransactionId    string             `json:"transaction_id"`
	GatewayAccountId string             `json:"gateway_account_id"`
	IdAtGateway      string             `json:"id_at_gateway"`
	CurrencyCode     string             `json:"currency_code"`
	Amount           int64              `json:"amount"`
	Reason           string             `json:"reason"`
	Status           disputeEnum.Status `json:"status"`
	Type             disputeEnum.Type   `json:"type"`
	IsPartialDispute bool               `json:"is_partial_dispute"`
	CreatedAt        int64              `json:"created_at"`
	ResourceVersion  int64              `json:"resource_version"`
	UpdatedAt        int64              `json:"updated_at"`
	Object           string             `json:"object"`
}
type ListRequestParams struct {
	Limit         *int32                  `json:"limit,omitempty"`
	Offset        string                  `json:"offset,omitempty"`
	Id            *filter.StringFilter    `json:"id,omitempty"`
	Status        *filter.EnumFilter      `json:"status,omitempty"`
	Type          *filter.EnumFilter      `json:"type,omitempty"`
	CustomerId    *filter.StringFilter    `json:"customer_id,omitempty"`
	TransactionId *filter.StringFilter    `json:"transaction_id,omitempty"`
	Amount        *filter.NumberFilter    `json:"amount,omitempty"`
	CreatedAt     *filter.TimestampFilter `json:"created_at,omitempty"`
}
