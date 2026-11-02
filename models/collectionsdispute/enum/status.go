package enum

type Status string

const (
	StatusOpen      Status = "open"
	StatusResolved  Status = "resolved"
	StatusWithdrawn Status = "withdrawn"
)
