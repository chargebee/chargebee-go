package chargebee

type RecordedPurchaseSource string

const (
	RecordedPurchaseSourceAppleAppStore   RecordedPurchaseSource = "apple_app_store"
	RecordedPurchaseSourceGooglePlayStore RecordedPurchaseSource = "google_play_store"
)

type RecordedPurchaseStatus string

const (
	RecordedPurchaseStatusInProcess RecordedPurchaseStatus = "in_process"
	RecordedPurchaseStatusCompleted RecordedPurchaseStatus = "completed"
	RecordedPurchaseStatusFailed    RecordedPurchaseStatus = "failed"
	RecordedPurchaseStatusIgnored   RecordedPurchaseStatus = "ignored"
)

// just struct
type RecordedPurchase struct {
	Id                             string                                           `json:"id"`
	CustomerId                     string                                           `json:"customer_id"`
	AppId                          string                                           `json:"app_id"`
	Source                         RecordedPurchaseSource                           `json:"source"`
	Status                         RecordedPurchaseStatus                           `json:"status"`
	OmnichannelTransactionId       string                                           `json:"omnichannel_transaction_id"`
	CreatedAt                      int64                                            `json:"created_at"`
	ResourceVersion                int64                                            `json:"resource_version"`
	LinkedOmnichannelSubscriptions []*RecordedPurchaseLinkedOmnichannelSubscription `json:"linked_omnichannel_subscriptions"`
	LinkedOmnichannelOneTimeOrders []*RecordedPurchaseLinkedOmnichannelOneTimeOrder `json:"linked_omnichannel_one_time_orders"`
	ErrorDetail                    *RecordedPurchaseErrorDetail                     `json:"error_detail"`
	Object                         string                                           `json:"object"`
}

// sub resources
type RecordedPurchaseLinkedOmnichannelSubscription struct {
	OmnichannelSubscriptionId string `json:"omnichannel_subscription_id"`
	Object                    string `json:"object"`
}
type RecordedPurchaseLinkedOmnichannelOneTimeOrder struct {
	OmnichannelOneTimeOrderId string `json:"omnichannel_one_time_order_id"`
	Object                    string `json:"object"`
}
type RecordedPurchaseErrorDetail struct {
	ErrorMessage string `json:"error_message"`
	Object       string `json:"object"`
}

// operations
// input params
type RecordedPurchaseCreateRequest struct {
	AppId                   string                                         `json:"app_id"`
	Customer                *RecordedPurchaseCreateCustomer                `json:"customer,omitempty"`
	AppleAppStore           *RecordedPurchaseCreateAppleAppStore           `json:"apple_app_store,omitempty"`
	GooglePlayStore         *RecordedPurchaseCreateGooglePlayStore         `json:"google_play_store,omitempty"`
	OmnichannelSubscription *RecordedPurchaseCreateOmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *RecordedPurchaseCreateRequest) payload() any { return r }

// input sub resource params single
type RecordedPurchaseCreateCustomer struct {
	Id string `json:"id"`
}

// input sub resource params single
type RecordedPurchaseCreateAppleAppStore struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Receipt       string `json:"receipt,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
}

// input sub resource params single
type RecordedPurchaseCreateGooglePlayStore struct {
	PurchaseToken string `json:"purchase_token,omitempty"`
	ProductId     string `json:"product_id,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
}

// input sub resource params single
type RecordedPurchaseCreateOmnichannelSubscription struct {
	Id string `json:"id,omitempty"`
}

// operation response
type RecordedPurchaseCreateResponse struct {
	RecordedPurchase *RecordedPurchase `json:"recorded_purchase,omitempty"`
	Customer         Customer          `json:"customer,omitempty"`
	apiResponse
}

// operation response
type RecordedPurchaseRetrieveResponse struct {
	RecordedPurchase *RecordedPurchase `json:"recorded_purchase,omitempty"`
	apiResponse
}
