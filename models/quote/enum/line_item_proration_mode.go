package enum

type LineItemProrationMode string

const (
	LineItemProrationModeReset                 LineItemProrationMode = "reset"
	LineItemProrationModeDelta                 LineItemProrationMode = "delta"
	LineItemProrationModeServicePeriodRevision LineItemProrationMode = "service_period_revision"
	LineItemProrationModeAdjustedTerm          LineItemProrationMode = "adjusted_term"
)
