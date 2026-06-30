package usage

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/usage"
	"net/url"
)

func Create(id string, params *usage.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/usages", url.PathEscape(id)), params).WithTelemetryResource("usage").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string, params *usage.RetrieveRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/usages", url.PathEscape(id)), params).WithTelemetryResource("usage").WithTelemetryOperation("retrieve")
}
func Delete(id string, params *usage.DeleteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/delete_usage", url.PathEscape(id)), params).WithTelemetryResource("usage").WithTelemetryOperation("delete").SetIdempotency(true)
}
func List(params *usage.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/usages"), params).WithTelemetryResource("usage").WithTelemetryOperation("list")
}
func Pdf(params *usage.PdfRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/usages/pdf"), params).WithTelemetryResource("usage").WithTelemetryOperation("pdf").SetIdempotency(true)
}
