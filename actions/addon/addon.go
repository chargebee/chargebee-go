package addon

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/addon"
	"net/url"
)

func Create(params *addon.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons"), params).WithTelemetryResource("addon").WithTelemetryOperation("create").SetIdempotency(true)
}
func Update(id string, params *addon.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v", url.PathEscape(id)), params).WithTelemetryResource("addon").WithTelemetryOperation("update").SetIdempotency(true)
}
func List(params *addon.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/addons"), params).WithTelemetryResource("addon").WithTelemetryOperation("list")
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/addons/%v", url.PathEscape(id)), nil).WithTelemetryResource("addon").WithTelemetryOperation("retrieve")
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("addon").WithTelemetryOperation("delete").SetIdempotency(true)
}
func Copy(params *addon.CopyRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/copy"), params).WithTelemetryResource("addon").WithTelemetryOperation("copy").SetIdempotency(true)
}
func Unarchive(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/addons/%v/unarchive", url.PathEscape(id)), nil).WithTelemetryResource("addon").WithTelemetryOperation("unarchive").SetIdempotency(true)
}
