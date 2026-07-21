package enum

type AsyncResponseStatus string

const (
	AsyncResponseStatusSuccess AsyncResponseStatus = "success"
	AsyncResponseStatusFailed  AsyncResponseStatus = "failed"
)
