package chargebee

import (
	"bytes"
	"io"
	"net/url"
	"strings"
)

func (request RequestObj) Request() (*Result, error) {
	result, err := request.RequestWithEnv(DefaultConfig())
	return result, err
}
func (request RequestObj) RequestWithEnv(env Environment) (*Result, error) {
	path, body := getBody(request.Method, request.Path, request.Params)
	req, err := newRequest(env, request.Method, path, body, request.Header)
	if err != nil {
		panic(err)
	}
	if request.Context != nil {
		req = req.WithContext(request.Context)
	}
	res, requestError := Do(req)
	result := &Result{}
	if requestError != nil {
		return result, requestError
	}

	if err := UnmarshalJSON(res.Body, result); err != nil {
		return result, nil
	}
	result.responseHeaders = res.Headers

	return result, requestError
}
func getBody(method string, path string, form *url.Values) (string, io.Reader) {
	var body io.Reader
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}
	return path, body
}
