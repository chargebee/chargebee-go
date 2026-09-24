package enum

type Status string

const (
	StatusScheduled     Status = "scheduled"
	StatusRescheduled   Status = "rescheduled"
	StatusSucceeded     Status = "succeeded"
	StatusFailed        Status = "failed"
	StatusDeferred      Status = "deferred"
	StatusDelivered     Status = "delivered"
	StatusOpened        Status = "opened"
	StatusBounced       Status = "bounced"
	StatusDropped       Status = "dropped"
	StatusActive        Status = "active"
	StatusArchived      Status = "archived"
	StatusDeleted       Status = "deleted"
	StatusAvailable     Status = "available"
	StatusExhausted     Status = "exhausted"
	StatusInGracePeriod Status = "in_grace_period"
)
