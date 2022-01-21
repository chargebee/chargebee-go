package enum

type Duration string

const (
	DurationAllBillingCycles       Duration = "all_billing_cycles"
	DurationSpecificNumberOfCycles Duration = "specific_number_of_cycles"
	DurationTillSpecificDate       Duration = "till_specific_date"
)
