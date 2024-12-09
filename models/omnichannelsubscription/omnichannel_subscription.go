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
	Object                       string                                                     `json:"object"`
}
type ListRequestParams struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
type OmnichannelTransactionsForOmnichannelSubscriptionRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
