package enum

type PeriodUnit string

const (
	PeriodUnitDay           PeriodUnit = "day"
	PeriodUnitWeek          PeriodUnit = "week"
	PeriodUnitMonth         PeriodUnit = "month"
	PeriodUnitYear          PeriodUnit = "year"
	PeriodUnitNotApplicable PeriodUnit = "not_applicable"
)
