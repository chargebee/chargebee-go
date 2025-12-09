package chargebee

type Status string

const (
	StatusActive        Status = "active"
	StatusExpired       Status = "expired"
	StatusCancelled     Status = "cancelled"
	StatusInDunning     Status = "in_dunning"
	StatusInGracePeriod Status = "in_grace_period"
	StatusPaused        Status = "paused"
)

type AutoRenewStatus string

const (
	AutoRenewStatusOff AutoRenewStatus = "off"
	AutoRenewStatusOn  AutoRenewStatus = "on"
)

type ExpirationReason string

const (
	ExpirationReasonBillingError        ExpirationReason = "billing_error"
	ExpirationReasonProductNotAvailable ExpirationReason = "product_not_available"
	ExpirationReasonOther               ExpirationReason = "other"
)

type CancellationReason string

const (
	CancellationReasonCustomerCancelled                    CancellationReason = "customer_cancelled"
	CancellationReasonCustomerDidNotConsentToPriceIncrease CancellationReason = "customer_did_not_consent_to_price_increase"
	CancellationReasonRefundedDueToAppIssue                CancellationReason = "refunded_due_to_app_issue"
	CancellationReasonRefundedForOtherReason               CancellationReason = "refunded_for_other_reason"
	CancellationReasonMerchantRevoked                      CancellationReason = "merchant_revoked"
)

type OmnichannelSubscriptionItem struct {
	Id                                string                                                               `json:"id"`
	ItemIdAtSource                    string                                                               `json:"item_id_at_source"`
	ItemParentIdAtSource              string                                                               `json:"item_parent_id_at_source"`
	Status                            Status                                                               `json:"status"`
	AutoRenewStatus                   AutoRenewStatus                                                      `json:"auto_renew_status"`
	CurrentTermStart                  int64                                                                `json:"current_term_start"`
	CurrentTermEnd                    int64                                                                `json:"current_term_end"`
	ExpiredAt                         int64                                                                `json:"expired_at"`
	ExpirationReason                  ExpirationReason                                                     `json:"expiration_reason"`
	CancelledAt                       int64                                                                `json:"cancelled_at"`
	CancellationReason                CancellationReason                                                   `json:"cancellation_reason"`
	GracePeriodExpiresAt              int64                                                                `json:"grace_period_expires_at"`
	ResumesAt                         int64                                                                `json:"resumes_at"`
	HasScheduledChanges               bool                                                                 `json:"has_scheduled_changes"`
	ResourceVersion                   int64                                                                `json:"resource_version"`
	OmnichannelSubscriptionItemOffers []*omnichannelsubscriptionitemoffer.OmnichannelSubscriptionItemOffer `json:"omnichannel_subscription_item_offers"`
	UpcomingRenewal                   UpcomingRenewal                                                      `json:"upcoming_renewal"`
	LinkedItem                        LinkedItem                                                           `json:"linked_item"`
	Object                            string                                                               `json:"object"`
}
type UpcomingRenewal struct {
	PriceCurrency string `json:"price_currency"`
	PriceUnits    int64  `json:"price_units"`
	PriceNanos    int64  `json:"price_nanos"`
	Object        string `json:"object"`
}
type LinkedItem struct {
	Id       string `json:"id"`
	LinkedAt int64  `json:"linked_at"`
	Object   string `json:"object"`
}
type ListOmniSubItemScheduleChangesRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type ListOmniSubItemScheduleChangesOmnichannelSubscriptionItemResponse struct {
	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`
}

type ListOmniSubItemScheduleChangesResponse struct {
	List       []*ListOmniSubItemScheduleChangesOmnichannelSubscriptionItemResponse `json:"list,omitempty"`
	NextOffset string                                                               `json:"next_offset,omitempty"`
}
