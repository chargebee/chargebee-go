package enum

type Status string

const (
	StatusActive   Status = "active"
	StatusExpired  Status = "expired"
	StatusArchived Status = "archived"
	StatusDeleted  Status = "deleted"
)
