package enum

type Status string

const (
	StatusAvailable     Status = "available"
	StatusExhausted     Status = "exhausted"
	StatusScheduled     Status = "scheduled"
	StatusInGracePeriod Status = "in_grace_period"
)
