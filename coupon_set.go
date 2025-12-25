package chargebee

import (
	"encoding/json"
)

// just struct
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

// sub resources
// operations
// input params
type CouponSetCreateRequest struct {
	CouponId   string                 `json:"coupon_id"`
	Name       string                 `json:"name"`
	Id         string                 `json:"id"`
	MetaData   map[string]interface{} `json:"meta_data,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CouponSetCreateRequest) payload() any { return r }

type CouponSetAddCouponCodesRequest struct {
	Code       []string `json:"code,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CouponSetAddCouponCodesRequest) payload() any { return r }

type CouponSetListRequest struct {
	Limit         *int32        `json:"limit,omitempty"`
	Offset        string        `json:"offset,omitempty"`
	Id            *StringFilter `json:"id,omitempty"`
	Name          *StringFilter `json:"name,omitempty"`
	CouponId      *StringFilter `json:"coupon_id,omitempty"`
	TotalCount    *NumberFilter `json:"total_count,omitempty"`
	RedeemedCount *NumberFilter `json:"redeemed_count,omitempty"`
	ArchivedCount *NumberFilter `json:"archived_count,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *CouponSetListRequest) payload() any { return r }

type CouponSetUpdateRequest struct {
	Name       string                 `json:"name,omitempty"`
	MetaData   map[string]interface{} `json:"meta_data,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CouponSetUpdateRequest) payload() any { return r }

// operation response
type CouponSetCreateResponse struct {
	CouponSet *CouponSet `json:"coupon_set,omitempty"`
	apiResponse
}

// operation response
type CouponSetAddCouponCodesResponse struct {
	CouponSet *CouponSet `json:"coupon_set,omitempty"`
	apiResponse
}

// operation sub response
type CouponSetListCouponSetResponse struct {
	CouponSet *CouponSet `json:"coupon_set,omitempty"`
}

// operation response
type CouponSetListResponse struct {
	List       []*CouponSetListCouponSetResponse `json:"list,omitempty"`
	NextOffset string                            `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CouponSetRetrieveResponse struct {
	CouponSet *CouponSet `json:"coupon_set,omitempty"`
	apiResponse
}

// operation response
type CouponSetUpdateResponse struct {
	CouponSet *CouponSet `json:"coupon_set,omitempty"`
	apiResponse
}

// operation response
type CouponSetDeleteResponse struct {
	CouponSet *CouponSet `json:"coupon_set,omitempty"`
	apiResponse
}

// operation response
type CouponSetDeleteUnusedCouponCodesResponse struct {
	CouponSet *CouponSet `json:"coupon_set,omitempty"`
	apiResponse
}
