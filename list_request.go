package chargebee

func (request RequestObj) ListRequest() (*ResultList, error) {
	result, err := request.ListRequestWithEnv(DefaultConfig())
	return result, err
}
func (request RequestObj) ListRequestWithEnv(env Environment) (*ResultList, error) {
	path, body := getBody(request.Method, request.Path, request.Params)
	req, err := newRequest(env, request.Method, path, body, request.Header)
	if err != nil {
		panic(err)
	}
	if request.Context != nil {
		req = req.WithContext(request.Context)
	}
	res, requestError := Do(req)
	result := &ResultList{}

	if requestError != nil {
		return result, requestError
	}

	if err := UnmarshalJSON(res.Body, result); err != nil {
		return result, err
	}

	result.responseHeaders = res.Headers
	return result, nil
}
