package couponcode

import (
	"github.com/chargebee/chargebee-go/filter"
	couponCodeEnum "github.com/chargebee/chargebee-go/models/couponcode/enum"
)

type CouponCode struct {
	Code          string                `json:"code"`
	Status        couponCodeEnum.Status `json:"status"`
	CouponId      string                `json:"coupon_id"`
	CouponSetId   string                `json:"coupon_set_id"`
	CouponSetName string                `json:"coupon_set_name"`
	Object        string                `json:"object"`
}
type CreateRequestParams struct {
	CouponId      string `json:"coupon_id"`
	CouponSetName string `json:"coupon_set_name"`
	Code          string `json:"code"`
}
type ListRequestParams struct {
	Limit         *int32               `json:"limit,omitempty"`
	Offset        string               `json:"offset,omitempty"`
	Code          *filter.StringFilter `json:"code,omitempty"`
	CouponId      *filter.StringFilter `json:"coupon_id,omitempty"`
	CouponSetName *filter.StringFilter `json:"coupon_set_name,omitempty"`
	Status        *filter.EnumFilter   `json:"status,omitempty"`
}
