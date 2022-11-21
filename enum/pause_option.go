package enum

type PauseOption string

const (
	PauseOptionImmediately   PauseOption = "immediately"
	PauseOptionEndOfTerm     PauseOption = "end_of_term"
	PauseOptionSpecificDate  PauseOption = "specific_date"
	PauseOptionBillingCycles PauseOption = "billing_cycles"
)
