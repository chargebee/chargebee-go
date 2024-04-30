package enum

type DiscountsToAddDurationType string

const (
	DiscountsToAddDurationTypeOneTime       DiscountsToAddDurationType = "one_time"
	DiscountsToAddDurationTypeForever       DiscountsToAddDurationType = "forever"
	DiscountsToAddDurationTypeLimitedPeriod DiscountsToAddDurationType = "limited_period"
)
