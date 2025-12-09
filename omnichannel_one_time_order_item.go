package chargebee

type CancellationReason string

const (
	CancellationReasonCustomerCancelled                    CancellationReason = "customer_cancelled"
	CancellationReasonCustomerDidNotConsentToPriceIncrease CancellationReason = "customer_did_not_consent_to_price_increase"
	CancellationReasonRefundedDueToAppIssue                CancellationReason = "refunded_due_to_app_issue"
	CancellationReasonRefundedForOtherReason               CancellationReason = "refunded_for_other_reason"
	CancellationReasonMerchantRevoked                      CancellationReason = "merchant_revoked"
)

type OmnichannelOneTimeOrderItem struct {
	Id                 string             `json:"id"`
	ItemIdAtSource     string             `json:"item_id_at_source"`
	ItemTypeAtSource   string             `json:"item_type_at_source"`
	Quantity           int32              `json:"quantity"`
	CancelledAt        int64              `json:"cancelled_at"`
	CancellationReason CancellationReason `json:"cancellation_reason"`
	CreatedAt          int64              `json:"created_at"`
	ResourceVersion    int64              `json:"resource_version"`
	Object             string             `json:"object"`
}
