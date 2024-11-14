package enum

type Status string

const (
	StatusInProcess Status = "in_process"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)
