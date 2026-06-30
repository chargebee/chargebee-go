package itemfamily

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/itemfamily"
	"net/url"
)

func Create(params *itemfamily.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/item_families"), params).WithTelemetryResource("itemFamily").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/item_families/%v", url.PathEscape(id)), nil).WithTelemetryResource("itemFamily").WithTelemetryOperation("retrieve")
}
func List(params *itemfamily.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/item_families"), params).WithTelemetryResource("itemFamily").WithTelemetryOperation("list")
}
func Update(id string, params *itemfamily.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/item_families/%v", url.PathEscape(id)), params).WithTelemetryResource("itemFamily").WithTelemetryOperation("update").SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/item_families/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("itemFamily").WithTelemetryOperation("delete").SetIdempotency(true)
}
