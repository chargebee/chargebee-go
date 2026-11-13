package gatewayrequestlog

import (
	"github.com/chargebee/chargebee-go/v3/filter"
)

type GatewayRequestLog struct {
	LogId                   string `json:"log_id"`
	Url                     string `json:"url"`
	StatusCode              string `json:"status_code"`
	Method                  string `json:"method"`
	RequestBody             string `json:"request_body"`
	ResponseBody            string `json:"response_body"`
	RequestHeaders          string `json:"request_headers"`
	ResponseHeaders         string `json:"response_headers"`
	Gateway                 string `json:"gateway"`
	CustomerHandle          string `json:"customer_handle"`
	InvoiceHandle           string `json:"invoice_handle"`
	TransactionHandle       string `json:"transaction_handle"`
	PaymentMethodHandle     string `json:"payment_method_handle"`
	GatewayCredentialHandle string `json:"gateway_credential_handle"`
	PaymentAttemptHandle    string `json:"payment_attempt_handle"`
	PaymentIntentHandle     string `json:"payment_intent_handle"`
	PaymentMethodName       string `json:"payment_method_name"`
	OperationName           string `json:"operation_name"`
	OperationType           string `json:"operation_type"`
	ActionType              string `json:"action_type"`
	Source                  string `json:"source"`
	TraceIdHash             string `json:"trace_id_hash"`
	InvocationTime          int64  `json:"invocation_time"`
	ResponseReceivedAt      int64  `json:"response_received_at"`
	ResponseTime            int64  `json:"response_time"`
	Object                  string `json:"object"`
}
type ListRequestParams struct {
	Limit                   *int32                  `json:"limit,omitempty"`
	Offset                  string                  `json:"offset,omitempty"`
	LogId                   *filter.StringFilter    `json:"log_id,omitempty"`
	Gateway                 *filter.StringFilter    `json:"gateway,omitempty"`
	StatusCode              *filter.StringFilter    `json:"status_code,omitempty"`
	Method                  *filter.StringFilter    `json:"method,omitempty"`
	Url                     *filter.StringFilter    `json:"url,omitempty"`
	CustomerHandle          *filter.StringFilter    `json:"customer_handle,omitempty"`
	InvoiceHandle           *filter.StringFilter    `json:"invoice_handle,omitempty"`
	TransactionHandle       *filter.StringFilter    `json:"transaction_handle,omitempty"`
	PaymentMethodHandle     *filter.StringFilter    `json:"payment_method_handle,omitempty"`
	GatewayCredentialHandle *filter.StringFilter    `json:"gateway_credential_handle,omitempty"`
	PaymentAttemptHandle    *filter.StringFilter    `json:"payment_attempt_handle,omitempty"`
	PaymentIntentHandle     *filter.StringFilter    `json:"payment_intent_handle,omitempty"`
	OperationName           *filter.StringFilter    `json:"operation_name,omitempty"`
	OperationType           *filter.StringFilter    `json:"operation_type,omitempty"`
	Source                  *filter.StringFilter    `json:"source,omitempty"`
	InvocationTime          *filter.TimestampFilter `json:"invocation_time,omitempty"`
	ResponseReceivedAt      *filter.TimestampFilter `json:"response_received_at,omitempty"`
	SortBy                  *filter.SortFilter      `json:"sort_by,omitempty"`
}
