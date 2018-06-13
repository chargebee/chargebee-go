package enum

type Status string

const (
	StatusValid               Status = "valid"
	StatusExpiring            Status = "expiring"
	StatusExpired             Status = "expired"
	StatusInvalid             Status = "invalid"
	StatusPendingVerification Status = "pending_verification"
)
