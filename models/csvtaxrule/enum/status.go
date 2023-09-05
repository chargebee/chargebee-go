package enum

type Status string

const (
	StatusActive    Status = "active"
	StatusExpired   Status = "expired"
	StatusScheduled Status = "scheduled"
)
