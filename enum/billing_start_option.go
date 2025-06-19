package enum

type BillingStartOption string

const (
	BillingStartOptionImmediately    BillingStartOption = "immediately"
	BillingStartOptionOnSpecificDate BillingStartOption = "on_specific_date"
)
