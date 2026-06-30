package alert

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/alert"
	"net/url"
)

func Create(params *alert.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/alerts"), params).WithTelemetryResource("alert").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/alerts/%v", url.PathEscape(id)), nil).WithTelemetryResource("alert").WithTelemetryOperation("retrieve")
}
func List(params *alert.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/alerts"), params).WithTelemetryResource("alert").WithTelemetryOperation("list")
}
func Update(id string, params *alert.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/alerts/%v", url.PathEscape(id)), params).WithTelemetryResource("alert").WithTelemetryOperation("update").SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/alerts/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("alert").WithTelemetryOperation("delete").SetIdempotency(true)
}
func ApplicationAlertsForSubscription(id string, params *alert.ApplicationAlertsForSubscriptionRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions/%v/applicable_alerts", url.PathEscape(id)), params).WithTelemetryResource("alert").WithTelemetryOperation("applicationAlertsForSubscription")
}
