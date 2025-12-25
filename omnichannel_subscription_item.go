package chargebee

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

type OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferCategory string

const (
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferCategoryIntroductory        OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferCategory = "introductory"
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferCategoryPromotional         OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferCategory = "promotional"
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferCategoryDeveloperDetermined OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferCategory = "developer_determined"
)

type OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferType string

const (
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferTypeFreeTrial  OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferType = "free_trial"
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferTypePayUpFront OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferType = "pay_up_front"
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferTypePayAsYouGo OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferType = "pay_as_you_go"
)

type OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferDiscountType string

const (
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferDiscountTypeFixedAmount OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferDiscountType = "fixed_amount"
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferDiscountTypePercentage  OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferDiscountType = "percentage"
	OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferDiscountTypePrice       OmnichannelSubscriptionItemOmnichannelSubscriptionItemOfferDiscountType = "price"
)

// just struct
type OmnichannelSubscriptionItem struct {
	Id                                string                                        `json:"id"`
	ItemIdAtSource                    string                                        `json:"item_id_at_source"`
	ItemParentIdAtSource              string                                        `json:"item_parent_id_at_source"`
	Status                            OmnichannelSubscriptionItemStatus             `json:"status"`
	AutoRenewStatus                   OmnichannelSubscriptionItemAutoRenewStatus    `json:"auto_renew_status"`
	CurrentTermStart                  int64                                         `json:"current_term_start"`
	CurrentTermEnd                    int64                                         `json:"current_term_end"`
	ExpiredAt                         int64                                         `json:"expired_at"`
	ExpirationReason                  OmnichannelSubscriptionItemExpirationReason   `json:"expiration_reason"`
	CancelledAt                       int64                                         `json:"cancelled_at"`
	CancellationReason                OmnichannelSubscriptionItemCancellationReason `json:"cancellation_reason"`
	GracePeriodExpiresAt              int64                                         `json:"grace_period_expires_at"`
	ResumesAt                         int64                                         `json:"resumes_at"`
	HasScheduledChanges               bool                                          `json:"has_scheduled_changes"`
	ResourceVersion                   int64                                         `json:"resource_version"`
	OmnichannelSubscriptionItemOffers []*OmnichannelSubscriptionItemOffer           `json:"omnichannel_subscription_item_offers"`
	UpcomingRenewal                   UpcomingRenewal                               `json:"upcoming_renewal"`
	LinkedItem                        LinkedItem                                    `json:"linked_item"`
	Object                            string                                        `json:"object"`
}

// sub resources
type OmnichannelSubscriptionItemUpcomingRenewal struct {
	PriceCurrency string `json:"price_currency"`
	PriceUnits    int64  `json:"price_units"`
	PriceNanos    int64  `json:"price_nanos"`
	Object        string `json:"object"`
}

type OmnichannelSubscriptionItemLinkedItem struct {
	Id       string `json:"id"`
	LinkedAt int64  `json:"linked_at"`
	Object   string `json:"object"`
}

// operations
// input params
type OmnichannelSubscriptionItemListOmniSubItemScheduleChangesRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *OmnichannelSubscriptionItemListOmniSubItemScheduleChangesRequest) payload() any { return r }

// operation sub response
type OmnichannelSubscriptionItemListOmniSubItemScheduleChangesOmnichannelSubscriptionItemResponse struct {
	OmnichannelSubscriptionItemScheduledChange OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
}

// operation response
type OmnichannelSubscriptionItemListOmniSubItemScheduleChangesResponse struct {
	List       []*OmnichannelSubscriptionItemListOmniSubItemScheduleChangesOmnichannelSubscriptionItemResponse `json:"list,omitempty"`
	NextOffset string                                                                                          `json:"next_offset,omitempty"`
	apiResponse
}
