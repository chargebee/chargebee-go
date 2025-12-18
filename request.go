package chargebee

import (
	"bytes"
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

// Request represents a ChargebeeRequest to the Chargebee API.
// It is embedded by other concrete ChargebeeRequest types like
// AddonCreateRequest, CustomerUpdateRequest, etc
type ChargebeeRequest[ReqType any, ResType responseSetter] struct {
	requestParams *ReqType
	urlParams     *url.Values
	method        string
	path          string
	headers       *http.Header
	context       context.Context `form:"-" json:"-"`
	subDomain     string
	jsonBody      string
	isJsonRequest bool
	isIdempotent  bool
	isListRequest bool
}

type BlankRequest struct{}

func NewRequest[ReqType any, ResType responseSetter](method, path string, req *ReqType) *ChargebeeRequest[ReqType, ResType] {
	return &ChargebeeRequest[ReqType, ResType]{
		requestParams: req,
		method:        method,
		path:          path,
	}
}

func (r *ChargebeeRequest[ReqType, ResType]) prepare() error {
	if r.isJsonRequest {
		if r.requestParams != nil && strings.ToUpper(r.method) == "POST" {
			if jsonData, err := json.Marshal(r.requestParams); err != nil {
				return fmt.Errorf("failed to marshal params: %w", err)
			} else {
				r.jsonBody = string(jsonData)
			}
		}
	} else if r.isListRequest {
		if r.requestParams != nil {
			form := SerializeListParams(r.requestParams)
			r.urlParams = form
		}
	} else {
		if r.requestParams != nil {
			form := SerializeParams(r.requestParams)
			r.urlParams = form
		}
	}

	return nil
}

// AddParams add a new key-value pair to the RequestObj.Params.
// This is used to add extra/custom_field params  in the request data.
func (r *ChargebeeRequest[ReqType, ResType]) AddParam(key string, value string) {
	if r.urlParams == nil {
		r.urlParams = &url.Values{}
	}
	r.urlParams.Set(key, value)
}

func (r *ChargebeeRequest[ReqType, ResType]) AddHeader(key string, value string) {
	if r.headers == nil {
		r.headers = &http.Header{}
	}
	r.headers.Add(key, value)
}

func (r *ChargebeeRequest[ReqType, ResType]) SetIdempotencyKey(idempotencyKey string) {
	r.AddHeader(IdempotencyHeader, idempotencyKey)
}

func (r *ChargebeeRequest[ReqType, ResType]) SetSubDomain(subDomain string) {
	r.subDomain = subDomain
}

func (r *ChargebeeRequest[ReqType, ResType]) SetIdempotency(idempotent bool) {
	r.isIdempotent = idempotent
}

// Context used for request. It may carry deadlines, cancelation signals,
// and other request-scoped values across API boundaries and between
// processes.
func (r *ChargebeeRequest[ReqType, ResType]) WithContext(ctx context.Context) {
	r.context = ctx
}

func (r *ChargebeeRequest[ReqType, ResType]) Send(transport *Transport) (ResType, error) {
	result, err := r.SendWithEnv(transport, DefaultConfig())
	return *result, err
}

func (r *ChargebeeRequest[ReqType, ResType]) SendWithEnv(transport *Transport, env Environment) (*ResType, error) {
	err := r.prepare()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %w", err)
	}
	var body io.Reader
	var path string

	// Now use 'request' instead of 'req' for accessing fields
	if r.isJsonRequest {
		path, body = getJsonBody(r.method, r.path, r.jsonBody)
	} else {
		path, body = getBody(r.method, r.path, r.urlParams)
	}

	reqObj, err := newRequest(env, r.method, path, body, r.headers, r.subDomain, r.isJsonRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create new chargebee request: %w", err)
	}

	if r.context != nil {
		reqObj = reqObj.WithContext(context.WithValue(r.context, cbEnvKey, env))
	} else {
		reqObj = reqObj.WithContext(context.WithValue(reqObj.Context(), cbEnvKey, env))
	}

	res, requestError := Do(reqObj, r.isIdempotent)
	result := new(ResType)
	if requestError != nil {
		return result, requestError
	}

	if err := UnmarshalJSON(res.Body, result); err != nil && res.StatusCode != http.StatusNoContent {
		return result, err
	}

	// if resultWithResponseMeta, ok := result.(responseSetter); ok {
	// 	resultWithResponseMeta.SetResponseMeta(&ResponseMeta{
	// 		Headers:    res.Headers,
	// 		Status:     res.Status,
	// 		StatusCode: res.StatusCode,
	// 	})
	// }

	(*result).SetResponseMeta(&ResponseMeta{
		Headers:    res.Headers,
		Status:     res.Status,
		StatusCode: res.StatusCode,
	})

	return result, requestError
}

func basicAuth(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func newRequest(env Environment, method string, path string, body io.Reader, headers *http.Header, subDomain string, isJsonRequest bool) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = env.apiBaseUrl(subDomain) + path
	httpReq, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
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

func addCustomHeaders(httpReq *http.Request, headers *http.Header) {
	for k, v := range *headers {
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
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
