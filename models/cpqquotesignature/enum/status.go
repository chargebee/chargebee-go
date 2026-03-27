package enum

type Status string

const (
	StatusDraft     Status = "draft"
	StatusActive    Status = "active"
	StatusSigned    Status = "signed"
	StatusExpired   Status = "expired"
	StatusCancelled Status = "cancelled"
	StatusDeclined  Status = "declined"
)
