package chargebee

type OmnichannelOneTimeOrderSource string

const (
	OmnichannelOneTimeOrderSourceAppleAppStore   OmnichannelOneTimeOrderSource = "apple_app_store"
	OmnichannelOneTimeOrderSourceGooglePlayStore OmnichannelOneTimeOrderSource = "google_play_store"
)

type OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReason string

const (
	OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReasonCustomerCancelled                    OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReason = "customer_cancelled"
	OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReasonCustomerDidNotConsentToPriceIncrease OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReason = "customer_did_not_consent_to_price_increase"
	OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReasonRefundedDueToAppIssue                OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReason = "refunded_due_to_app_issue"
	OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReasonRefundedForOtherReason               OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReason = "refunded_for_other_reason"
	OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReasonMerchantRevoked                      OmnichannelOneTimeOrderOmnichannelOneTimeOrderItemCancellationReason = "merchant_revoked"
)

type OmnichannelOneTimeOrderOmnichannelTransactionType string

const (
	OmnichannelOneTimeOrderOmnichannelTransactionTypePurchase OmnichannelOneTimeOrderOmnichannelTransactionType = "purchase"
	OmnichannelOneTimeOrderOmnichannelTransactionTypeRenewal  OmnichannelOneTimeOrderOmnichannelTransactionType = "renewal"
)

// just struct
type OmnichannelOneTimeOrder struct {
	Id                           string                         `json:"id"`
	AppId                        string                         `json:"app_id"`
	CustomerId                   string                         `json:"customer_id"`
	IdAtSource                   string                         `json:"id_at_source"`
	Origin                       string                         `json:"origin"`
	Source                       OmnichannelOneTimeOrderSource  `json:"source"`
	CreatedAt                    int64                          `json:"created_at"`
	ResourceVersion              int64                          `json:"resource_version"`
	OmnichannelOneTimeOrderItems []*OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_items"`
	PurchaseTransaction          *OmnichannelTransaction        `json:"purchase_transaction"`
	Object                       string                         `json:"object"`
}

// sub resources
// operations
// input params
type OmnichannelOneTimeOrderListRequest struct {
	Limit      *int32        `json:"limit,omitempty"`
	Offset     string        `json:"offset,omitempty"`
	Source     *EnumFilter   `json:"source,omitempty"`
	CustomerId *StringFilter `json:"customer_id,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *OmnichannelOneTimeOrderListRequest) payload() any { return r }

// operation response
type OmnichannelOneTimeOrderRetrieveResponse struct {
	OmnichannelOneTimeOrder *OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`
	apiResponse
}

// operation sub response
type OmnichannelOneTimeOrderListOmnichannelOneTimeOrderResponse struct {
	OmnichannelOneTimeOrder *OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`
}

// operation response
type OmnichannelOneTimeOrderListResponse struct {
	List       []*OmnichannelOneTimeOrderListOmnichannelOneTimeOrderResponse `json:"list,omitempty"`
	NextOffset string                                                        `json:"next_offset,omitempty"`
	apiResponse
}
