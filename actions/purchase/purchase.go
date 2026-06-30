package purchase

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/purchase"
)

func Create(params *purchase.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/purchases"), params).WithTelemetryResource("purchase").WithTelemetryOperation("create").SetIdempotency(true)
}
func Estimate(params *purchase.EstimateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/purchases/estimate"), params).WithTelemetryResource("purchase").WithTelemetryOperation("estimate")
}
