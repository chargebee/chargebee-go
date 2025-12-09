package chargebee

type Source string

const (
	SourceAppleAppStore   Source = "apple_app_store"
	SourceGooglePlayStore Source = "google_play_store"
)

type OmnichannelOneTimeOrderItemCancellationReason string

const (
	OmnichannelOneTimeOrderItemCancellationReasonCustomerCancelled                    OmnichannelOneTimeOrderItemCancellationReason = "customer_cancelled"
	OmnichannelOneTimeOrderItemCancellationReasonCustomerDidNotConsentToPriceIncrease OmnichannelOneTimeOrderItemCancellationReason = "customer_did_not_consent_to_price_increase"
	OmnichannelOneTimeOrderItemCancellationReasonRefundedDueToAppIssue                OmnichannelOneTimeOrderItemCancellationReason = "refunded_due_to_app_issue"
	OmnichannelOneTimeOrderItemCancellationReasonRefundedForOtherReason               OmnichannelOneTimeOrderItemCancellationReason = "refunded_for_other_reason"
	OmnichannelOneTimeOrderItemCancellationReasonMerchantRevoked                      OmnichannelOneTimeOrderItemCancellationReason = "merchant_revoked"
)

type OmnichannelTransactionType string

const (
	OmnichannelTransactionTypePurchase OmnichannelTransactionType = "purchase"
	OmnichannelTransactionTypeRenewal  OmnichannelTransactionType = "renewal"
)

type OmnichannelOneTimeOrder struct {
	Id                           string                                                     `json:"id"`
	AppId                        string                                                     `json:"app_id"`
	CustomerId                   string                                                     `json:"customer_id"`
	IdAtSource                   string                                                     `json:"id_at_source"`
	Origin                       string                                                     `json:"origin"`
	Source                       Source                                                     `json:"source"`
	CreatedAt                    int64                                                      `json:"created_at"`
	ResourceVersion              int64                                                      `json:"resource_version"`
	OmnichannelOneTimeOrderItems []*omnichannelonetimeorderitem.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_items"`
	PurchaseTransaction          *omnichanneltransaction.OmnichannelTransaction             `json:"purchase_transaction"`
	Object                       string                                                     `json:"object"`
}
type ListRequest struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	Source     *filter.EnumFilter   `json:"source,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}

type RetrieveResponse struct {
	OmnichannelOneTimeOrder *OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`
}

type ListOmnichannelOneTimeOrderResponse struct {
	OmnichannelOneTimeOrder *OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`
}

type ListResponse struct {
	List       []*ListOmnichannelOneTimeOrderResponse `json:"list,omitempty"`
	NextOffset string                                 `json:"next_offset,omitempty"`
}
