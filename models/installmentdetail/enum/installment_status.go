package enum

type InstallmentStatus string

const (
	InstallmentStatusPosted     InstallmentStatus = "posted"
	InstallmentStatusPaymentDue InstallmentStatus = "payment_due"
	InstallmentStatusPaid       InstallmentStatus = "paid"
)
