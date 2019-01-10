package enum

type PaymentStatus string

const (
	PaymentStatusNotPaid PaymentStatus = "not_paid"
	PaymentStatusPaid    PaymentStatus = "paid"
)
