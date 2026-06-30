package comment

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/comment"
	"net/url"
)

func Create(params *comment.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/comments"), params).WithTelemetryResource("comment").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/comments/%v", url.PathEscape(id)), nil).WithTelemetryResource("comment").WithTelemetryOperation("retrieve")
}
func List(params *comment.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/comments"), params).WithTelemetryResource("comment").WithTelemetryOperation("list")
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/comments/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("comment").WithTelemetryOperation("delete").SetIdempotency(true)
}
