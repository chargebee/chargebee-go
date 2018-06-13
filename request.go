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
	res, err1 := Do(req)
	result := &Result{}
	if err1 != nil {
		return result, err1
	}
	err1 = UnmarshalJSON([]byte(res), result)
	return result, err1
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
