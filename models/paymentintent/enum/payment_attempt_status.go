package enum

type PaymentAttemptStatus string

const (
	PaymentAttemptStatusInited                 PaymentAttemptStatus = "inited"
	PaymentAttemptStatusRequiresIdentification PaymentAttemptStatus = "requires_identification"
	PaymentAttemptStatusRequiresChallenge      PaymentAttemptStatus = "requires_challenge"
	PaymentAttemptStatusRequiresRedirection    PaymentAttemptStatus = "requires_redirection"
	PaymentAttemptStatusAuthorized             PaymentAttemptStatus = "authorized"
	PaymentAttemptStatusRefused                PaymentAttemptStatus = "refused"
)
