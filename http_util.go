package chargebee

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// RequestObj is the structure that contains the properties of regular request data.
type RequestObj struct {
	Params        *url.Values
	Method        string
	Path          string
	Header        map[string]string
	Context       context.Context `form:"-"`
	subDomain     string
	isJsonRequest bool
	JsonBody      string
	idempotent    bool
}

// ListRequestObj is the structure that contains the properties of list request data.
type ListRequestObj struct {
	Params        *url.Values
	Method        string
	Path          string
	Header        map[string]string
	Context       context.Context `form:"-"`
	subDomain     string
	isJsonRequest bool
	JsonBody      string
	idempotent    bool
}

type Request = RequestObj
type ListRequest = ListRequestObj

func basicAuth(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// Send prepares a RequestObj for Request operation.
func Send(method string, path string, params interface{}) (Request, error) {
	var form *url.Values

	if params != nil {
		var err error
		form, err = SerializeParams(params)
		if err != nil {
			return Request{}, err
		}
	}

	return Request{
		Params: form,
		Method: method,
		Path:   path,
	}, nil
}

func SendJsonRequest(method string, path string, params interface{}) (Request, error) {
	var body string

	if params != nil {
		if strings.ToUpper(method) == "POST" {
			jsonData, err := json.Marshal(params)
			if err != nil {
				return Request{}, err
			}
			body = string(jsonData)
		}

	}

	return Request{
		Method:        method,
		Path:          path,
		JsonBody:      body,
		isJsonRequest: true,
	}, nil
}

// SendList prepares a ListRequest for ListRequest operation.
func SendList(method string, path string, params interface{}) (ListRequest, error) {
	var form *url.Values
	if params != nil {
		var err error
		form, err = SerializeListParams(params)
		if err != nil {
			return ListRequest{}, err
		}
	}

	return ListRequest{
		Params: form,
		Method: method,
		Path:   path,
	}, nil
}

// AddParams add a new key-value pair to the RequestObj.Params.
// This is used to add extra/custom_field params  in the request data.
func (request Request) AddParams(key string, value interface{}) Request {
	request.Params.Set(key, fmt.Sprintf("%v", value))
	return request
}

// Headers add a new key-value pair to the RequestObj.Header .
// This is used to add custom headers .
func (request Request) Headers(key string, value string) Request {
	if request.Header == nil {
		request.Header = make(map[string]string)
	}
	request.Header[key] = value
	return request
}

// This is used to add idempotency key .
func (request Request) SetIdempotencyKey(idempotencyKey string) Request {
	if request.Header == nil {
		request.Header = make(map[string]string)
	}
	request.Header[IdempotencyHeader] = idempotencyKey
	return request
}

func (request Request) SetSubDomain(subDomain string) Request {
	request.subDomain = subDomain
	return request
}

func (request Request) SetIdempotency(idempotent bool) Request {
	request.idempotent = idempotent
	return request
}

// Context used for request. It may carry deadlines, cancelation signals,
// and other request-scoped values across API boundaries and between
// processes.
func (request Request) Contexts(ctx context.Context) Request {
	request.Context = ctx
	return request
}

// ListRequestObj methods

// AddParams add a new key-value pair to the ListRequest.Params.
// This is used to add extra/custom_field params in the request data.
func (request ListRequest) AddParams(key string, value interface{}) ListRequest {
	request.Params.Set(key, fmt.Sprintf("%v", value))
	return request
}

// Headers add a new key-value pair to the ListRequest.Header.
// This is used to add custom headers.
func (request ListRequest) Headers(key string, value string) ListRequest {
	if request.Header == nil {
		request.Header = make(map[string]string)
	}
	request.Header[key] = value
	return request
}

// SetIdempotencyKey is used to add idempotency key.
func (request ListRequest) SetIdempotencyKey(idempotencyKey string) ListRequest {
	if request.Header == nil {
		request.Header = make(map[string]string)
	}
	request.Header[IdempotencyHeader] = idempotencyKey
	return request
}

func (request ListRequest) SetSubDomain(subDomain string) ListRequest {
	request.subDomain = subDomain
	return request
}

func (request ListRequest) SetIdempotency(idempotent bool) ListRequest {
	request.idempotent = idempotent
	return request
}

// Context used for request. It may carry deadlines, cancelation signals,
// and other request-scoped values across API boundaries and between
// processes.
func (request ListRequest) Contexts(ctx context.Context) ListRequest {
	request.Context = ctx
	return request
}

// Common utility functions

func newRequest(env Environment, method string, path string, body io.Reader, headers map[string]string, subDomain string, isJsonRequest bool) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = env.apiBaseUrl(subDomain) + path
	httpReq, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	addHeaders(httpReq, env, isJsonRequest)
	addCustomHeaders(httpReq, headers)
	return httpReq, nil
}
func addHeaders(httpReq *http.Request, env Environment, isJsonRequest bool) {
	httpReq.Header.Add("Accept-Charset", Charset)
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("User-Agent", "ChargeBee-Go-Client v"+Version)
	httpReq.Header.Add("Authorization", "Basic "+basicAuth(env.Key))
	httpReq.Header.Add("Lang-Version", runtime.Version())
	httpReq.Header.Add("OS-Version", runtime.GOOS+" "+runtime.GOARCH)
	if isJsonRequest {
		httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	} else {
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
}
func addCustomHeaders(httpReq *http.Request, headers map[string]string) {
	for k, v := range headers {
		httpReq.Header.Add(k, v)
	}
}
func ensureIdempotencyKey(req *http.Request, isIdempotent bool) {
	if isIdempotent && req.Header.Get("X-CB-Idempotency-Key") == "" {
		req.Header.Set("X-CB-Idempotency-Key", uuid.NewString())
	}
}

func ensureRetryCountHeader(req *http.Request, attempt int) {
	req.Header.Set("X-CB-Retry-Attempt", strconv.Itoa(attempt))
}

// UnmarshalJSON is used to unmarshal the response to Result / ResultList struct.
func UnmarshalJSON(response []byte, result interface{}) error {
	err := json.Unmarshal(response, result)
	if err != nil {
		return err
	}
	if err := customFieldExtraction(result, response); err != nil {
		return err
	}
	return nil
}

// GetMap is used to unmarshal the json.RawMessage to map[string]interface{}.
func GetMap(rawMessage json.RawMessage) (map[string]interface{}, error) {
	data := json.NewDecoder(strings.NewReader(string(rawMessage)))
	data.UseNumber()
	var m map[string]interface{}
	if err := data.Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}
