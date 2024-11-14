package omnichannelsubscription

import (
	omnichannelSubscriptionEnum "github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription/enum"
)

type OmnichannelSubscription struct {
	Id                           string                             `json:"id"`
	IdAtSource                   string                             `json:"id_at_source"`
	AppId                        string                             `json:"app_id"`
	Source                       omnichannelSubscriptionEnum.Source `json:"source"`
	CustomerId                   string                             `json:"customer_id"`
	CreatedAt                    int64                              `json:"created_at"`
	OmnichannelSubscriptionItems []*OmnichannelSubscriptionItem     `json:"omnichannel_subscription_items"`
	Object                       string                             `json:"object"`
}
type OmnichannelSubscriptionItem struct {
	Id                 string                                                                    `json:"id"`
	IdAtSource         string                                                                    `json:"id_at_source"`
	Status             omnichannelSubscriptionEnum.OmnichannelSubscriptionItemStatus             `json:"status"`
	CurrentTermStart   int64                                                                     `json:"current_term_start"`
	CurrentTermEnd     int64                                                                     `json:"current_term_end"`
	ExpiredAt          int64                                                                     `json:"expired_at"`
	ExpirationReason   omnichannelSubscriptionEnum.OmnichannelSubscriptionItemExpirationReason   `json:"expiration_reason"`
	CancelledAt        int64                                                                     `json:"cancelled_at"`
	CancellationReason omnichannelSubscriptionEnum.OmnichannelSubscriptionItemCancellationReason `json:"cancellation_reason"`
	Object             string                                                                    `json:"object"`
}
type ListRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type OmnichannelTransactionsForOmnichannelSubscriptionRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
