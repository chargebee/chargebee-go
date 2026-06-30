package chargebee

import (
	"context"
)

func (request ListRequest) ListRequest() (*ResultList, error) {
	result, err := request.ListRequestWithEnv(DefaultConfig())
	return result, err
}

func (request ListRequest) ListRequestWithEnv(env Environment) (*ResultList, error) {
	path, _ := getBody(request.Method, request.Path, request.Params)
	httpURL := buildTelemetryHTTPURL(env, request.subDomain, path)

	result := &ResultList{}
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
		res, requestError := req.doListRequestWithEnv(env)
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

func (request ListRequest) doListRequestWithEnv(env Environment) (*ResultList, error) {
	path, body := getBody(request.Method, request.Path, request.Params)
	req, err := newRequest(env, request.Method, path, body, request.Header, request.subDomain, false)
	if err != nil {
		panic(err)
	}
	if request.Context != nil {
		req = req.WithContext(context.WithValue(request.Context, cbEnvKey, env))
	} else {
		req = req.WithContext(context.WithValue(req.Context(), cbEnvKey, env))
	}
	res, requestError := Do(req, request.idempotent)
	result := &ResultList{}

	if requestError != nil {
		return result, requestError
	}

	if err := UnmarshalJSON(res.Body, result); err != nil {
		return result, err
	}

	result.responseHeaders = res.Headers
	result.httpStatusCode = res.StatusCode
	return result, nil
}
