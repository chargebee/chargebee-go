package enum

type Status string

const (
	StatusActive     Status = "active"
	StatusCompleted  Status = "completed"
	StatusCancelled  Status = "cancelled"
	StatusTerminated Status = "terminated"
)
