package enum

type LinkedPaymentStatus string

const (
	LinkedPaymentStatusInProgress     LinkedPaymentStatus = "in_progress"
	LinkedPaymentStatusSuccess        LinkedPaymentStatus = "success"
	LinkedPaymentStatusVoided         LinkedPaymentStatus = "voided"
	LinkedPaymentStatusFailure        LinkedPaymentStatus = "failure"
	LinkedPaymentStatusTimeout        LinkedPaymentStatus = "timeout"
	LinkedPaymentStatusNeedsAttention LinkedPaymentStatus = "needs_attention"
)
