package enum

type BillingAlignmentMode string

const (
	BillingAlignmentModeImmediate BillingAlignmentMode = "immediate"
	BillingAlignmentModeDelayed   BillingAlignmentMode = "delayed"
)
