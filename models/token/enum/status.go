package enum

type Status string

const (
	StatusNew      Status = "new"
	StatusExpired  Status = "expired"
	StatusConsumed Status = "consumed"
)
