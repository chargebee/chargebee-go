package enum

type PaymentMethodStatus string

const (
	PaymentMethodStatusValid               PaymentMethodStatus = "valid"
	PaymentMethodStatusExpiring            PaymentMethodStatus = "expiring"
	PaymentMethodStatusExpired             PaymentMethodStatus = "expired"
	PaymentMethodStatusInvalid             PaymentMethodStatus = "invalid"
	PaymentMethodStatusPendingVerification PaymentMethodStatus = "pending_verification"
)
