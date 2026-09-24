package enum

type Status string

const (
	StatusInitiated      Status = "initiated"
	StatusFundsWithdrawn Status = "funds_withdrawn"
	StatusInReview       Status = "in_review"
	StatusCancelled      Status = "cancelled"
	StatusLost           Status = "lost"
	StatusWon            Status = "won"
)
