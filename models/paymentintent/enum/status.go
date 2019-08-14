package enum

type Status string

const (
	StatusInited     Status = "inited"
	StatusInProgress Status = "in_progress"
	StatusAuthorized Status = "authorized"
	StatusConsumed   Status = "consumed"
	StatusExpired    Status = "expired"
)
