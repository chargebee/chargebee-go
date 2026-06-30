package webhookendpoint

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/webhookendpoint"
	"net/url"
)

func Create(params *webhookendpoint.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/webhook_endpoints"), params).WithTelemetryResource("webhookEndpoint").WithTelemetryOperation("create").SetIdempotency(true)
}
func Update(id string, params *webhookendpoint.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/webhook_endpoints/%v", url.PathEscape(id)), params).WithTelemetryResource("webhookEndpoint").WithTelemetryOperation("update").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/webhook_endpoints/%v", url.PathEscape(id)), nil).WithTelemetryResource("webhookEndpoint").WithTelemetryOperation("retrieve")
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/webhook_endpoints/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("webhookEndpoint").WithTelemetryOperation("delete").SetIdempotency(true)
}
func List(params *webhookendpoint.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/webhook_endpoints"), params).WithTelemetryResource("webhookEndpoint").WithTelemetryOperation("list")
}
