package chargebee

// just struct
type BillingConfiguration struct {
	IsCalendarBillingEnabled bool                               `json:"is_calendar_billing_enabled"`
	BillingDates             []*BillingConfigurationBillingDate `json:"billing_dates"`
	Object                   string                             `json:"object"`
}

// sub resources
type BillingConfigurationBillingDate struct {
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
	Object    string `json:"object"`
}

// operations
// input params
