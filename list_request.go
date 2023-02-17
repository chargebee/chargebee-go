package chargebee

import "context"

func (request RequestObj) ListRequest() (*ResultList, error) {
	return request.ListRequestWithContext(context.Background())
}
func (request RequestObj) ListRequestWithContext(ctx context.Context) (*ResultList, error) {
	return request.ListRequestWithContextWithEnv(ctx, DefaultConfig())
}
func (request RequestObj) ListRequestWithEnv(env Environment) (*ResultList, error) {
	return request.ListRequestWithContextWithEnv(context.Background(), env)
}
func (request RequestObj) ListRequestWithContextWithEnv(ctx context.Context, env Environment) (*ResultList, error) {
	path, body := getBody(request.Method, request.Path, request.Params)
	req, err := newRequest(ctx, env, request.Method, path, body, request.Header)
	if err != nil {
		panic(err)
	}
	res, err1 := Do(req)
	result := &ResultList{}
	if err1 != nil {
		return result, err1
	}
	err1 = UnmarshalJSON([]byte(res), result)
	return result, nil
}
