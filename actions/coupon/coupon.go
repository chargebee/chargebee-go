package coupon

import (
	"fmt"
	"net/url"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/coupon"
)

func Create(params *coupon.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons"), params)
}
func List(params *coupon.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/coupons"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *coupon.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/delete", url.PathEscape(id)), nil)
}
func Copy(params *coupon.CopyRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/copy"), params)
}
func Unarchive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/unarchive", url.PathEscape(id)), nil)
}
