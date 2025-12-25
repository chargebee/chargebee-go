package chargebee

type CouponCodeStatus string

const (
	CouponCodeStatusNotRedeemed CouponCodeStatus = "not_redeemed"
	CouponCodeStatusRedeemed    CouponCodeStatus = "redeemed"
	CouponCodeStatusArchived    CouponCodeStatus = "archived"
)

// just struct
type CouponCode struct {
	Code          string           `json:"code"`
	Status        CouponCodeStatus `json:"status"`
	CouponId      string           `json:"coupon_id"`
	CouponSetId   string           `json:"coupon_set_id"`
	CouponSetName string           `json:"coupon_set_name"`
	Object        string           `json:"object"`
}

// sub resources
// operations
// input params
type CouponCodeCreateRequest struct {
	CouponId      string `json:"coupon_id"`
	CouponSetName string `json:"coupon_set_name"`
	Code          string `json:"code"`
	apiRequest    `json:"-" form:"-"`
}

func (r *CouponCodeCreateRequest) payload() any { return r }

type CouponCodeListRequest struct {
	Limit         *int32        `json:"limit,omitempty"`
	Offset        string        `json:"offset,omitempty"`
	Code          *StringFilter `json:"code,omitempty"`
	CouponId      *StringFilter `json:"coupon_id,omitempty"`
	CouponSetName *StringFilter `json:"coupon_set_name,omitempty"`
	Status        *EnumFilter   `json:"status,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *CouponCodeListRequest) payload() any { return r }

// operation response
type CouponCodeCreateResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
	apiResponse
}

// operation response
type CouponCodeRetrieveResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
	apiResponse
}

// operation sub response
type CouponCodeListCouponCodeResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
}

// operation response
type CouponCodeListResponse struct {
	List       []*CouponCodeListCouponCodeResponse `json:"list,omitempty"`
	NextOffset string                              `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CouponCodeArchiveResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
	apiResponse
}
