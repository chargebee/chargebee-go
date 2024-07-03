package enum

type CancelOption string

const (
	CancelOptionImmediately      CancelOption = "immediately"
	CancelOptionEndOfTerm        CancelOption = "end_of_term"
	CancelOptionSpecificDate     CancelOption = "specific_date"
	CancelOptionEndOfBillingTerm CancelOption = "end_of_billing_term"
)
