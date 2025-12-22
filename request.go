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

// apiRequest represents a typed request to the Chargebee API.
// It is embedded by other concrete types like AddonCreateRequest, CustomerUpdateRequest, etc
type apiRequest[ReqType requestWrapper[ReqType], ResType responseWrapper] struct {
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

type requestWrapper[T any] interface {
	payload() T
}

type BlankRequest[ResType responseWrapper] struct {
	apiRequest[BlankRequest[ResType], ResType]
}

func (r BlankRequest[ResType]) payload() BlankRequest[ResType] {
	return r
}

func NewRequest[ReqType requestWrapper[ReqType], ResType responseWrapper](method, path string, req *ReqType) *apiRequest[ReqType, ResType] {
	return &apiRequest[ReqType, ResType]{
		requestParams: req,
		method:        method,
		path:          path,
	}
}

func NewBlankRequest[ResType responseWrapper]() *apiRequest[BlankRequest[ResType], ResType] {
	return &apiRequest[BlankRequest[ResType], ResType]{
		requestParams: &BlankRequest[ResType]{},
	}
}

func (r *apiRequest[ReqType, ResType]) Bind(owner requestWrapper[ReqType]) {
	payload := owner.payload()
	r.requestParams = &payload
}

func (r *apiRequest[ReqType, ResType]) prepare() error {
	if r.requestParams == nil {
		return fmt.Errorf("requestParams not initialized")
	}
	// if request, ok := any(r).(requestWrapper[ReqType]); ok {
	// 	payload := request.payload()
	// 	if any(payload) != (&BlankRequest[ResType]{}) {
	// 		r.requestParams = &payload
	// 	}
	// } else {
	// 	return fmt.Errorf("cannot get payload from request")
	// }

	// if r.requestParams != nil && (any(*r.requestParams) != (BlankRequest[ResType]{})) {
	if r.requestParams != nil {
		if r.isJsonRequest && strings.ToUpper(r.method) == "POST" {
			if jsonData, err := json.Marshal(r.requestParams); err != nil {
				return fmt.Errorf("failed to marshal params: %w", err)
			} else {
				r.jsonBody = string(jsonData)
			}
		} else if r.isListRequest {
			r.urlParams = SerializeListParams(r.requestParams)
		} else {
			r.urlParams = SerializeParams(r.requestParams)
		}
	}
	return nil
}

func (r *apiRequest[ReqType, ResType]) AddParam(key string, value string) {
	if r.urlParams == nil {
		r.urlParams = &url.Values{}
	}
	r.urlParams.Set(key, value)
}

func (r *apiRequest[ReqType, ResType]) AddHeader(key string, value string) {
	if r.headers == nil {
		r.headers = &http.Header{}
	}
	r.headers.Add(key, value)
}

func (r *apiRequest[ReqType, ResType]) SetIdempotencyKey(idempotencyKey string) {
	r.AddHeader(IdempotencyHeader, idempotencyKey)
}

func (r *apiRequest[ReqType, ResType]) SetSubDomain(subDomain string) {
	r.subDomain = subDomain
}

func (r *apiRequest[ReqType, ResType]) SetIdempotency(idempotent bool) {
	r.isIdempotent = idempotent
}

func (r *apiRequest[ReqType, ResType]) SetContext(ctx context.Context) {
	r.context = ctx
}

// func (r *ChargebeeRequest[ReqType, ResType]) Send() (ResType, error) {
// 	result, err := r.SendWithEnv(DefaultConfig())
// 	return *result, err
// }

func (r *apiRequest[ReqType, ResType]) Send(cfg *ClientConfig) (*ResType, error) {
	err := r.prepare()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %w", err)
	}
	var body io.Reader
	var path string

	if r.isJsonRequest {
		path, body = getJsonBody(r.method, r.path, r.jsonBody)
	} else {
		path, body = getBody(r.method, r.path, r.urlParams)
	}

	reqObj, err := newRequest(cfg, r.method, path, body, r.headers, r.subDomain, r.isJsonRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create new chargebee request: %w", err)
	}

	if r.context != nil {
		reqObj = reqObj.WithContext(context.WithValue(r.context, cbEnvKey, cfg))
	} else {
		reqObj = reqObj.WithContext(context.WithValue(reqObj.Context(), cbEnvKey, cfg))
	}

	res, requestError := Do(reqObj, r.isIdempotent, cfg)
	result := new(ResType)
	if requestError != nil {
		return result, requestError
	}

	if err := UnmarshalJSON(res.body, result); err != nil && res.statusCode != http.StatusNoContent {
		return result, err
	}

	// if resultWithResponseMeta, ok := result.(response); ok {
	// 	resultWithResponseMeta.SetResponseMeta(&ResponseMeta{
	// 		Headers:    res.Headers,
	// 		Status:     res.Status,
	// 		StatusCode: res.StatusCode,
	// 	})
	// }

	(*result).SetMetadata(&apiResponse{
		headers:    res.headers,
		statusText: res.statusText,
		statusCode: res.statusCode,
		body:       res.body,
	})

	return result, requestError
}

func basicAuth(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func newRequest(cc *ClientConfig, method string, path string, body io.Reader, headers *http.Header, subDomain string, isJsonRequest bool) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = cc.apiBaseUrl(subDomain) + path
	httpReq, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	addHeaders(httpReq, cc, isJsonRequest)
	addCustomHeaders(httpReq, headers)
	return httpReq, err
}

func addHeaders(httpReq *http.Request, cc *ClientConfig, isJsonRequest bool) {
	httpReq.Header.Add("Accept-Charset", Charset)
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("User-Agent", "ChargeBee-Go-Client v"+Version)
	httpReq.Header.Add("Authorization", "Basic "+basicAuth(string(cc.ApiKey)))
	httpReq.Header.Add("Lang-Version", runtime.Version())
	httpReq.Header.Add("OS-Version", runtime.GOOS+" "+runtime.GOARCH)
	if isJsonRequest {
		httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	} else {
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
}

func addCustomHeaders(httpReq *http.Request, headers *http.Header) {
	if headers != nil {
		for k, v := range *headers {
			for _, vv := range v {
				httpReq.Header.Add(k, vv)
			}
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
