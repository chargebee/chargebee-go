package enum

type DiscountDurationType string

const (
	DiscountDurationTypeOneTime       DiscountDurationType = "one_time"
	DiscountDurationTypeForever       DiscountDurationType = "forever"
	DiscountDurationTypeLimitedPeriod DiscountDurationType = "limited_period"
)
