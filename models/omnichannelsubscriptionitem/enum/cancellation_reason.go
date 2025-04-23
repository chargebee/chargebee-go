package enum

type CancellationReason string

const (
	CancellationReasonCustomerCancelled                    CancellationReason = "customer_cancelled"
	CancellationReasonCustomerDidNotConsentToPriceIncrease CancellationReason = "customer_did_not_consent_to_price_increase"
	CancellationReasonRefundedDueToAppIssue                CancellationReason = "refunded_due_to_app_issue"
	CancellationReasonRefundedForOtherReason               CancellationReason = "refunded_for_other_reason"
)
