package chargebee

import (
	"encoding/json"
)

type CouponDiscountType string

const (
	CouponDiscountTypeFixedAmount   CouponDiscountType = "fixed_amount"
	CouponDiscountTypePercentage    CouponDiscountType = "percentage"
	CouponDiscountTypeOfferQuantity CouponDiscountType = "offer_quantity"
)

type CouponDurationType string

const (
	CouponDurationTypeOneTime       CouponDurationType = "one_time"
	CouponDurationTypeForever       CouponDurationType = "forever"
	CouponDurationTypeLimitedPeriod CouponDurationType = "limited_period"
)

type CouponStatus string

const (
	CouponStatusActive   CouponStatus = "active"
	CouponStatusExpired  CouponStatus = "expired"
	CouponStatusArchived CouponStatus = "archived"
	CouponStatusDeleted  CouponStatus = "deleted"
	CouponStatusFuture   CouponStatus = "future"
)

type CouponApplyDiscountOn string

const (
	CouponApplyDiscountOnPlans             CouponApplyDiscountOn = "plans"
	CouponApplyDiscountOnPlansAndAddons    CouponApplyDiscountOn = "plans_and_addons"
	CouponApplyDiscountOnPlansWithQuantity CouponApplyDiscountOn = "plans_with_quantity"
	CouponApplyDiscountOnNotApplicable     CouponApplyDiscountOn = "not_applicable"
)

type CouponApplyOn string

const (
	CouponApplyOnInvoiceAmount            CouponApplyOn = "invoice_amount"
	CouponApplyOnEachSpecifiedItem        CouponApplyOn = "each_specified_item"
	CouponApplyOnSpecifiedItemsTotal      CouponApplyOn = "specified_items_total"
	CouponApplyOnEachUnitOfSpecifiedItems CouponApplyOn = "each_unit_of_specified_items"
)

type CouponPeriodUnit string

const (
	CouponPeriodUnitDay   CouponPeriodUnit = "day"
	CouponPeriodUnitWeek  CouponPeriodUnit = "week"
	CouponPeriodUnitMonth CouponPeriodUnit = "month"
	CouponPeriodUnitYear  CouponPeriodUnit = "year"
)

type CouponAddonConstraint string

const (
	CouponAddonConstraintNone          CouponAddonConstraint = "none"
	CouponAddonConstraintAll           CouponAddonConstraint = "all"
	CouponAddonConstraintSpecific      CouponAddonConstraint = "specific"
	CouponAddonConstraintNotApplicable CouponAddonConstraint = "not_applicable"
)

type CouponPlanConstraint string

const (
	CouponPlanConstraintNone          CouponPlanConstraint = "none"
	CouponPlanConstraintAll           CouponPlanConstraint = "all"
	CouponPlanConstraintSpecific      CouponPlanConstraint = "specific"
	CouponPlanConstraintNotApplicable CouponPlanConstraint = "not_applicable"
)

type CouponItemConstraintItemType string

const (
	CouponItemConstraintItemTypePlan   CouponItemConstraintItemType = "plan"
	CouponItemConstraintItemTypeAddon  CouponItemConstraintItemType = "addon"
	CouponItemConstraintItemTypeCharge CouponItemConstraintItemType = "charge"
)

type CouponItemConstraintConstraint string

const (
	CouponItemConstraintConstraintNone     CouponItemConstraintConstraint = "none"
	CouponItemConstraintConstraintAll      CouponItemConstraintConstraint = "all"
	CouponItemConstraintConstraintSpecific CouponItemConstraintConstraint = "specific"
	CouponItemConstraintConstraintCriteria CouponItemConstraintConstraint = "criteria"
)

type CouponItemConstraintCriteriaItemType string

const (
	CouponItemConstraintCriteriaItemTypePlan   CouponItemConstraintCriteriaItemType = "plan"
	CouponItemConstraintCriteriaItemTypeAddon  CouponItemConstraintCriteriaItemType = "addon"
	CouponItemConstraintCriteriaItemTypeCharge CouponItemConstraintCriteriaItemType = "charge"
)

type CouponCouponConstraintEntityType string

const (
	CouponCouponConstraintEntityTypeCustomer CouponCouponConstraintEntityType = "customer"
)

type CouponCouponConstraintType string

const (
	CouponCouponConstraintTypeMaxRedemptions   CouponCouponConstraintType = "max_redemptions"
	CouponCouponConstraintTypeUniqueBy         CouponCouponConstraintType = "unique_by"
	CouponCouponConstraintTypeExistingCustomer CouponCouponConstraintType = "existing_customer"
	CouponCouponConstraintTypeNewCustomer      CouponCouponConstraintType = "new_customer"
)

// just struct
type Coupon struct {
	Id                     string                          `json:"id"`
	Name                   string                          `json:"name"`
	InvoiceName            string                          `json:"invoice_name"`
	DiscountType           CouponDiscountType              `json:"discount_type"`
	DiscountPercentage     float64                         `json:"discount_percentage"`
	DiscountAmount         int64                           `json:"discount_amount"`
	DiscountQuantity       int32                           `json:"discount_quantity"`
	CurrencyCode           string                          `json:"currency_code"`
	DurationType           CouponDurationType              `json:"duration_type"`
	ValidFrom              int64                           `json:"valid_from"`
	ValidTill              int64                           `json:"valid_till"`
	MaxRedemptions         int32                           `json:"max_redemptions"`
	Status                 CouponStatus                    `json:"status"`
	ApplyOn                CouponApplyOn                   `json:"apply_on"`
	PlanConstraint         CouponPlanConstraint            `json:"plan_constraint"`
	AddonConstraint        CouponAddonConstraint           `json:"addon_constraint"`
	CreatedAt              int64                           `json:"created_at"`
	ArchivedAt             int64                           `json:"archived_at"`
	ResourceVersion        int64                           `json:"resource_version"`
	UpdatedAt              int64                           `json:"updated_at"`
	IncludedInMrr          bool                            `json:"included_in_mrr"`
	Period                 int32                           `json:"period"`
	PeriodUnit             CouponPeriodUnit                `json:"period_unit"`
	PlanIds                []string                        `json:"plan_ids"`
	AddonIds               []string                        `json:"addon_ids"`
	ItemConstraints        []*CouponItemConstraint         `json:"item_constraints"`
	ItemConstraintCriteria []*CouponItemConstraintCriteria `json:"item_constraint_criteria"`
	Redemptions            int32                           `json:"redemptions"`
	InvoiceNotes           string                          `json:"invoice_notes"`
	MetaData               json.RawMessage                 `json:"meta_data"`
	CouponConstraints      []*CouponCouponConstraint       `json:"coupon_constraints"`
	Deleted                bool                            `json:"deleted"`
	CustomField            CustomField                     `json:"custom_field"`
	Object                 string                          `json:"object"`
}

// sub resources
type CouponItemConstraint struct {
	ItemType     CouponItemConstraintItemType   `json:"item_type"`
	Constraint   CouponItemConstraintConstraint `json:"constraint"`
	ItemPriceIds json.RawMessage                `json:"item_price_ids"`
	Object       string                         `json:"object"`
}
type CouponItemConstraintCriteria struct {
	ItemType         CouponItemConstraintCriteriaItemType `json:"item_type"`
	Currencies       json.RawMessage                      `json:"currencies"`
	ItemFamilyIds    json.RawMessage                      `json:"item_family_ids"`
	ItemPricePeriods json.RawMessage                      `json:"item_price_periods"`
	Object           string                               `json:"object"`
}
type CouponCouponConstraint struct {
	EntityType CouponCouponConstraintEntityType `json:"entity_type"`
	Type       CouponCouponConstraintType       `json:"type"`
	Value      string                           `json:"value"`
	Object     string                           `json:"object"`
}

// operations
// input params
type CouponCreateRequest struct {
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	InvoiceName        string                 `json:"invoice_name,omitempty"`
	DiscountType       CouponDiscountType     `json:"discount_type,omitempty"`
	DiscountAmount     *int64                 `json:"discount_amount,omitempty"`
	CurrencyCode       string                 `json:"currency_code,omitempty"`
	DiscountPercentage *float64               `json:"discount_percentage,omitempty"`
	DiscountQuantity   *int32                 `json:"discount_quantity,omitempty"`
	ApplyOn            CouponApplyOn          `json:"apply_on"`
	DurationType       CouponDurationType     `json:"duration_type,omitempty"`
	DurationMonth      *int32                 `json:"duration_month,omitempty"`
	ValidTill          *int64                 `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                 `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                 `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{} `json:"meta_data,omitempty"`
	IncludedInMrr      *bool                  `json:"included_in_mrr,omitempty"`
	Period             *int32                 `json:"period,omitempty"`
	PeriodUnit         CouponPeriodUnit       `json:"period_unit,omitempty"`
	PlanConstraint     CouponPlanConstraint   `json:"plan_constraint,omitempty"`
	AddonConstraint    CouponAddonConstraint  `json:"addon_constraint,omitempty"`
	PlanIds            []string               `json:"plan_ids,omitempty"`
	AddonIds           []string               `json:"addon_ids,omitempty"`
	Status             CouponStatus           `json:"status,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *CouponCreateRequest) payload() any { return r }

type CouponCreateForItemsRequest struct {
	Id                     string                                        `json:"id"`
	Name                   string                                        `json:"name"`
	InvoiceName            string                                        `json:"invoice_name,omitempty"`
	DiscountType           CouponDiscountType                            `json:"discount_type,omitempty"`
	DiscountAmount         *int64                                        `json:"discount_amount,omitempty"`
	CurrencyCode           string                                        `json:"currency_code,omitempty"`
	DiscountPercentage     *float64                                      `json:"discount_percentage,omitempty"`
	DiscountQuantity       *int32                                        `json:"discount_quantity,omitempty"`
	ApplyOn                CouponApplyOn                                 `json:"apply_on"`
	DurationType           CouponDurationType                            `json:"duration_type,omitempty"`
	DurationMonth          *int32                                        `json:"duration_month,omitempty"`
	ValidFrom              *int64                                        `json:"valid_from,omitempty"`
	ValidTill              *int64                                        `json:"valid_till,omitempty"`
	MaxRedemptions         *int32                                        `json:"max_redemptions,omitempty"`
	InvoiceNotes           string                                        `json:"invoice_notes,omitempty"`
	MetaData               map[string]interface{}                        `json:"meta_data,omitempty"`
	IncludedInMrr          *bool                                         `json:"included_in_mrr,omitempty"`
	Period                 *int32                                        `json:"period,omitempty"`
	PeriodUnit             CouponPeriodUnit                              `json:"period_unit,omitempty"`
	ItemConstraints        []*CouponCreateForItemsItemConstraint         `json:"item_constraints,omitempty"`
	ItemConstraintCriteria []*CouponCreateForItemsItemConstraintCriteria `json:"item_constraint_criteria,omitempty"`
	Status                 CouponStatus                                  `json:"status,omitempty"`
	CouponConstraints      []*CouponCreateForItemsCouponConstraint       `json:"coupon_constraints,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *CouponCreateForItemsRequest) payload() any { return r }

// input sub resource params multi
type CouponCreateForItemsItemConstraint struct {
	Constraint   CouponItemConstraintConstraint `json:"constraint"`
	ItemType     CouponItemConstraintItemType   `json:"item_type"`
	ItemPriceIds []interface{}                  `json:"item_price_ids,omitempty"`
}

// input sub resource params multi
type CouponCreateForItemsItemConstraintCriteria struct {
	ItemType         CouponItemConstraintCriteriaItemType `json:"item_type,omitempty"`
	ItemFamilyIds    []interface{}                        `json:"item_family_ids,omitempty"`
	Currencies       []interface{}                        `json:"currencies,omitempty"`
	ItemPricePeriods []interface{}                        `json:"item_price_periods,omitempty"`
}

// input sub resource params multi
type CouponCreateForItemsCouponConstraint struct {
	EntityType CouponCouponConstraintEntityType `json:"entity_type"`
	Type       CouponCouponConstraintType       `json:"type"`
	Value      string                           `json:"value,omitempty"`
}
type CouponUpdateForItemsRequest struct {
	Name                   string                                        `json:"name,omitempty"`
	InvoiceName            string                                        `json:"invoice_name,omitempty"`
	DiscountType           CouponDiscountType                            `json:"discount_type,omitempty"`
	DiscountAmount         *int64                                        `json:"discount_amount,omitempty"`
	CurrencyCode           string                                        `json:"currency_code,omitempty"`
	DiscountPercentage     *float64                                      `json:"discount_percentage,omitempty"`
	DiscountQuantity       *int32                                        `json:"discount_quantity,omitempty"`
	ApplyOn                CouponApplyOn                                 `json:"apply_on,omitempty"`
	DurationType           CouponDurationType                            `json:"duration_type,omitempty"`
	DurationMonth          *int32                                        `json:"duration_month,omitempty"`
	ValidFrom              *int64                                        `json:"valid_from,omitempty"`
	ValidTill              *int64                                        `json:"valid_till,omitempty"`
	MaxRedemptions         *int32                                        `json:"max_redemptions,omitempty"`
	InvoiceNotes           string                                        `json:"invoice_notes,omitempty"`
	MetaData               map[string]interface{}                        `json:"meta_data,omitempty"`
	IncludedInMrr          *bool                                         `json:"included_in_mrr,omitempty"`
	Period                 *int32                                        `json:"period,omitempty"`
	PeriodUnit             CouponPeriodUnit                              `json:"period_unit,omitempty"`
	ItemConstraints        []*CouponUpdateForItemsItemConstraint         `json:"item_constraints,omitempty"`
	ItemConstraintCriteria []*CouponUpdateForItemsItemConstraintCriteria `json:"item_constraint_criteria,omitempty"`
	CouponConstraints      []*CouponUpdateForItemsCouponConstraint       `json:"coupon_constraints,omitempty"`
	apiRequest             `json:"-" form:"-"`
}

func (r *CouponUpdateForItemsRequest) payload() any { return r }

// input sub resource params multi
type CouponUpdateForItemsItemConstraint struct {
	Constraint   CouponItemConstraintConstraint `json:"constraint"`
	ItemType     CouponItemConstraintItemType   `json:"item_type"`
	ItemPriceIds []interface{}                  `json:"item_price_ids,omitempty"`
}

// input sub resource params multi
type CouponUpdateForItemsItemConstraintCriteria struct {
	ItemType         CouponItemConstraintCriteriaItemType `json:"item_type,omitempty"`
	ItemFamilyIds    []interface{}                        `json:"item_family_ids,omitempty"`
	Currencies       []interface{}                        `json:"currencies,omitempty"`
	ItemPricePeriods []interface{}                        `json:"item_price_periods,omitempty"`
}

// input sub resource params multi
type CouponUpdateForItemsCouponConstraint struct {
	EntityType CouponCouponConstraintEntityType `json:"entity_type"`
	Type       CouponCouponConstraintType       `json:"type"`
	Value      string                           `json:"value,omitempty"`
}
type CouponListRequest struct {
	Limit        *int32           `json:"limit,omitempty"`
	Offset       string           `json:"offset,omitempty"`
	Id           *StringFilter    `json:"id,omitempty"`
	Name         *StringFilter    `json:"name,omitempty"`
	DiscountType *EnumFilter      `json:"discount_type,omitempty"`
	DurationType *EnumFilter      `json:"duration_type,omitempty"`
	Status       *EnumFilter      `json:"status,omitempty"`
	ApplyOn      *EnumFilter      `json:"apply_on,omitempty"`
	CreatedAt    *TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt    *TimestampFilter `json:"updated_at,omitempty"`
	SortBy       *SortFilter      `json:"sort_by,omitempty"`
	CurrencyCode *StringFilter    `json:"currency_code,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *CouponListRequest) payload() any { return r }

type CouponUpdateRequest struct {
	Name               string                 `json:"name,omitempty"`
	InvoiceName        string                 `json:"invoice_name,omitempty"`
	DiscountType       CouponDiscountType     `json:"discount_type,omitempty"`
	DiscountAmount     *int64                 `json:"discount_amount,omitempty"`
	CurrencyCode       string                 `json:"currency_code,omitempty"`
	DiscountPercentage *float64               `json:"discount_percentage,omitempty"`
	DiscountQuantity   *int32                 `json:"discount_quantity,omitempty"`
	ApplyOn            CouponApplyOn          `json:"apply_on,omitempty"`
	DurationType       CouponDurationType     `json:"duration_type,omitempty"`
	DurationMonth      *int32                 `json:"duration_month,omitempty"`
	ValidTill          *int64                 `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                 `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                 `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{} `json:"meta_data,omitempty"`
	IncludedInMrr      *bool                  `json:"included_in_mrr,omitempty"`
	Period             *int32                 `json:"period,omitempty"`
	PeriodUnit         CouponPeriodUnit       `json:"period_unit,omitempty"`
	PlanConstraint     CouponPlanConstraint   `json:"plan_constraint,omitempty"`
	AddonConstraint    CouponAddonConstraint  `json:"addon_constraint,omitempty"`
	PlanIds            []string               `json:"plan_ids,omitempty"`
	AddonIds           []string               `json:"addon_ids,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *CouponUpdateRequest) payload() any { return r }

type CouponCopyRequest struct {
	FromSite       string `json:"from_site"`
	IdAtFromSite   string `json:"id_at_from_site"`
	Id             string `json:"id,omitempty"`
	ForSiteMerging *bool  `json:"for_site_merging,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *CouponCopyRequest) payload() any { return r }

// operation response
type CouponCreateResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}

// operation response
type CouponCreateForItemsResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}

// operation response
type CouponUpdateForItemsResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}

// operation sub response
type CouponListCouponResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

// operation response
type CouponListResponse struct {
	List       []*CouponListCouponResponse `json:"list,omitempty"`
	NextOffset string                      `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CouponRetrieveResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}

// operation response
type CouponUpdateResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}

// operation response
type CouponDeleteResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}

// operation response
type CouponCopyResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}

// operation response
type CouponUnarchiveResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
	apiResponse
}
