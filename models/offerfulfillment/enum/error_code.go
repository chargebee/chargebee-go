package enum

type ErrorCode string

const (
	ErrorCodeBillingUpdateFailed       ErrorCode = "billing_update_failed"
	ErrorCodeCheckoutAbandoned         ErrorCode = "checkout_abandoned"
	ErrorCodeExternalFulfillmentFailed ErrorCode = "external_fulfillment_failed"
	ErrorCodeInternalError             ErrorCode = "internal_error"
	ErrorCodeFulfillmentExpired        ErrorCode = "fulfillment_expired"
)
