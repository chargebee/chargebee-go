package enum

type ExpirationReason string

const (
	ExpirationReasonBillingError                 ExpirationReason = "billing_error"
	ExpirationReasonProductNotAvailable          ExpirationReason = "product_not_available"
	ExpirationReasonOther                        ExpirationReason = "other"
	ExpirationReasonSubscriptionNotFoundInSource ExpirationReason = "subscription_not_found_in_source"
)
