package asyncresponse

import (
	"encoding/json"
	asyncResponseEnum "github.com/chargebee/chargebee-go/v3/models/asyncresponse/enum"
)

type AsyncResponse struct {
	ApiVersion  string                   `json:"api_version"`
	CreatedAt   int64                    `json:"created_at"`
	CompletedAt int64                    `json:"completed_at"`
	Status      asyncResponseEnum.Status `json:"status"`
	Request     *RequestAsyncApi         `json:"request"`
	ErrorDetail *Error                   `json:"error_detail"`
	Result      json.RawMessage          `json:"result"`
	Object      string                   `json:"object"`
}
type RequestAsyncApi struct {
	Id             string `json:"id"`
	Resource       string `json:"resource"`
	OperationType  string `json:"operation_type"`
	Method         string `json:"method"`
	Uri            string `json:"uri"`
	IdempotencyKey string `json:"idempotency_key"`
	Object         string `json:"object"`
}
type Error struct {
	Message        string `json:"message"`
	Type           string `json:"type"`
	ApiErrorCode   string `json:"api_error_code"`
	ErrorCode      string `json:"error_code"`
	ErrorMsg       string `json:"error_msg"`
	HttpStatusCode string `json:"http_status_code"`
	Object         string `json:"object"`
}
