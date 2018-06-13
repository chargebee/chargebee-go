package enum

type DurationType string

const (
	DurationTypeOneTime       DurationType = "one_time"
	DurationTypeForever       DurationType = "forever"
	DurationTypeLimitedPeriod DurationType = "limited_period"
)
