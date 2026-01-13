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

// apiRequest represents request metadata like headers, context, etc
// It is embedded by other concrete types like AddonCreateRequest,
// CustomerUpdateRequest, etc
type apiRequest struct {
	method    string          `json:"-" form:"-"`
	path      string          `json:"-" form:"-"`
	headers   *http.Header    `json:"-" form:"-"`
	context   context.Context `json:"-" form:"-"`
	subDomain string          `json:"-" form:"-"`
	// request payload for all url encoded requests
	urlParams *url.Values `json:"-" form:"-"`
	// request payload for json requests
	jsonBody      string `json:"-" form:"-"`
	isJsonRequest bool   `json:"-" form:"-"`
	isListRequest bool   `json:"-" form:"-"`
	isIdempotent  bool   `json:"-" form:"-"`
}

type requestWrapper interface {
	request() *apiRequest
	payload() any
}

func (r *apiRequest) request() *apiRequest {
	return r
}

type BlankRequest struct {
	apiRequest `json:"-" form:"-"`
}

func (r *BlankRequest) payload() any {
	return r
}

func (r *apiRequest) AddParam(key string, value string) {
	if r.urlParams == nil {
		r.urlParams = &url.Values{}
	}
	r.urlParams.Set(key, value)
}

func (r *apiRequest) AddCustomField(key, value string) {
	if !strings.HasPrefix(key, "cf_") {
		key = "cf_" + key
	}
	r.AddParam(key, value)
}

func (r *apiRequest) AddHeader(key string, value string) {
	if r.headers == nil {
		r.headers = &http.Header{}
	}
	r.headers.Add(key, value)
}

func (r *apiRequest) SetIdempotencyKey(idempotencyKey string) {
	r.AddHeader(IdempotencyHeader, idempotencyKey)
}

func (r *apiRequest) SetSubDomain(subDomain string) {
	r.subDomain = subDomain
}

func (r *apiRequest) SetIdempotency(idempotent bool) {
	r.isIdempotent = idempotent
}

func (r *apiRequest) SetContext(ctx context.Context) {
	r.context = ctx
}

func (r *apiRequest) prepare(payload any) error {
	if r == nil {
		return fmt.Errorf("request is nil")
	}
	if payload != nil {
		if r.isJsonRequest && strings.ToUpper(r.method) == "POST" {
			if jsonData, err := json.Marshal(payload); err != nil {
				return fmt.Errorf("failed to marshal payload: %w", err)
			} else {
				r.jsonBody = string(jsonData)
			}
		} else if r.isListRequest {
			r.urlParams = SerializeListParams(payload)
		} else {
			r.urlParams = SerializeParams(payload)
		}
	}
	return nil
}

func send[ResType responseWrapper](rw requestWrapper, cfg *ClientConfig) (ResType, error) {
	var result ResType
	req := rw.request()
	if err := req.prepare(rw.payload()); err != nil {
		return result, fmt.Errorf("failed to prepare request for sending: %w", err)
	}

	var body io.Reader
	var path string

	if req.isJsonRequest {
		path, body = getJsonBody(req.method, req.path, req.jsonBody)
	} else {
		path, body = getBody(req.method, req.path, req.urlParams)
	}

	reqObj, err := newRequest(cfg, req.method, path, body, req.headers, req.subDomain, req.isJsonRequest)
	if err != nil {
		return result, fmt.Errorf("failed to create new chargebee request: %w", err)
	}

	if req.context != nil {
		reqObj = reqObj.WithContext(context.WithValue(req.context, configCtxKey, cfg))
	} else {
		reqObj = reqObj.WithContext(context.WithValue(reqObj.Context(), configCtxKey, cfg))
	}

	res, requestError := Do(reqObj, req.isIdempotent, cfg)
	if requestError != nil {
		return result, requestError
	}

	if err := json.Unmarshal(res.Body, &result); err != nil && res.StatusCode != http.StatusNoContent {
		return result, err
	}

	result.setMeta(&apiResponse{
		Headers:    res.Headers,
		StatusText: res.StatusText,
		StatusCode: res.StatusCode,
		Body:       res.Body,
	})

	return result, requestError
}

func basicAuth(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func newRequest(cfg *ClientConfig, method string, path string, body io.Reader, headers *http.Header, subDomain string, isJsonRequest bool) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = cfg.apiBaseUrl(subDomain) + path
	httpReq, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	addHeaders(httpReq, cfg, isJsonRequest)
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

// GetMap is used to unmarshal the json.RawMessage to map[string]interface{}.
func GetMap(rawMessage json.RawMessage) (map[string]any, error) {
	data := json.NewDecoder(strings.NewReader(string(rawMessage)))
	data.UseNumber()
	var m map[string]any
	if err := data.Decode(&m); err != nil {
		return nil, fmt.Errorf("failed to decode to map[string]any: %w", err)
	}
	return m, nil
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
