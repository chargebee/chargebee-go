package chargebee

import "net/http"

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

type UnknownResponse struct {
	apiResponse
}
