package enum

type ActivePaymentAttemptStatus string

const (
	ActivePaymentAttemptStatusInited                 ActivePaymentAttemptStatus = "inited"
	ActivePaymentAttemptStatusRequiresIdentification ActivePaymentAttemptStatus = "requires_identification"
	ActivePaymentAttemptStatusRequiresChallenge      ActivePaymentAttemptStatus = "requires_challenge"
	ActivePaymentAttemptStatusRequiresRedirection    ActivePaymentAttemptStatus = "requires_redirection"
	ActivePaymentAttemptStatusAuthorized             ActivePaymentAttemptStatus = "authorized"
	ActivePaymentAttemptStatusRefused                ActivePaymentAttemptStatus = "refused"
	ActivePaymentAttemptStatusPendingAuthorization   ActivePaymentAttemptStatus = "pending_authorization"
)
