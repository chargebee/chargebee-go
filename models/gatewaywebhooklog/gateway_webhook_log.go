package gatewaywebhooklog

import (
	"github.com/chargebee/chargebee-go/v3/filter"
)

type GatewayWebhookLog struct {
	WebhookId               string `json:"webhook_id"`
	Url                     string `json:"url"`
	Method                  string `json:"method"`
	StatusCode              string `json:"status_code"`
	Gateway                 string `json:"gateway"`
	IdAtGateway             string `json:"id_at_gateway"`
	CustomerHandle          string `json:"customer_handle"`
	InvoiceHandle           string `json:"invoice_handle"`
	TransactionHandle       string `json:"transaction_handle"`
	PaymentMethodHandle     string `json:"payment_method_handle"`
	GatewayCredentialHandle string `json:"gateway_credential_handle"`
	PaymentAttemptHandle    string `json:"payment_attempt_handle"`
	PaymentIntentHandle     string `json:"payment_intent_handle"`
	PaymentMethodName       string `json:"payment_method_name"`
	RequestBody             string `json:"request_body"`
	ResponseBody            string `json:"response_body"`
	RequestHeaders          string `json:"request_headers"`
	ResponseHeaders         string `json:"response_headers"`
	ReceivedAt              int64  `json:"received_at"`
	ResponseTime            int64  `json:"response_time"`
	Object                  string `json:"object"`
}
type ListRequestParams struct {
	Limit                   *int32                  `json:"limit,omitempty"`
	Offset                  string                  `json:"offset,omitempty"`
	WebhookId               *filter.StringFilter    `json:"webhook_id,omitempty"`
	Gateway                 *filter.StringFilter    `json:"gateway,omitempty"`
	StatusCode              *filter.StringFilter    `json:"status_code,omitempty"`
	Method                  *filter.StringFilter    `json:"method,omitempty"`
	Url                     *filter.StringFilter    `json:"url,omitempty"`
	IdAtGateway             *filter.StringFilter    `json:"id_at_gateway,omitempty"`
	CustomerHandle          *filter.StringFilter    `json:"customer_handle,omitempty"`
	InvoiceHandle           *filter.StringFilter    `json:"invoice_handle,omitempty"`
	TransactionHandle       *filter.StringFilter    `json:"transaction_handle,omitempty"`
	PaymentMethodHandle     *filter.StringFilter    `json:"payment_method_handle,omitempty"`
	GatewayCredentialHandle *filter.StringFilter    `json:"gateway_credential_handle,omitempty"`
	PaymentAttemptHandle    *filter.StringFilter    `json:"payment_attempt_handle,omitempty"`
	PaymentIntentHandle     *filter.StringFilter    `json:"payment_intent_handle,omitempty"`
	ReceivedAt              *filter.TimestampFilter `json:"received_at,omitempty"`
	SortBy                  *filter.SortFilter      `json:"sort_by,omitempty"`
}
