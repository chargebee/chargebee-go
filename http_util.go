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
	"strings"
)

// RequestObj is the structure that contains the properties of any request data.
type RequestObj struct {
	Params        *url.Values
	Method        string
	Path          string
	Header        map[string]string
	Context       context.Context `form:"-"`
	subDomain     string
	isJsonRequest bool
	JsonBody      string
}

func basicAuth(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// Send prepares a RequestObj for Request operation.
func Send(method string, path string, params interface{}) RequestObj {
	var form *url.Values

	if params != nil {
		form = SerializeParams(params)
	}

	return RequestObj{
		Params: form,
		Method: method,
		Path:   path,
	}
}

func SendJsonRequest(method string, path string, params interface{}) RequestObj {
	var body string

	if params != nil {
		if strings.ToUpper(method) == "POST" {
			jsonData, err := json.Marshal(params)
			if err != nil {
				panic(err)
			}
			body = string(jsonData)
		}

	}

	return RequestObj{
		Method:        method,
		Path:          path,
		JsonBody:      body,
		isJsonRequest: true,
	}
}

// SendList prepares a RequestObj for ListRequest operation.
func SendList(method string, path string, params interface{}) RequestObj {
	var form *url.Values
	if params != nil {
		form = SerializeListParams(params)
	}

	return RequestObj{
		Params: form,
		Method: method,
		Path:   path,
	}
}

// AddParams add a new key-value pair to the RequestObj.Params.
// This is used to add extra/custom_field params  in the request data.
func (request RequestObj) AddParams(key string, value interface{}) RequestObj {
	request.Params.Set(key, fmt.Sprintf("%v", value))
	return request
}

// Headers add a new key-value pair to the RequestObj.Header .
// This is used to add custom headers .
func (request RequestObj) Headers(key string, value string) RequestObj {
	if request.Header == nil {
		request.Header = make(map[string]string)
	}
	request.Header[key] = value
	return request
}

// This is used to add idempotency key .
func (request RequestObj) SetIdempotencyKey(idempotencyKey string) RequestObj {
	if request.Header == nil {
		request.Header = make(map[string]string)
	}
	request.Header[IdempotencyHeader] = idempotencyKey
	return request
}

func (request RequestObj) SetSubDomain(subDomain string) RequestObj {
	request.subDomain = subDomain
	return request
}

// Context used for request. It may carry deadlines, cancelation signals,
// and other request-scoped values across API boundaries and between
// processes.
func (request RequestObj) Contexts(ctx context.Context) RequestObj {
	request.Context = ctx
	return request
}

func newRequest(env Environment, method string, path string, body io.Reader, headers map[string]string, subDomain string, isJsonRequest bool) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = env.apiBaseUrl(subDomain) + path
	httpReq, err := http.NewRequest(method, path, body)
	if err != nil {
		panic(err)
	}
	addHeaders(httpReq, env, isJsonRequest)
	addCustomHeaders(httpReq, headers)
	return httpReq, err
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

// UnmarshalJSON is used to unmarshal the response to Result / ResultList struct.
func UnmarshalJSON(response []byte, result interface{}) error {
	err := json.Unmarshal(response, result)
	if err != nil {
		return err
	}
	customFieldExtraction(result, response)
	return nil
}

// GetMap is used to unmarshal the json.RawMessage to map[string]interface{}.
func GetMap(rawMessage json.RawMessage) map[string]interface{} {
	data := json.NewDecoder(strings.NewReader(string(rawMessage)))
	data.UseNumber()
	var m map[string]interface{}
	if err := data.Decode(&m); err != nil {
		panic(err)
	}
	return m
}
