package recordedpurchase

import (
	recordedPurchaseEnum "github.com/chargebee/chargebee-go/v3/models/recordedpurchase/enum"
)

type RecordedPurchase struct {
	Id                             string                           `json:"id"`
	CustomerId                     string                           `json:"customer_id"`
	AppId                          string                           `json:"app_id"`
	Source                         recordedPurchaseEnum.Source      `json:"source"`
	Status                         recordedPurchaseEnum.Status      `json:"status"`
	OmnichannelTransactionId       string                           `json:"omnichannel_transaction_id"`
	CreatedAt                      int64                            `json:"created_at"`
	ResourceVersion                int64                            `json:"resource_version"`
	LinkedOmnichannelSubscriptions []*LinkedOmnichannelSubscription `json:"linked_omnichannel_subscriptions"`
	LinkedOmnichannelOneTimeOrders []*LinkedOmnichannelOneTimeOrder `json:"linked_omnichannel_one_time_orders"`
	ErrorDetail                    *ErrorDetail                     `json:"error_detail"`
	Object                         string                           `json:"object"`
}
type LinkedOmnichannelSubscription struct {
	OmnichannelSubscriptionId string `json:"omnichannel_subscription_id"`
	Object                    string `json:"object"`
}
type LinkedOmnichannelOneTimeOrder struct {
	OmnichannelOneTimeOrderId string `json:"omnichannel_one_time_order_id"`
	Object                    string `json:"object"`
}
type ErrorDetail struct {
	ErrorMessage string `json:"error_message"`
	Object       string `json:"object"`
}
type CreateRequestParams struct {
	AppId                   string                               `json:"app_id"`
	Customer                *CreateCustomerParams                `json:"customer,omitempty"`
	AppleAppStore           *CreateAppleAppStoreParams           `json:"apple_app_store,omitempty"`
	GooglePlayStore         *CreateGooglePlayStoreParams         `json:"google_play_store,omitempty"`
	OmnichannelSubscription *CreateOmnichannelSubscriptionParams `json:"omnichannel_subscription,omitempty"`
}
type CreateCustomerParams struct {
	Id string `json:"id"`
}
type CreateAppleAppStoreParams struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Receipt       string `json:"receipt,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
}
type CreateGooglePlayStoreParams struct {
	PurchaseToken string `json:"purchase_token,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
}
type CreateOmnichannelSubscriptionParams struct {
	Id string `json:"id,omitempty"`
}
