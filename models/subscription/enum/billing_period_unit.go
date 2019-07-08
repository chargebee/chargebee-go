package enum

type BillingPeriodUnit string

const (
	BillingPeriodUnitDay   BillingPeriodUnit = "day"
	BillingPeriodUnitWeek  BillingPeriodUnit = "week"
	BillingPeriodUnitMonth BillingPeriodUnit = "month"
	BillingPeriodUnitYear  BillingPeriodUnit = "year"
)
