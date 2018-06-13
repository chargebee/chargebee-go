package enum

type Type string

const (
	TypeAuthorization   Type = "authorization"
	TypePayment         Type = "payment"
	TypeRefund          Type = "refund"
	TypePaymentReversal Type = "payment_reversal"
)
