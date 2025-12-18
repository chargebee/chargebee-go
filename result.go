package chargebee

import (
	"net/http"
	"strconv"
)

type ResultList struct {
	List            []*Result `json:"list"`
	NextOffset      string    `json:"next_offset"`
	responseHeaders http.Header
	httpStatusCode  int
}

type Result struct {
	FailedEvents    interface{} `json:"failed_events,omitempty"`
	ExpiresAt       interface{} `json:"expires_at,omitempty"`
	BatchId         interface{} `json:"batch_id,omitempty"`
	Success         interface{} `json:"success,omitempty"`
	ScheduledAt     interface{} `json:"scheduled_at,omitempty"`
	List            interface{} `json:"list,omitempty"`
	responseHeaders http.Header
	httpStatusCode  int
}

func (rl *ResultList) GetResponseHeaders() http.Header {
	return rl.responseHeaders
}

func (rl *ResultList) GetHttpStatusCode() int {
	return rl.httpStatusCode
}

func (r *Result) GetResponseHeaders() http.Header {
	return r.responseHeaders
}

func (r *Result) GetHttpStatusCode() int {
	return r.httpStatusCode
}

func (r *Result) IsIdempotencyReplayed() bool {
	value := r.responseHeaders.Get(IdempotencyReplayHeader)
	replayed, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return replayed
}
