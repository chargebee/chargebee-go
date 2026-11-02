package enum

type Status string

const (
	StatusValid               Status = "valid"
	StatusInvalid             Status = "invalid"
	StatusPendingVerification Status = "pending_verification"
)
