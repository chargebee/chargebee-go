package chargebee

type Status string

const (
	StatusNotRedeemed Status = "not_redeemed"
	StatusRedeemed    Status = "redeemed"
	StatusArchived    Status = "archived"
)

type CouponCode struct {
	Code          string `json:"code"`
	Status        Status `json:"status"`
	CouponId      string `json:"coupon_id"`
	CouponSetId   string `json:"coupon_set_id"`
	CouponSetName string `json:"coupon_set_name"`
	Object        string `json:"object"`
}
type CreateRequest struct {
	CouponId      string `json:"coupon_id"`
	CouponSetName string `json:"coupon_set_name"`
	Code          string `json:"code"`
}
type ListRequest struct {
	Limit         *int32               `json:"limit,omitempty"`
	Offset        string               `json:"offset,omitempty"`
	Code          *filter.StringFilter `json:"code,omitempty"`
	CouponId      *filter.StringFilter `json:"coupon_id,omitempty"`
	CouponSetName *filter.StringFilter `json:"coupon_set_name,omitempty"`
	Status        *filter.EnumFilter   `json:"status,omitempty"`
}

type CreateResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
}

type RetrieveResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
}

type ListCouponCodeResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
}

type ListResponse struct {
	List       []*ListCouponCodeResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
}

type ArchiveResponse struct {
	CouponCode *CouponCode `json:"coupon_code,omitempty"`
}
