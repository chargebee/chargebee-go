package promotionalcredit

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"
	"net/url"
)

func Add(params *promotionalcredit.AddRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/add"), params).WithTelemetryResource("promotionalCredit").WithTelemetryOperation("add").SetIdempotency(true)
}
func Deduct(params *promotionalcredit.DeductRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/deduct"), params).WithTelemetryResource("promotionalCredit").WithTelemetryOperation("deduct").SetIdempotency(true)
}
func Set(params *promotionalcredit.SetRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/promotional_credits/set"), params).WithTelemetryResource("promotionalCredit").WithTelemetryOperation("set").SetIdempotency(true)
}
func List(params *promotionalcredit.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/promotional_credits"), params).WithTelemetryResource("promotionalCredit").WithTelemetryOperation("list")
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/promotional_credits/%v", url.PathEscape(id)), nil).WithTelemetryResource("promotionalCredit").WithTelemetryOperation("retrieve")
}
