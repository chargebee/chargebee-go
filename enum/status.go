package enum

type Status string

const (
	StatusActive        Status = "active"
	StatusArchived      Status = "archived"
	StatusDeleted       Status = "deleted"
	StatusAvailable     Status = "available"
	StatusExhausted     Status = "exhausted"
	StatusScheduled     Status = "scheduled"
	StatusInGracePeriod Status = "in_grace_period"
)
