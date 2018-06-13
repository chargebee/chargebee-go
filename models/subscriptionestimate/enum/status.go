package enum

type Status string

const (
	StatusFuture      Status = "future"
	StatusInTrial     Status = "in_trial"
	StatusActive      Status = "active"
	StatusNonRenewing Status = "non_renewing"
	StatusPaused      Status = "paused"
	StatusCancelled   Status = "cancelled"
)
