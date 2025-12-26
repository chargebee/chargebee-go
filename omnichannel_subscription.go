package chargebee

type OmnichannelSubscriptionSource string

const (
	OmnichannelSubscriptionSourceAppleAppStore   OmnichannelSubscriptionSource = "apple_app_store"
	OmnichannelSubscriptionSourceGooglePlayStore OmnichannelSubscriptionSource = "google_play_store"
)

type OmnichannelSubscriptionOmnichannelSubscriptionItemStatus string

const (
	OmnichannelSubscriptionOmnichannelSubscriptionItemStatusActive        OmnichannelSubscriptionOmnichannelSubscriptionItemStatus = "active"
	OmnichannelSubscriptionOmnichannelSubscriptionItemStatusExpired       OmnichannelSubscriptionOmnichannelSubscriptionItemStatus = "expired"
	OmnichannelSubscriptionOmnichannelSubscriptionItemStatusCancelled     OmnichannelSubscriptionOmnichannelSubscriptionItemStatus = "cancelled"
	OmnichannelSubscriptionOmnichannelSubscriptionItemStatusInDunning     OmnichannelSubscriptionOmnichannelSubscriptionItemStatus = "in_dunning"
	OmnichannelSubscriptionOmnichannelSubscriptionItemStatusInGracePeriod OmnichannelSubscriptionOmnichannelSubscriptionItemStatus = "in_grace_period"
	OmnichannelSubscriptionOmnichannelSubscriptionItemStatusPaused        OmnichannelSubscriptionOmnichannelSubscriptionItemStatus = "paused"
)

type OmnichannelSubscriptionOmnichannelSubscriptionItemAutoRenewStatus string

const (
	OmnichannelSubscriptionOmnichannelSubscriptionItemAutoRenewStatusOff OmnichannelSubscriptionOmnichannelSubscriptionItemAutoRenewStatus = "off"
	OmnichannelSubscriptionOmnichannelSubscriptionItemAutoRenewStatusOn  OmnichannelSubscriptionOmnichannelSubscriptionItemAutoRenewStatus = "on"
)

type OmnichannelSubscriptionOmnichannelSubscriptionItemExpirationReason string

const (
	OmnichannelSubscriptionOmnichannelSubscriptionItemExpirationReasonBillingError        OmnichannelSubscriptionOmnichannelSubscriptionItemExpirationReason = "billing_error"
	OmnichannelSubscriptionOmnichannelSubscriptionItemExpirationReasonProductNotAvailable OmnichannelSubscriptionOmnichannelSubscriptionItemExpirationReason = "product_not_available"
	OmnichannelSubscriptionOmnichannelSubscriptionItemExpirationReasonOther               OmnichannelSubscriptionOmnichannelSubscriptionItemExpirationReason = "other"
)

type OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReason string

const (
	OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReasonCustomerCancelled                    OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReason = "customer_cancelled"
	OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReasonCustomerDidNotConsentToPriceIncrease OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReason = "customer_did_not_consent_to_price_increase"
	OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReasonRefundedDueToAppIssue                OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReason = "refunded_due_to_app_issue"
	OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReasonRefundedForOtherReason               OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReason = "refunded_for_other_reason"
	OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReasonMerchantRevoked                      OmnichannelSubscriptionOmnichannelSubscriptionItemCancellationReason = "merchant_revoked"
)

type OmnichannelSubscriptionInitialPurchaseTransactionType string

const (
	OmnichannelSubscriptionInitialPurchaseTransactionTypePurchase OmnichannelSubscriptionInitialPurchaseTransactionType = "purchase"
	OmnichannelSubscriptionInitialPurchaseTransactionTypeRenewal  OmnichannelSubscriptionInitialPurchaseTransactionType = "renewal"
)

// just struct
type OmnichannelSubscription struct {
	Id                           string                         `json:"id"`
	IdAtSource                   string                         `json:"id_at_source"`
	AppId                        string                         `json:"app_id"`
	Source                       OmnichannelSubscriptionSource  `json:"source"`
	CustomerId                   string                         `json:"customer_id"`
	CreatedAt                    int64                          `json:"created_at"`
	ResourceVersion              int64                          `json:"resource_version"`
	OmnichannelSubscriptionItems []*OmnichannelSubscriptionItem `json:"omnichannel_subscription_items"`
	InitialPurchaseTransaction   *OmnichannelTransaction        `json:"initial_purchase_transaction"`
	Object                       string                         `json:"object"`
}

// sub resources
// operations
// input params
type OmnichannelSubscriptionListRequest struct {
	Limit      *int32        `json:"limit,omitempty"`
	Offset     string        `json:"offset,omitempty"`
	Source     *EnumFilter   `json:"source,omitempty"`
	CustomerId *StringFilter `json:"customer_id,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *OmnichannelSubscriptionListRequest) payload() any { return r }

type OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionRequest) payload() any {
	return r
}

type OmnichannelSubscriptionMoveRequest struct {
	ToCustomerId string `json:"to_customer_id"`
	apiRequest   `json:"-" form:"-"`
}

func (r *OmnichannelSubscriptionMoveRequest) payload() any { return r }

// operation response
type OmnichannelSubscriptionRetrieveResponse struct {
	OmnichannelSubscription *OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
	apiResponse
}

// operation sub response
type OmnichannelSubscriptionListOmnichannelSubscriptionResponse struct {
	OmnichannelSubscription *OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
}

// operation response
type OmnichannelSubscriptionListResponse struct {
	List       []*OmnichannelSubscriptionListOmnichannelSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                                        `json:"next_offset,omitempty"`
	apiResponse
}

// operation sub response
type OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionOmnichannelSubscriptionResponse struct {
	OmnichannelTransaction OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`
}

// operation response
type OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionResponse struct {
	List       []*OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionOmnichannelSubscriptionResponse `json:"list,omitempty"`
	NextOffset string                                                                                                     `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type OmnichannelSubscriptionMoveResponse struct {
	OmnichannelSubscription *OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`
	apiResponse
}
