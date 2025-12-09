package chargebee

type BillingConfiguration struct {
	IsCalendarBillingEnabled bool           `json:"is_calendar_billing_enabled"`
	BillingDates             []*BillingDate `json:"billing_dates"`
	Object                   string         `json:"object"`
}
type BillingDate struct {
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
	Object    string `json:"object"`
}
