package enum

type Status string

const (
	StatusPosted     Status = "posted"
	StatusPaymentDue Status = "payment_due"
	StatusPaid       Status = "paid"
)
