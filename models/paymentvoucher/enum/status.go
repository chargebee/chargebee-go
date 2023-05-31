package enum

type Status string

const (
	StatusActive   Status = "active"
	StatusConsumed Status = "consumed"
	StatusExpired  Status = "expired"
	StatusFailure  Status = "failure"
)
