package coupon

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	couponEnum "github.com/chargebee/chargebee-go/models/coupon/enum"
)

type Coupon struct {
	Id                     string                     `json:"id"`
	Name                   string                     `json:"name"`
	InvoiceName            string                     `json:"invoice_name"`
	DiscountType           couponEnum.DiscountType    `json:"discount_type"`
	DiscountPercentage     float64                    `json:"discount_percentage"`
	DiscountAmount         int32                      `json:"discount_amount"`
	DiscountQuantity       int32                      `json:"discount_quantity"`
	CurrencyCode           string                     `json:"currency_code"`
	DurationType           couponEnum.DurationType    `json:"duration_type"`
	DurationMonth          int32                      `json:"duration_month"`
	ValidTill              int64                      `json:"valid_till"`
	MaxRedemptions         int32                      `json:"max_redemptions"`
	Status                 couponEnum.Status          `json:"status"`
	ApplyDiscountOn        couponEnum.ApplyDiscountOn `json:"apply_discount_on"`
	ApplyOn                couponEnum.ApplyOn         `json:"apply_on"`
	PlanConstraint         couponEnum.PlanConstraint  `json:"plan_constraint"`
	AddonConstraint        couponEnum.AddonConstraint `json:"addon_constraint"`
	CreatedAt              int64                      `json:"created_at"`
	ArchivedAt             int64                      `json:"archived_at"`
	ResourceVersion        int64                      `json:"resource_version"`
	UpdatedAt              int64                      `json:"updated_at"`
	IncludedInMrr          bool                       `json:"included_in_mrr"`
	Period                 int32                      `json:"period"`
	PeriodUnit             enum.PeriodUnit            `json:"period_unit"`
	PlanIds                []string                   `json:"plan_ids"`
	AddonIds               []string                   `json:"addon_ids"`
	ItemConstraints        []*ItemConstraint          `json:"item_constraints"`
	ItemConstraintCriteria []*ItemConstraintCriteria  `json:"item_constraint_criteria"`
	Redemptions            int32                      `json:"redemptions"`
	InvoiceNotes           string                     `json:"invoice_notes"`
	MetaData               json.RawMessage            `json:"meta_data"`
	Object                 string                     `json:"object"`
}
type ItemConstraint struct {
	ItemType     couponEnum.ItemConstraintItemType   `json:"item_type"`
	Constraint   couponEnum.ItemConstraintConstraint `json:"constraint"`
	ItemPriceIds json.RawMessage                     `json:"item_price_ids"`
	Object       string                              `json:"object"`
}
type ItemConstraintCriteria struct {
	ItemType         couponEnum.ItemConstraintCriteriaItemType `json:"item_type"`
	Currencies       json.RawMessage                           `json:"currencies"`
	ItemFamilyIds    json.RawMessage                           `json:"item_family_ids"`
	ItemPricePeriods json.RawMessage                           `json:"item_price_periods"`
	Object           string                                    `json:"object"`
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
	DurationType       couponEnum.DurationType    `json:"duration_type"`
	DurationMonth      *int32                     `json:"duration_month,omitempty"`
	ValidTill          *int64                     `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                     `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                     `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{}     `json:"meta_data,omitempty"`
	IncludedInMrr      *bool                      `json:"included_in_mrr,omitempty"`
	Period             *int32                     `json:"period,omitempty"`
	PeriodUnit         enum.PeriodUnit            `json:"period_unit,omitempty"`
	PlanConstraint     couponEnum.PlanConstraint  `json:"plan_constraint,omitempty"`
	AddonConstraint    couponEnum.AddonConstraint `json:"addon_constraint,omitempty"`
	PlanIds            []string                   `json:"plan_ids,omitempty"`
	AddonIds           []string                   `json:"addon_ids,omitempty"`
	Status             couponEnum.Status          `json:"status,omitempty"`
}
type CreateForItemsRequestParams struct {
	Id                     string                                        `json:"id"`
	Name                   string                                        `json:"name"`
	InvoiceName            string                                        `json:"invoice_name,omitempty"`
	DiscountType           couponEnum.DiscountType                       `json:"discount_type"`
	DiscountAmount         *int32                                        `json:"discount_amount,omitempty"`
	CurrencyCode           string                                        `json:"currency_code,omitempty"`
	DiscountPercentage     *float64                                      `json:"discount_percentage,omitempty"`
	DiscountQuantity       *int32                                        `json:"discount_quantity,omitempty"`
	ApplyOn                couponEnum.ApplyOn                            `json:"apply_on"`
	DurationType           couponEnum.DurationType                       `json:"duration_type"`
	DurationMonth          *int32                                        `json:"duration_month,omitempty"`
	ValidTill              *int64                                        `json:"valid_till,omitempty"`
	MaxRedemptions         *int32                                        `json:"max_redemptions,omitempty"`
	InvoiceNotes           string                                        `json:"invoice_notes,omitempty"`
	MetaData               map[string]interface{}                        `json:"meta_data,omitempty"`
	IncludedInMrr          *bool                                         `json:"included_in_mrr,omitempty"`
	Period                 *int32                                        `json:"period,omitempty"`
	PeriodUnit             enum.PeriodUnit                               `json:"period_unit,omitempty"`
	ItemConstraints        []*CreateForItemsItemConstraintParams         `json:"item_constraints,omitempty"`
	ItemConstraintCriteria []*CreateForItemsItemConstraintCriteriaParams `json:"item_constraint_criteria,omitempty"`
	Status                 couponEnum.Status                             `json:"status,omitempty"`
}
type CreateForItemsItemConstraintParams struct {
	Constraint   couponEnum.ItemConstraintConstraint `json:"constraint"`
	ItemType     couponEnum.ItemConstraintItemType   `json:"item_type"`
	ItemPriceIds []map[string]interface{}            `json:"item_price_ids,omitempty"`
}
type CreateForItemsItemConstraintCriteriaParams struct {
	ItemType         couponEnum.ItemConstraintCriteriaItemType `json:"item_type,omitempty"`
	ItemFamilyIds    []map[string]interface{}                  `json:"item_family_ids,omitempty"`
	Currencies       []map[string]interface{}                  `json:"currencies,omitempty"`
	ItemPricePeriods []map[string]interface{}                  `json:"item_price_periods,omitempty"`
}
type UpdateForItemsRequestParams struct {
	Name                   string                                        `json:"name,omitempty"`
	InvoiceName            string                                        `json:"invoice_name,omitempty"`
	DiscountType           couponEnum.DiscountType                       `json:"discount_type,omitempty"`
	DiscountAmount         *int32                                        `json:"discount_amount,omitempty"`
	CurrencyCode           string                                        `json:"currency_code,omitempty"`
	DiscountPercentage     *float64                                      `json:"discount_percentage,omitempty"`
	ApplyOn                couponEnum.ApplyOn                            `json:"apply_on,omitempty"`
	DurationType           couponEnum.DurationType                       `json:"duration_type,omitempty"`
	DurationMonth          *int32                                        `json:"duration_month,omitempty"`
	ValidTill              *int64                                        `json:"valid_till,omitempty"`
	MaxRedemptions         *int32                                        `json:"max_redemptions,omitempty"`
	InvoiceNotes           string                                        `json:"invoice_notes,omitempty"`
	MetaData               map[string]interface{}                        `json:"meta_data,omitempty"`
	IncludedInMrr          *bool                                         `json:"included_in_mrr,omitempty"`
	Period                 *int32                                        `json:"period,omitempty"`
	PeriodUnit             enum.PeriodUnit                               `json:"period_unit,omitempty"`
	ItemConstraints        []*UpdateForItemsItemConstraintParams         `json:"item_constraints,omitempty"`
	ItemConstraintCriteria []*UpdateForItemsItemConstraintCriteriaParams `json:"item_constraint_criteria,omitempty"`
}
type UpdateForItemsItemConstraintParams struct {
	Constraint   couponEnum.ItemConstraintConstraint `json:"constraint"`
	ItemType     couponEnum.ItemConstraintItemType   `json:"item_type"`
	ItemPriceIds []map[string]interface{}            `json:"item_price_ids,omitempty"`
}
type UpdateForItemsItemConstraintCriteriaParams struct {
	ItemType         couponEnum.ItemConstraintCriteriaItemType `json:"item_type,omitempty"`
	ItemFamilyIds    []map[string]interface{}                  `json:"item_family_ids,omitempty"`
	Currencies       []map[string]interface{}                  `json:"currencies,omitempty"`
	ItemPricePeriods []map[string]interface{}                  `json:"item_price_periods,omitempty"`
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
	CurrencyCode *filter.StringFilter    `json:"currency_code,omitempty"`
}
type UpdateRequestParams struct {
	Name               string                     `json:"name,omitempty"`
	InvoiceName        string                     `json:"invoice_name,omitempty"`
	DiscountType       couponEnum.DiscountType    `json:"discount_type,omitempty"`
	DiscountAmount     *int32                     `json:"discount_amount,omitempty"`
	CurrencyCode       string                     `json:"currency_code,omitempty"`
	DiscountPercentage *float64                   `json:"discount_percentage,omitempty"`
	ApplyOn            couponEnum.ApplyOn         `json:"apply_on,omitempty"`
	DurationType       couponEnum.DurationType    `json:"duration_type,omitempty"`
	DurationMonth      *int32                     `json:"duration_month,omitempty"`
	ValidTill          *int64                     `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                     `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                     `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{}     `json:"meta_data,omitempty"`
	IncludedInMrr      *bool                      `json:"included_in_mrr,omitempty"`
	Period             *int32                     `json:"period,omitempty"`
	PeriodUnit         enum.PeriodUnit            `json:"period_unit,omitempty"`
	PlanConstraint     couponEnum.PlanConstraint  `json:"plan_constraint,omitempty"`
	AddonConstraint    couponEnum.AddonConstraint `json:"addon_constraint,omitempty"`
	PlanIds            []string                   `json:"plan_ids,omitempty"`
	AddonIds           []string                   `json:"addon_ids,omitempty"`
}
type CopyRequestParams struct {
	FromSite       string `json:"from_site"`
	IdAtFromSite   string `json:"id_at_from_site"`
	Id             string `json:"id,omitempty"`
	ForSiteMerging *bool  `json:"for_site_merging,omitempty"`
}
