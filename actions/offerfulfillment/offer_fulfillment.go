package offerfulfillment

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/offerfulfillment"
	"net/url"
)

func OfferFulfillments(params *offerfulfillment.OfferFulfillmentsRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/offer_fulfillments"), params).WithTelemetryResource("offerFulfillment").WithTelemetryOperation("offerFulfillments").SetSubDomain("grow")
}
func OfferFulfillmentsGet(id string) chargebee.Request {
	return chargebee.SendJsonRequest("GET", fmt.Sprintf("/offer_fulfillments/%v", url.PathEscape(id)), nil).WithTelemetryResource("offerFulfillment").WithTelemetryOperation("offerFulfillmentsGet").SetSubDomain("grow")
}
func OfferFulfillmentsUpdate(id string, params *offerfulfillment.OfferFulfillmentsUpdateRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/offer_fulfillments/%v", url.PathEscape(id)), params).WithTelemetryResource("offerFulfillment").WithTelemetryOperation("offerFulfillmentsUpdate").SetSubDomain("grow")
}
