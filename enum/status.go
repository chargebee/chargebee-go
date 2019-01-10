package enum

type Status string

const (
	StatusScheduled Status = "scheduled"
	StatusUnclaimed Status = "unclaimed"
	StatusClaimed   Status = "claimed"
	StatusCancelled Status = "cancelled"
	StatusExpired   Status = "expired"
)
