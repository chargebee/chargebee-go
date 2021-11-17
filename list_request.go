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
	opts := []RequestOption{}
	if env.HttpClient != nil {
		opts = append(opts, SetClient(env.HttpClient))
	}
	res, err1 := Do(req, opts...)
	result := &ResultList{}
	if err1 != nil {
		return result, err1
	}
	err1 = UnmarshalJSON([]byte(res), result)
	return result, nil
}
