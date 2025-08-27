package couponcode

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/couponcode"
	"net/url"
)

func Create(params *couponcode.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_codes"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/coupon_codes/%v", url.PathEscape(id)), nil)
}
func List(params *couponcode.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/coupon_codes"), params)
}
func Archive(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_codes/%v/archive", url.PathEscape(id)), nil).SetIdempotency(true)
}
