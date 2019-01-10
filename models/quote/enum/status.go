package enum

type Status string

const (
	StatusOpen     Status = "open"
	StatusAccepted Status = "accepted"
	StatusDeclined Status = "declined"
	StatusInvoiced Status = "invoiced"
	StatusClosed   Status = "closed"
)
