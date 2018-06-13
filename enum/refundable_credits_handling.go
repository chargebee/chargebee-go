package enum

type RefundableCreditsHandling string

const (
	RefundableCreditsHandlingNoAction       RefundableCreditsHandling = "no_action"
	RefundableCreditsHandlingScheduleRefund RefundableCreditsHandling = "schedule_refund"
)
