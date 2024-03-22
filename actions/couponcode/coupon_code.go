package couponcode

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/couponcode"
)

func Create(params *couponcode.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_codes"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/coupon_codes/%v", chargebee.IDEscape(id)), nil)
}
func List(params *couponcode.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/coupon_codes"), params)
}
func Archive(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/coupon_codes/%v/archive", chargebee.IDEscape(id)), nil)
}
