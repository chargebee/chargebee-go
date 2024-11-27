package enum

type CancellationReason string

const (
	CancellationReasonCustomerCancelled                    CancellationReason = "customer_cancelled"
	CancellationReasonCustomerDidNotConsentToPriceIncrease CancellationReason = "customer_did_not_consent_to_price_increase"
)
