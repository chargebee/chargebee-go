package gift

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/gift"
	"net/url"
)

func Create(params *gift.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/gifts"), params).WithTelemetryResource("gift").WithTelemetryOperation("create").SetIdempotency(true)
}
func CreateForItems(params *gift.CreateForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/create_for_items"), params).WithTelemetryResource("gift").WithTelemetryOperation("createForItems").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/gifts/%v", url.PathEscape(id)), nil).WithTelemetryResource("gift").WithTelemetryOperation("retrieve")
}
func List(params *gift.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/gifts"), params).WithTelemetryResource("gift").WithTelemetryOperation("list")
}
func Claim(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/claim", url.PathEscape(id)), nil).WithTelemetryResource("gift").WithTelemetryOperation("claim").SetIdempotency(true)
}
func Cancel(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/cancel", url.PathEscape(id)), nil).WithTelemetryResource("gift").WithTelemetryOperation("cancel").SetIdempotency(true)
}
func UpdateGift(id string, params *gift.UpdateGiftRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/update_gift", url.PathEscape(id)), params).WithTelemetryResource("gift").WithTelemetryOperation("updateGift").SetIdempotency(true)
}
