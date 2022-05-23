package enum

type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
	StatusDraft    Status = "draft"
)
