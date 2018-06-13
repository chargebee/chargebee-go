package enum

type BillingDateMode string

const (
	BillingDateModeUsingDefaults BillingDateMode = "using_defaults"
	BillingDateModeManuallySet   BillingDateMode = "manually_set"
)
