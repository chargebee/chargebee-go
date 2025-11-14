package chargebee

import "context"

func (request ListRequest) ListRequest() (*ResultList, error) {
	env, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	result, err := request.ListRequestWithEnv(env)
	return result, err
}
func (request ListRequest) ListRequestWithEnv(env Environment) (*ResultList, error) {
	path, body := getBody(request.Method, request.Path, request.Params)
	req, err := newRequest(env, request.Method, path, body, request.Header, request.subDomain, false)
	if err != nil {
		return nil, err
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
