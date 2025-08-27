package chargebee

func (request ListRequest) ListRequest() (*ResultList, error) {
	result, err := request.ListRequestWithEnv(DefaultConfig())
	return result, err
}
func (request ListRequest) ListRequestWithEnv(env Environment) (*ResultList, error) {
	path, body := getBody(request.Method, request.Path, request.Params)
	req, err := newRequest(env, request.Method, path, body, request.Header, request.subDomain, false)
	if err != nil {
		panic(err)
	}
	if request.Context != nil {
		req = req.WithContext(request.Context)
	}
	res, requestError := Do(req, request.isJsonRequest)
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
