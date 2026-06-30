package coupon

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/coupon"
	"net/url"
)

func Create(params *coupon.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupons"), params).WithTelemetryResource("coupon").WithTelemetryOperation("create").SetIdempotency(true)
}
func CreateForItems(params *coupon.CreateForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/create_for_items"), params).WithTelemetryResource("coupon").WithTelemetryOperation("createForItems").SetIdempotency(true)
}
func UpdateForItems(id string, params *coupon.UpdateForItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/update_for_items", url.PathEscape(id)), params).WithTelemetryResource("coupon").WithTelemetryOperation("updateForItems").SetIdempotency(true)
}
func List(params *coupon.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/coupons"), params).WithTelemetryResource("coupon").WithTelemetryOperation("list")
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), nil).WithTelemetryResource("coupon").WithTelemetryOperation("retrieve")
}
func Update(id string, params *coupon.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), params).WithTelemetryResource("coupon").WithTelemetryOperation("update").SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("coupon").WithTelemetryOperation("delete").SetIdempotency(true)
}
func Copy(params *coupon.CopyRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/copy"), params).WithTelemetryResource("coupon").WithTelemetryOperation("copy").SetIdempotency(true)
}
func Unarchive(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/unarchive", url.PathEscape(id)), nil).WithTelemetryResource("coupon").WithTelemetryOperation("unarchive").SetIdempotency(true)
}
