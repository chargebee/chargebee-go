package enum

type OmnichannelSubscriptionItemCancellationReason string

const (
	OmnichannelSubscriptionItemCancellationReasonCustomerCancelled                    OmnichannelSubscriptionItemCancellationReason = "customer_cancelled"
	OmnichannelSubscriptionItemCancellationReasonCustomerDidNotConsentToPriceIncrease OmnichannelSubscriptionItemCancellationReason = "customer_did_not_consent_to_price_increase"
	OmnichannelSubscriptionItemCancellationReasonRefundedDueToAppIssue                OmnichannelSubscriptionItemCancellationReason = "refunded_due_to_app_issue"
	OmnichannelSubscriptionItemCancellationReasonRefundedForOtherReason               OmnichannelSubscriptionItemCancellationReason = "refunded_for_other_reason"
	OmnichannelSubscriptionItemCancellationReasonMerchantRevoked                      OmnichannelSubscriptionItemCancellationReason = "merchant_revoked"
)
