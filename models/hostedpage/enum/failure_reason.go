package enum

type FailureReason string

const (
	FailureReasonCardError   FailureReason = "card_error"
	FailureReasonServerError FailureReason = "server_error"
)
