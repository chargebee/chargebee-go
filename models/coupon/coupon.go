package coupon

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/filter"
	couponEnum "github.com/chargebee/chargebee-go/models/coupon/enum"
)

type Coupon struct {
	Id                 string                     `json:"id"`
	Name               string                     `json:"name"`
	InvoiceName        string                     `json:"invoice_name"`
	DiscountType       couponEnum.DiscountType    `json:"discount_type"`
	DiscountPercentage float64                    `json:"discount_percentage"`
	DiscountAmount     int32                      `json:"discount_amount"`
	DiscountQuantity   int32                      `json:"discount_quantity"`
	CurrencyCode       string                     `json:"currency_code"`
	DurationType       couponEnum.DurationType    `json:"duration_type"`
	DurationMonth      int32                      `json:"duration_month"`
	ValidTill          int64                      `json:"valid_till"`
	MaxRedemptions     int32                      `json:"max_redemptions"`
	Status             couponEnum.Status          `json:"status"`
	ApplyDiscountOn    couponEnum.ApplyDiscountOn `json:"apply_discount_on"`
	ApplyOn            couponEnum.ApplyOn         `json:"apply_on"`
	PlanConstraint     couponEnum.PlanConstraint  `json:"plan_constraint"`
	AddonConstraint    couponEnum.AddonConstraint `json:"addon_constraint"`
	CreatedAt          int64                      `json:"created_at"`
	ArchivedAt         int64                      `json:"archived_at"`
	ResourceVersion    int64                      `json:"resource_version"`
	UpdatedAt          int64                      `json:"updated_at"`
	PlanIds            []string                   `json:"plan_ids"`
	AddonIds           []string                   `json:"addon_ids"`
	Redemptions        int32                      `json:"redemptions"`
	InvoiceNotes       string                     `json:"invoice_notes"`
	MetaData           json.RawMessage            `json:"meta_data"`
	Object             string                     `json:"object"`
}
type CreateRequestParams struct {
	Id                 string                     `json:"id"`
	Name               string                     `json:"name"`
	InvoiceName        string                     `json:"invoice_name,omitempty"`
	DiscountType       couponEnum.DiscountType    `json:"discount_type"`
	DiscountAmount     *int32                     `json:"discount_amount,omitempty"`
	CurrencyCode       string                     `json:"currency_code,omitempty"`
	DiscountPercentage *float64                   `json:"discount_percentage,omitempty"`
	DiscountQuantity   *int32                     `json:"discount_quantity,omitempty"`
	ApplyOn            couponEnum.ApplyOn         `json:"apply_on"`
	PlanConstraint     couponEnum.PlanConstraint  `json:"plan_constraint,omitempty"`
	AddonConstraint    couponEnum.AddonConstraint `json:"addon_constraint,omitempty"`
	PlanIds            []string                   `json:"plan_ids,omitempty"`
	AddonIds           []string                   `json:"addon_ids,omitempty"`
	DurationType       couponEnum.DurationType    `json:"duration_type"`
	DurationMonth      *int32                     `json:"duration_month,omitempty"`
	ValidTill          *int64                     `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                     `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                     `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{}     `json:"meta_data,omitempty"`
	Status             couponEnum.Status          `json:"status,omitempty"`
}
type ListRequestParams struct {
	Limit        *int32                  `json:"limit,omitempty"`
	Offset       string                  `json:"offset,omitempty"`
	Id           *filter.StringFilter    `json:"id,omitempty"`
	Name         *filter.StringFilter    `json:"name,omitempty"`
	DiscountType *filter.EnumFilter      `json:"discount_type,omitempty"`
	DurationType *filter.EnumFilter      `json:"duration_type,omitempty"`
	Status       *filter.EnumFilter      `json:"status,omitempty"`
	ApplyOn      *filter.EnumFilter      `json:"apply_on,omitempty"`
	CreatedAt    *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt    *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy       *filter.SortFilter      `json:"sort_by,omitempty"`
}
type UpdateRequestParams struct {
	Name               string                     `json:"name,omitempty"`
	InvoiceName        string                     `json:"invoice_name,omitempty"`
	DiscountType       couponEnum.DiscountType    `json:"discount_type,omitempty"`
	DiscountAmount     *int32                     `json:"discount_amount,omitempty"`
	CurrencyCode       string                     `json:"currency_code,omitempty"`
	DiscountPercentage *float64                   `json:"discount_percentage,omitempty"`
	ApplyOn            couponEnum.ApplyOn         `json:"apply_on,omitempty"`
	PlanConstraint     couponEnum.PlanConstraint  `json:"plan_constraint,omitempty"`
	AddonConstraint    couponEnum.AddonConstraint `json:"addon_constraint,omitempty"`
	PlanIds            []string                   `json:"plan_ids,omitempty"`
	AddonIds           []string                   `json:"addon_ids,omitempty"`
	DurationType       couponEnum.DurationType    `json:"duration_type,omitempty"`
	DurationMonth      *int32                     `json:"duration_month,omitempty"`
	ValidTill          *int64                     `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                     `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                     `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{}     `json:"meta_data,omitempty"`
}
type CopyRequestParams struct {
	FromSite       string `json:"from_site"`
	IdAtFromSite   string `json:"id_at_from_site"`
	Id             string `json:"id,omitempty"`
	ForSiteMerging *bool  `json:"for_site_merging,omitempty"`
}
