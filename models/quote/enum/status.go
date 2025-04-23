package enum

type Status string

const (
	StatusOpen             Status = "open"
	StatusAccepted         Status = "accepted"
	StatusDeclined         Status = "declined"
	StatusInvoiced         Status = "invoiced"
	StatusClosed           Status = "closed"
	StatusPendingApproval  Status = "pending_approval"
	StatusApprovalRejected Status = "approval_rejected"
	StatusProposed         Status = "proposed"
	StatusVoided           Status = "voided"
	StatusExpired          Status = "expired"
)
