package enum

type Status string

const (
	StatusScheduled Status = "scheduled"
	StatusFailed    Status = "failed"
	StatusSucceeded Status = "succeeded"
)
