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
	LinkedOmnichannelSubscriptions []*LinkedOmnichannelSubscription `json:"linked_omnichannel_subscriptions"`
	ErrorDetail                    *ErrorDetail                     `json:"error_detail"`
	Object                         string                           `json:"object"`
}
type LinkedOmnichannelSubscription struct {
	OmnichannelSubscriptionId string `json:"omnichannel_subscription_id"`
	Object                    string `json:"object"`
}
type ErrorDetail struct {
	ErrorMessage string `json:"error_message"`
	Object       string `json:"object"`
}
type CreateRequestParams struct {
	AppId         string                     `json:"app_id"`
	Customer      *CreateCustomerParams      `json:"customer,omitempty"`
	AppleAppStore *CreateAppleAppStoreParams `json:"apple_app_store,omitempty"`
}
type CreateCustomerParams struct {
	Id string `json:"id"`
}
type CreateAppleAppStoreParams struct {
	TransactionId string `json:"transaction_id,omitempty"`
}
