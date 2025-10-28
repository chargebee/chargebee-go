package enum

type AuthorizationReason string

const (
	AuthorizationReasonBlockingFunds    AuthorizationReason = "blocking_funds"
	AuthorizationReasonVerification     AuthorizationReason = "verification"
	AuthorizationReasonScheduledCapture AuthorizationReason = "scheduled_capture"
)
