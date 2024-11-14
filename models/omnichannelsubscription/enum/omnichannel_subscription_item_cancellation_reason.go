package enum

type OmnichannelSubscriptionItemCancellationReason string

const (
	OmnichannelSubscriptionItemCancellationReasonCustomerCancelled                    OmnichannelSubscriptionItemCancellationReason = "customer_cancelled"
	OmnichannelSubscriptionItemCancellationReasonCustomerDidNotConsentToPriceIncrease OmnichannelSubscriptionItemCancellationReason = "customer_did_not_consent_to_price_increase"
)
