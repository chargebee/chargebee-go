package enum

type Status string

const (
	StatusPaid       Status = "paid"
	StatusPosted     Status = "posted"
	StatusPaymentDue Status = "payment_due"
	StatusNotPaid    Status = "not_paid"
	StatusVoided     Status = "voided"
	StatusPending    Status = "pending"
)
