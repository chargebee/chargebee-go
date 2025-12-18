package chargebee

import "net/http"

type responseSetter interface {
	SetResponseMeta(responseMeta *ResponseMeta)
}

type ResponseMeta struct {
	Headers    http.Header
	Status     string
	StatusCode int
}

func (r *ResponseMeta) SetResponseMeta(responseMeta *ResponseMeta) {
	r.Headers = responseMeta.Headers
	r.Status = responseMeta.Status
	r.StatusCode = responseMeta.StatusCode
}

type CBResponse struct {
	ResponseMeta
	Body []byte
}
