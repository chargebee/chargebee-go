package enum

type BillingDayOfWeekMode string

const (
	BillingDayOfWeekModeUsingDefaults BillingDayOfWeekMode = "using_defaults"
	BillingDayOfWeekModeManuallySet   BillingDayOfWeekMode = "manually_set"
)
