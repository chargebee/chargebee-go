package chargebee

type Source string

const (
	SourceAppleAppStore   Source = "apple_app_store"
	SourceGooglePlayStore Source = "google_play_store"
)

type Status string

const (
	StatusInProcess Status = "in_process"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
	StatusIgnored   Status = "ignored"
)

type RecordedPurchase struct {
	Id                             string                           `json:"id"`
	CustomerId                     string                           `json:"customer_id"`
	AppId                          string                           `json:"app_id"`
	Source                         Source                           `json:"source"`
	Status                         Status                           `json:"status"`
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
type CreateRequest struct {
	AppId                   string                         `json:"app_id"`
	Customer                *CreateCustomer                `json:"customer,omitempty"`
	AppleAppStore           *CreateAppleAppStore           `json:"apple_app_store,omitempty"`
	GooglePlayStore         *CreateGooglePlayStore         `json:"google_play_store,omitempty"`
	OmnichannelSubscription *CreateOmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
}
type CreateCustomer struct {
	Id string `json:"id"`
}
type CreateAppleAppStore struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Receipt       string `json:"receipt,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
}
type CreateGooglePlayStore struct {
	PurchaseToken string `json:"purchase_token,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
}
type CreateOmnichannelSubscription struct {
	Id string `json:"id,omitempty"`
}

type CreateResponse struct {
	RecordedPurchase *RecordedPurchase  `json:"recorded_purchase,omitempty"`
	Customer         *customer.Customer `json:"customer,omitempty"`
}

type RetrieveResponse struct {
	RecordedPurchase *RecordedPurchase `json:"recorded_purchase,omitempty"`
}
