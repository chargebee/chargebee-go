package enum

type Status string

const (
	StatusAdjusted  Status = "adjusted"
	StatusRefunded  Status = "refunded"
	StatusRefundDue Status = "refund_due"
	StatusVoided    Status = "voided"
)
