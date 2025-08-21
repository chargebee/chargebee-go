package enum

type OmnichannelOneTimeOrderItemCancellationReason string

const (
	OmnichannelOneTimeOrderItemCancellationReasonCustomerCancelled                    OmnichannelOneTimeOrderItemCancellationReason = "customer_cancelled"
	OmnichannelOneTimeOrderItemCancellationReasonCustomerDidNotConsentToPriceIncrease OmnichannelOneTimeOrderItemCancellationReason = "customer_did_not_consent_to_price_increase"
	OmnichannelOneTimeOrderItemCancellationReasonRefundedDueToAppIssue                OmnichannelOneTimeOrderItemCancellationReason = "refunded_due_to_app_issue"
	OmnichannelOneTimeOrderItemCancellationReasonRefundedForOtherReason               OmnichannelOneTimeOrderItemCancellationReason = "refunded_for_other_reason"
	OmnichannelOneTimeOrderItemCancellationReasonMerchantRevoked                      OmnichannelOneTimeOrderItemCancellationReason = "merchant_revoked"
)
