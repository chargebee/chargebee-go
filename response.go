package chargebee

import (
	"net/http"
	"strconv"
)

type responseWrapper interface {
	setMeta(meta *apiResponse)
}

type apiResponse struct {
	Headers    http.Header
	StatusText string
	StatusCode int
	Body       []byte `json:"-"`
}

func (r *apiResponse) setMeta(meta *apiResponse) {
	r.Headers = meta.Headers
	r.StatusText = meta.StatusText
	r.StatusCode = meta.StatusCode
	r.Body = meta.Body
}

func (r *apiResponse) IsIdempotencyReplayed() bool {
	value := r.Headers.Get(IdempotencyReplayHeader)
	if replayed, err := strconv.ParseBool(value); err != nil {
		return false
	} else {
		return replayed
	}
}

type UnknownResponse struct {
	apiResponse
}
