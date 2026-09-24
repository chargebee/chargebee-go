package enum

type OmnichannelSubscriptionItemExpirationReason string

const (
	OmnichannelSubscriptionItemExpirationReasonBillingError                 OmnichannelSubscriptionItemExpirationReason = "billing_error"
	OmnichannelSubscriptionItemExpirationReasonProductNotAvailable          OmnichannelSubscriptionItemExpirationReason = "product_not_available"
	OmnichannelSubscriptionItemExpirationReasonOther                        OmnichannelSubscriptionItemExpirationReason = "other"
	OmnichannelSubscriptionItemExpirationReasonSubscriptionNotFoundInSource OmnichannelSubscriptionItemExpirationReason = "subscription_not_found_in_source"
)
