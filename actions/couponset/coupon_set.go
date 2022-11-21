package couponset

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/couponset"
	"net/url"
)

func Create(params *couponset.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets"), params)
}
func AddCouponCodes(id string, params *couponset.AddCouponCodesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/add_coupon_codes", url.PathEscape(id)), params)
}
func List(params *couponset.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/coupon_sets"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/coupon_sets/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *couponset.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/update", url.PathEscape(id)), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/delete", url.PathEscape(id)), nil)
}
func DeleteUnusedCouponCodes(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_sets/%v/delete_unused_coupon_codes", url.PathEscape(id)), nil)
}
