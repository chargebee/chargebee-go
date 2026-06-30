package couponset

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/couponset"
	"net/url"
)

func Create(params *couponset.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets"), params).WithTelemetryResource("couponSet").WithTelemetryOperation("create").SetIdempotency(true)
}
func AddCouponCodes(id string, params *couponset.AddCouponCodesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/add_coupon_codes", url.PathEscape(id)), params).WithTelemetryResource("couponSet").WithTelemetryOperation("addCouponCodes").SetIdempotency(true)
}
func List(params *couponset.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/coupon_sets"), params).WithTelemetryResource("couponSet").WithTelemetryOperation("list")
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/coupon_sets/%v", url.PathEscape(id)), nil).WithTelemetryResource("couponSet").WithTelemetryOperation("retrieve")
}
func Update(id string, params *couponset.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/update", url.PathEscape(id)), params).WithTelemetryResource("couponSet").WithTelemetryOperation("update").SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/delete", url.PathEscape(id)), nil).WithTelemetryResource("couponSet").WithTelemetryOperation("delete").SetIdempotency(true)
}
func DeleteUnusedCouponCodes(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/delete_unused_coupon_codes", url.PathEscape(id)), nil).WithTelemetryResource("couponSet").WithTelemetryOperation("deleteUnusedCouponCodes").SetIdempotency(true)
}
