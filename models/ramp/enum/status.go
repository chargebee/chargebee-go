package enum

type Status string

const (
	StatusScheduled Status = "scheduled"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
)
