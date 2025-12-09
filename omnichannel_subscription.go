package chargebee

type Source string

const (
	SourceAppleAppStore   Source = "apple_app_store"
	SourceGooglePlayStore Source = "google_play_store"
)

type OmnichannelSubscriptionItemStatus string

const (
	OmnichannelSubscriptionItemStatusActive        OmnichannelSubscriptionItemStatus = "active"
	OmnichannelSubscriptionItemStatusExpired       OmnichannelSubscriptionItemStatus = "expired"
	OmnichannelSubscriptionItemStatusCancelled     OmnichannelSubscriptionItemStatus = "cancelled"
	OmnichannelSubscriptionItemStatusInDunning     OmnichannelSubscriptionItemStatus = "in_dunning"
	OmnichannelSubscriptionItemStatusInGracePeriod OmnichannelSubscriptionItemStatus = "in_grace_period"
	OmnichannelSubscriptionItemStatusPaused        OmnichannelSubscriptionItemStatus = "paused"
)

type OmnichannelSubscriptionItemAutoRenewStatus string

const (
	OmnichannelSubscriptionItemAutoRenewStatusOff OmnichannelSubscriptionItemAutoRenewStatus = "off"
	OmnichannelSubscriptionItemAutoRenewStatusOn  OmnichannelSubscriptionItemAutoRenewStatus = "on"
)

type OmnichannelSubscriptionItemExpirationReason string

const (
	OmnichannelSubscriptionItemExpirationReasonBillingError        OmnichannelSubscriptionItemExpirationReason = "billing_error"
	OmnichannelSubscriptionItemExpirationReasonProductNotAvailable OmnichannelSubscriptionItemExpirationReason = "product_not_available"
	OmnichannelSubscriptionItemExpirationReasonOther               OmnichannelSubscriptionItemExpirationReason = "other"
)

type OmnichannelSubscriptionItemCancellationReason string

const (
	OmnichannelSubscriptionItemCancellationReasonCustomerCancelled                    OmnichannelSubscriptionItemCancellationReason = "customer_cancelled"
	OmnichannelSubscriptionItemCancellationReasonCustomerDidNotConsentToPriceIncrease OmnichannelSubscriptionItemCancellationReason = "customer_did_not_consent_to_price_increase"
	OmnichannelSubscriptionItemCancellationReasonRefundedDueToAppIssue                OmnichannelSubscriptionItemCancellationReason = "refunded_due_to_app_issue"
	OmnichannelSubscriptionItemCancellationReasonRefundedForOtherReason               OmnichannelSubscriptionItemCancellationReason = "refunded_for_other_reason"
	OmnichannelSubscriptionItemCancellationReasonMerchantRevoked                      OmnichannelSubscriptionItemCancellationReason = "merchant_revoked"
)

type OmnichannelTransactionType string

const (
	OmnichannelTransactionTypePurchase OmnichannelTransactionType = "purchase"
	OmnichannelTransactionTypeRenewal  OmnichannelTransactionType = "renewal"
)

type OmnichannelSubscription struct {
	Id                           string                                                     `json:"id"`
	IdAtSource                   string                                                     `json:"id_at_source"`
	AppId                        string                                                     `json:"app_id"`
	Source                       Source                                                     `json:"source"`
	CustomerId                   string                                                     `json:"customer_id"`
	CreatedAt                    int64                                                      `json:"created_at"`
	ResourceVersion              int64                                                      `json:"resource_version"`
	OmnichannelSubscriptionItems []*omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_items"`
	InitialPurchaseTransaction   *omnichanneltransaction.OmnichannelTransaction             `json:"initial_purchase_transaction"`
	Object                       string                                                     `json:"object"`
}
type ListRequest struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	Source     *filter.EnumFilter   `json:"source,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
type OmnichannelTransactionsForOmnichannelSubscriptionRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type MoveRequest struct {
	ToCustomerId string `json:"to_customer_id"`
}

type RetrieveResponse struct {
	OmnichannelSubscription *OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
}

type ListOmnichannelSubscriptionResponse struct {
	OmnichannelSubscription *OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
}

type ListResponse struct {
	List       []*ListOmnichannelSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                 `json:"next_offset,omitempty"`
}

type OmnichannelTransactionsForOmnichannelSubscriptionOmnichannelSubscriptionResponse struct {
	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`
}

type OmnichannelTransactionsForOmnichannelSubscriptionResponse struct {
	List       []*OmnichannelTransactionsForOmnichannelSubscriptionOmnichannelSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                                                              `json:"next_offset,omitempty"`
}

type MoveResponse struct {
	OmnichannelSubscription *OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
}
