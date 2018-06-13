package enum

type CancelReason string

const (
	CancelReasonNotPaid                         CancelReason = "not_paid"
	CancelReasonNoCard                          CancelReason = "no_card"
	CancelReasonFraudReviewFailed               CancelReason = "fraud_review_failed"
	CancelReasonNonCompliantEuCustomer          CancelReason = "non_compliant_eu_customer"
	CancelReasonTaxCalculationFailed            CancelReason = "tax_calculation_failed"
	CancelReasonCurrencyIncompatibleWithGateway CancelReason = "currency_incompatible_with_gateway"
	CancelReasonNonCompliantCustomer            CancelReason = "non_compliant_customer"
)
