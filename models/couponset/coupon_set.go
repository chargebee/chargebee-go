package couponset

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/filter"
)

type CouponSet struct {
	Id            string          `json:"id"`
	CouponId      string          `json:"coupon_id"`
	Name          string          `json:"name"`
	TotalCount    int32           `json:"total_count"`
	RedeemedCount int32           `json:"redeemed_count"`
	ArchivedCount int32           `json:"archived_count"`
	MetaData      json.RawMessage `json:"meta_data"`
	Object        string          `json:"object"`
}
type CreateRequestParams struct {
	CouponId string                 `json:"coupon_id"`
	Name     string                 `json:"name"`
	Id       string                 `json:"id"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
}
type AddCouponCodesRequestParams struct {
	Code []string `json:"code,omitempty"`
}
type ListRequestParams struct {
	Limit         *int32               `json:"limit,omitempty"`
	Offset        string               `json:"offset,omitempty"`
	Id            *filter.StringFilter `json:"id,omitempty"`
	Name          *filter.StringFilter `json:"name,omitempty"`
	CouponId      *filter.StringFilter `json:"coupon_id,omitempty"`
	TotalCount    *filter.NumberFilter `json:"total_count,omitempty"`
	RedeemedCount *filter.NumberFilter `json:"redeemed_count,omitempty"`
	ArchivedCount *filter.NumberFilter `json:"archived_count,omitempty"`
}
type UpdateRequestParams struct {
	Name     string                 `json:"name,omitempty"`
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
}
