package enum

type Status string

const (
	StatusActive    Status = "active"
	StatusFulfilled Status = "fulfilled"
	StatusBroken    Status = "broken"
	StatusCancelled Status = "cancelled"
)
