package paymentintent

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/paymentintent"
	"net/url"
)

func Create(params *paymentintent.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_intents"), params).WithTelemetryResource("paymentIntent").WithTelemetryOperation("create").SetIdempotency(true)
}
func Update(id string, params *paymentintent.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/payment_intents/%v", url.PathEscape(id)), params).WithTelemetryResource("paymentIntent").WithTelemetryOperation("update").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/payment_intents/%v", url.PathEscape(id)), nil).WithTelemetryResource("paymentIntent").WithTelemetryOperation("retrieve")
}
