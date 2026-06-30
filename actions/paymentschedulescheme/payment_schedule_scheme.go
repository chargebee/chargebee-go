package paymentschedulescheme

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/paymentschedulescheme"
	"net/url"
)

func Create(params *paymentschedulescheme.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_schedule_schemes"), params).WithTelemetryResource("paymentScheduleScheme").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/payment_schedule_schemes/%v", url.PathEscape(id)), nil).WithTelemetryResource("paymentScheduleScheme").WithTelemetryOperation("retrieve")
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_schedule_schemes/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("paymentScheduleScheme").WithTelemetryOperation("delete").SetIdempotency(true)
}
