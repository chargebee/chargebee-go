package chargebee

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (request Request) Request() (*Result, error) {
	result, err := request.RequestWithEnv(DefaultConfig())
	return result, err
}

func (request Request) RequestWithEnv(env Environment) (*Result, error) {
	var path string
	if request.isJsonRequest {
		path, _ = getJsonBody(request.Method, request.Path, request.JsonBody)
	} else {
		path, _ = getBody(request.Method, request.Path, request.Params)
	}
	httpURL := buildTelemetryHTTPURL(env, request.subDomain, path)

	result := &Result{}
	err := executeWithTelemetry(env, telemetryExecuteInput{
		resource:       request.telemetryResource,
		operation:      request.telemetryOperation,
		method:         request.Method,
		httpURL:        httpURL,
		requestHeaders: request.Header,
	}, func(headers map[string]string) (int, error) {
		req := request
		for key, value := range headers {
			req = req.Headers(key, value)
		}
		res, requestError := req.doRequestWithEnv(env)
		if requestError != nil {
			return 0, requestError
		}
		*result = *res
		return res.httpStatusCode, nil
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func (request Request) doRequestWithEnv(env Environment) (*Result, error) {
	var body io.Reader
	var path string
	if request.isJsonRequest {
		path, body = getJsonBody(request.Method, request.Path, request.JsonBody)
	} else {
		path, body = getBody(request.Method, request.Path, request.Params)
	}
	req, err := newRequest(env, request.Method, path, body, request.Header, request.subDomain, request.isJsonRequest)
	if err != nil {
		panic(err)
	}
	if request.Context != nil {
		req = req.WithContext(context.WithValue(request.Context, cbEnvKey, env))
	} else {
		req = req.WithContext(context.WithValue(req.Context(), cbEnvKey, env))
	}
	res, requestError := Do(req, request.idempotent)
	result := &Result{}
	if requestError != nil {
		return result, requestError
	}

	if err := UnmarshalJSON(res.Body, result); err != nil && res.StatusCode != http.StatusNoContent {
		return result, err
	}
	result.responseHeaders = res.Headers
	result.httpStatusCode = res.StatusCode
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

func getJsonBody(method string, path string, jsonBody string) (string, io.Reader) {
	var body io.Reader
	if jsonBody != "" {
		body = bytes.NewBufferString(jsonBody)
	}
	return path, body
}
