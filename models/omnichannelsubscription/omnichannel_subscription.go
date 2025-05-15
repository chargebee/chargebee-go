package omnichannelsubscription

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	omnichannelSubscriptionEnum "github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription/enum"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"
)

type OmnichannelSubscription struct {
	Id                           string                                                     `json:"id"`
	IdAtSource                   string                                                     `json:"id_at_source"`
	AppId                        string                                                     `json:"app_id"`
	Source                       omnichannelSubscriptionEnum.Source                         `json:"source"`
	CustomerId                   string                                                     `json:"customer_id"`
	CreatedAt                    int64                                                      `json:"created_at"`
	ResourceVersion              int64                                                      `json:"resource_version"`
	OmnichannelSubscriptionItems []*omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_items"`
	InitialPurchaseTransaction   OmnichannelTransaction                                     `json:"initial_purchase_transaction"`
	Object                       string                                                     `json:"object"`
}
type OmnichannelTransaction struct {
	Id              string                                                     `json:"id"`
	IdAtSource      string                                                     `json:"id_at_source"`
	AppId           string                                                     `json:"app_id"`
	PriceCurrency   string                                                     `json:"price_currency"`
	PriceUnits      int64                                                      `json:"price_units"`
	PriceNanos      int64                                                      `json:"price_nanos"`
	Type            omnichannelSubscriptionEnum.InitialPurchaseTransactionType `json:"type"`
	TransactedAt    int64                                                      `json:"transacted_at"`
	CreatedAt       int64                                                      `json:"created_at"`
	ResourceVersion int64                                                      `json:"resource_version"`
	Object          string                                                     `json:"object"`
}
type ListRequestParams struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	Source     *filter.EnumFilter   `json:"source,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
type OmnichannelTransactionsForOmnichannelSubscriptionRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
