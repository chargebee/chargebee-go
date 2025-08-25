package coupon

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/coupon"
	"net/url"
)

func Create(params *coupon.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons"), params).SetIdempotency(true)
}
func CreateForItems(params *coupon.CreateForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/create_for_items"), params).SetIdempotency(true)
}
func UpdateForItems(id string, params *coupon.UpdateForItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/update_for_items", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *coupon.ListRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/coupons"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *coupon.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Copy(params *coupon.CopyRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/copy"), params).SetIdempotency(true)
}
func Unarchive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupons/%v/unarchive", url.PathEscape(id)), nil).SetIdempotency(true)
}
