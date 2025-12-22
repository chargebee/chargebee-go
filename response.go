package chargebee

import "net/http"

type responseWrapper interface {
	SetMetadata(meta *apiResponse)
	Meta() *apiResponse
}

type apiResponse struct {
	headers    http.Header
	statusText string
	statusCode int
	body       []byte
}

func (r *apiResponse) Meta() *apiResponse {

}

func (r *apiResponse) SetMetadata(rm *apiResponse) {
	r.headers = rm.headers
	r.statusText = rm.statusText
	r.statusCode = rm.statusCode
	r.body = rm.body
}

type UnknownResponse struct {
	apiResponse
}
