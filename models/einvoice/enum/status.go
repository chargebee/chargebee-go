package enum

type Status string

const (
	StatusScheduled  Status = "scheduled"
	StatusSkipped    Status = "skipped"
	StatusInProgress Status = "in_progress"
	StatusSuccess    Status = "success"
	StatusFailed     Status = "failed"
	StatusRegistered Status = "registered"
)
