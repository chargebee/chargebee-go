package chargebee

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"strings"
)

func (request RequestObj) Request() (*Result, error) {
	return request.RequestWithContext(context.Background())
}
func (request RequestObj) RequestWithContext(ctx context.Context) (*Result, error) {
	return request.RequestWithContextWithEnv(ctx, DefaultConfig())
}
func (request RequestObj) RequestWithEnv(env Environment) (*Result, error) {
	return request.RequestWithContextWithEnv(context.Background(), env)
}
func (request RequestObj) RequestWithContextWithEnv(ctx context.Context, env Environment) (*Result, error) {
	path, body := getBody(request.Method, request.Path, request.Params)
	req, err := newRequest(ctx, env, request.Method, path, body, request.Header)
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
