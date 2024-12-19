package enum

type Status string

const (
	StatusActive        Status = "active"
	StatusExpired       Status = "expired"
	StatusCancelled     Status = "cancelled"
	StatusInDunning     Status = "in_dunning"
	StatusInGracePeriod Status = "in_grace_period"
)
