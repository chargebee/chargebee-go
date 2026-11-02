package enum

type Status string

const (
	StatusActive    Status = "active"
	StatusCompleted Status = "completed"
	StatusPaused    Status = "paused"
	StatusFailed    Status = "failed"
)
