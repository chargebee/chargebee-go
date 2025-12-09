package chargebee

type DiscountType string

const (
	DiscountTypeFixedAmount   DiscountType = "fixed_amount"
	DiscountTypePercentage    DiscountType = "percentage"
	DiscountTypeOfferQuantity DiscountType = "offer_quantity"
)

type DurationType string

const (
	DurationTypeOneTime       DurationType = "one_time"
	DurationTypeForever       DurationType = "forever"
	DurationTypeLimitedPeriod DurationType = "limited_period"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusExpired  Status = "expired"
	StatusArchived Status = "archived"
	StatusDeleted  Status = "deleted"
	StatusFuture   Status = "future"
)

type ApplyDiscountOn string

const (
	ApplyDiscountOnPlans             ApplyDiscountOn = "plans"
	ApplyDiscountOnPlansAndAddons    ApplyDiscountOn = "plans_and_addons"
	ApplyDiscountOnPlansWithQuantity ApplyDiscountOn = "plans_with_quantity"
	ApplyDiscountOnNotApplicable     ApplyDiscountOn = "not_applicable"
)

type ApplyOn string

const (
	ApplyOnInvoiceAmount            ApplyOn = "invoice_amount"
	ApplyOnEachSpecifiedItem        ApplyOn = "each_specified_item"
	ApplyOnSpecifiedItemsTotal      ApplyOn = "specified_items_total"
	ApplyOnEachUnitOfSpecifiedItems ApplyOn = "each_unit_of_specified_items"
)

type AddonConstraint string

const (
	AddonConstraintNone          AddonConstraint = "none"
	AddonConstraintAll           AddonConstraint = "all"
	AddonConstraintSpecific      AddonConstraint = "specific"
	AddonConstraintNotApplicable AddonConstraint = "not_applicable"
)

type PlanConstraint string

const (
	PlanConstraintNone          PlanConstraint = "none"
	PlanConstraintAll           PlanConstraint = "all"
	PlanConstraintSpecific      PlanConstraint = "specific"
	PlanConstraintNotApplicable PlanConstraint = "not_applicable"
)

type ItemConstraintItemType string

const (
	ItemConstraintItemTypePlan   ItemConstraintItemType = "plan"
	ItemConstraintItemTypeAddon  ItemConstraintItemType = "addon"
	ItemConstraintItemTypeCharge ItemConstraintItemType = "charge"
)

type ItemConstraintConstraint string

const (
	ItemConstraintConstraintNone     ItemConstraintConstraint = "none"
	ItemConstraintConstraintAll      ItemConstraintConstraint = "all"
	ItemConstraintConstraintSpecific ItemConstraintConstraint = "specific"
	ItemConstraintConstraintCriteria ItemConstraintConstraint = "criteria"
)

type ItemConstraintCriteriaItemType string

const (
	ItemConstraintCriteriaItemTypePlan   ItemConstraintCriteriaItemType = "plan"
	ItemConstraintCriteriaItemTypeAddon  ItemConstraintCriteriaItemType = "addon"
	ItemConstraintCriteriaItemTypeCharge ItemConstraintCriteriaItemType = "charge"
)

type CouponConstraintEntityType string

const (
	CouponConstraintEntityTypeCustomer CouponConstraintEntityType = "customer"
)

type CouponConstraintType string

const (
	CouponConstraintTypeMaxRedemptions   CouponConstraintType = "max_redemptions"
	CouponConstraintTypeUniqueBy         CouponConstraintType = "unique_by"
	CouponConstraintTypeExistingCustomer CouponConstraintType = "existing_customer"
	CouponConstraintTypeNewCustomer      CouponConstraintType = "new_customer"
)

type Coupon struct {
	Id                     string                    `json:"id"`
	Name                   string                    `json:"name"`
	InvoiceName            string                    `json:"invoice_name"`
	DiscountType           DiscountType              `json:"discount_type"`
	DiscountPercentage     float64                   `json:"discount_percentage"`
	DiscountAmount         int64                     `json:"discount_amount"`
	DiscountQuantity       int32                     `json:"discount_quantity"`
	CurrencyCode           string                    `json:"currency_code"`
	DurationType           DurationType              `json:"duration_type"`
	DurationMonth          int32                     `json:"duration_month"`
	ValidFrom              int64                     `json:"valid_from"`
	ValidTill              int64                     `json:"valid_till"`
	MaxRedemptions         int32                     `json:"max_redemptions"`
	Status                 Status                    `json:"status"`
	ApplyDiscountOn        ApplyDiscountOn           `json:"apply_discount_on"`
	ApplyOn                ApplyOn                   `json:"apply_on"`
	PlanConstraint         PlanConstraint            `json:"plan_constraint"`
	AddonConstraint        AddonConstraint           `json:"addon_constraint"`
	CreatedAt              int64                     `json:"created_at"`
	ArchivedAt             int64                     `json:"archived_at"`
	ResourceVersion        int64                     `json:"resource_version"`
	UpdatedAt              int64                     `json:"updated_at"`
	IncludedInMrr          bool                      `json:"included_in_mrr"`
	Period                 int32                     `json:"period"`
	PeriodUnit             enum.PeriodUnit           `json:"period_unit"`
	PlanIds                []string                  `json:"plan_ids"`
	AddonIds               []string                  `json:"addon_ids"`
	ItemConstraints        []*ItemConstraint         `json:"item_constraints"`
	ItemConstraintCriteria []*ItemConstraintCriteria `json:"item_constraint_criteria"`
	Redemptions            int32                     `json:"redemptions"`
	InvoiceNotes           string                    `json:"invoice_notes"`
	MetaData               json.RawMessage           `json:"meta_data"`
	CouponConstraints      []*CouponConstraint       `json:"coupon_constraints"`
	Deleted                bool                      `json:"deleted"`
	CustomField            map[string]interface{}    `json:"custom_field"`
	Object                 string                    `json:"object"`
}
type ItemConstraint struct {
	ItemType     ItemConstraintItemType   `json:"item_type"`
	Constraint   ItemConstraintConstraint `json:"constraint"`
	ItemPriceIds json.RawMessage          `json:"item_price_ids"`
	Object       string                   `json:"object"`
}
type ItemConstraintCriteria struct {
	ItemType         ItemConstraintCriteriaItemType `json:"item_type"`
	Currencies       json.RawMessage                `json:"currencies"`
	ItemFamilyIds    json.RawMessage                `json:"item_family_ids"`
	ItemPricePeriods json.RawMessage                `json:"item_price_periods"`
	Object           string                         `json:"object"`
}
type CouponConstraint struct {
	EntityType CouponConstraintEntityType `json:"entity_type"`
	Type       CouponConstraintType       `json:"type"`
	Value      string                     `json:"value"`
	Object     string                     `json:"object"`
}
type CreateRequest struct {
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	InvoiceName        string                 `json:"invoice_name,omitempty"`
	DiscountType       DiscountType           `json:"discount_type,omitempty"`
	DiscountAmount     *int64                 `json:"discount_amount,omitempty"`
	CurrencyCode       string                 `json:"currency_code,omitempty"`
	DiscountPercentage *float64               `json:"discount_percentage,omitempty"`
	DiscountQuantity   *int32                 `json:"discount_quantity,omitempty"`
	ApplyOn            ApplyOn                `json:"apply_on"`
	DurationType       DurationType           `json:"duration_type,omitempty"`
	DurationMonth      *int32                 `json:"duration_month,omitempty"`
	ValidTill          *int64                 `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                 `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                 `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{} `json:"meta_data,omitempty"`
	IncludedInMrr      *bool                  `json:"included_in_mrr,omitempty"`
	Period             *int32                 `json:"period,omitempty"`
	PeriodUnit         enum.PeriodUnit        `json:"period_unit,omitempty"`
	PlanConstraint     PlanConstraint         `json:"plan_constraint,omitempty"`
	AddonConstraint    AddonConstraint        `json:"addon_constraint,omitempty"`
	PlanIds            []string               `json:"plan_ids,omitempty"`
	AddonIds           []string               `json:"addon_ids,omitempty"`
	Status             Status                 `json:"status,omitempty"`
}
type CreateForItemsRequest struct {
	Id                     string                                  `json:"id"`
	Name                   string                                  `json:"name"`
	InvoiceName            string                                  `json:"invoice_name,omitempty"`
	DiscountType           DiscountType                            `json:"discount_type,omitempty"`
	DiscountAmount         *int64                                  `json:"discount_amount,omitempty"`
	CurrencyCode           string                                  `json:"currency_code,omitempty"`
	DiscountPercentage     *float64                                `json:"discount_percentage,omitempty"`
	DiscountQuantity       *int32                                  `json:"discount_quantity,omitempty"`
	ApplyOn                ApplyOn                                 `json:"apply_on"`
	DurationType           DurationType                            `json:"duration_type,omitempty"`
	DurationMonth          *int32                                  `json:"duration_month,omitempty"`
	ValidFrom              *int64                                  `json:"valid_from,omitempty"`
	ValidTill              *int64                                  `json:"valid_till,omitempty"`
	MaxRedemptions         *int32                                  `json:"max_redemptions,omitempty"`
	InvoiceNotes           string                                  `json:"invoice_notes,omitempty"`
	MetaData               map[string]interface{}                  `json:"meta_data,omitempty"`
	IncludedInMrr          *bool                                   `json:"included_in_mrr,omitempty"`
	Period                 *int32                                  `json:"period,omitempty"`
	PeriodUnit             enum.PeriodUnit                         `json:"period_unit,omitempty"`
	ItemConstraints        []*CreateForItemsItemConstraint         `json:"item_constraints,omitempty"`
	ItemConstraintCriteria []*CreateForItemsItemConstraintCriteria `json:"item_constraint_criteria,omitempty"`
	Status                 Status                                  `json:"status,omitempty"`
	CouponConstraints      []*CreateForItemsCouponConstraint       `json:"coupon_constraints,omitempty"`
}
type CreateForItemsItemConstraint struct {
	Constraint   coupon.ItemConstraintConstraint `json:"constraint"`
	ItemType     coupon.ItemConstraintItemType   `json:"item_type"`
	ItemPriceIds []interface{}                   `json:"item_price_ids,omitempty"`
}
type CreateForItemsItemConstraintCriteria struct {
	ItemType         coupon.ItemConstraintCriteriaItemType `json:"item_type,omitempty"`
	ItemFamilyIds    []interface{}                         `json:"item_family_ids,omitempty"`
	Currencies       []interface{}                         `json:"currencies,omitempty"`
	ItemPricePeriods []interface{}                         `json:"item_price_periods,omitempty"`
}
type CreateForItemsCouponConstraint struct {
	EntityType coupon.CouponConstraintEntityType `json:"entity_type"`
	Type       coupon.CouponConstraintType       `json:"type"`
	Value      string                            `json:"value,omitempty"`
}
type UpdateForItemsRequest struct {
	Name                   string                                  `json:"name,omitempty"`
	InvoiceName            string                                  `json:"invoice_name,omitempty"`
	DiscountType           DiscountType                            `json:"discount_type,omitempty"`
	DiscountAmount         *int64                                  `json:"discount_amount,omitempty"`
	CurrencyCode           string                                  `json:"currency_code,omitempty"`
	DiscountPercentage     *float64                                `json:"discount_percentage,omitempty"`
	DiscountQuantity       *int32                                  `json:"discount_quantity,omitempty"`
	ApplyOn                ApplyOn                                 `json:"apply_on,omitempty"`
	DurationType           DurationType                            `json:"duration_type,omitempty"`
	DurationMonth          *int32                                  `json:"duration_month,omitempty"`
	ValidFrom              *int64                                  `json:"valid_from,omitempty"`
	ValidTill              *int64                                  `json:"valid_till,omitempty"`
	MaxRedemptions         *int32                                  `json:"max_redemptions,omitempty"`
	InvoiceNotes           string                                  `json:"invoice_notes,omitempty"`
	MetaData               map[string]interface{}                  `json:"meta_data,omitempty"`
	IncludedInMrr          *bool                                   `json:"included_in_mrr,omitempty"`
	Period                 *int32                                  `json:"period,omitempty"`
	PeriodUnit             enum.PeriodUnit                         `json:"period_unit,omitempty"`
	ItemConstraints        []*UpdateForItemsItemConstraint         `json:"item_constraints,omitempty"`
	ItemConstraintCriteria []*UpdateForItemsItemConstraintCriteria `json:"item_constraint_criteria,omitempty"`
	CouponConstraints      []*UpdateForItemsCouponConstraint       `json:"coupon_constraints,omitempty"`
}
type UpdateForItemsItemConstraint struct {
	Constraint   coupon.ItemConstraintConstraint `json:"constraint"`
	ItemType     coupon.ItemConstraintItemType   `json:"item_type"`
	ItemPriceIds []interface{}                   `json:"item_price_ids,omitempty"`
}
type UpdateForItemsItemConstraintCriteria struct {
	ItemType         coupon.ItemConstraintCriteriaItemType `json:"item_type,omitempty"`
	ItemFamilyIds    []interface{}                         `json:"item_family_ids,omitempty"`
	Currencies       []interface{}                         `json:"currencies,omitempty"`
	ItemPricePeriods []interface{}                         `json:"item_price_periods,omitempty"`
}
type UpdateForItemsCouponConstraint struct {
	EntityType coupon.CouponConstraintEntityType `json:"entity_type"`
	Type       coupon.CouponConstraintType       `json:"type"`
	Value      string                            `json:"value,omitempty"`
}
type ListRequest struct {
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
type UpdateRequest struct {
	Name               string                 `json:"name,omitempty"`
	InvoiceName        string                 `json:"invoice_name,omitempty"`
	DiscountType       DiscountType           `json:"discount_type,omitempty"`
	DiscountAmount     *int64                 `json:"discount_amount,omitempty"`
	CurrencyCode       string                 `json:"currency_code,omitempty"`
	DiscountPercentage *float64               `json:"discount_percentage,omitempty"`
	DiscountQuantity   *int32                 `json:"discount_quantity,omitempty"`
	ApplyOn            ApplyOn                `json:"apply_on,omitempty"`
	DurationType       DurationType           `json:"duration_type,omitempty"`
	DurationMonth      *int32                 `json:"duration_month,omitempty"`
	ValidTill          *int64                 `json:"valid_till,omitempty"`
	MaxRedemptions     *int32                 `json:"max_redemptions,omitempty"`
	InvoiceNotes       string                 `json:"invoice_notes,omitempty"`
	MetaData           map[string]interface{} `json:"meta_data,omitempty"`
	IncludedInMrr      *bool                  `json:"included_in_mrr,omitempty"`
	Period             *int32                 `json:"period,omitempty"`
	PeriodUnit         enum.PeriodUnit        `json:"period_unit,omitempty"`
	PlanConstraint     PlanConstraint         `json:"plan_constraint,omitempty"`
	AddonConstraint    AddonConstraint        `json:"addon_constraint,omitempty"`
	PlanIds            []string               `json:"plan_ids,omitempty"`
	AddonIds           []string               `json:"addon_ids,omitempty"`
}
type CopyRequest struct {
	FromSite       string `json:"from_site"`
	IdAtFromSite   string `json:"id_at_from_site"`
	Id             string `json:"id,omitempty"`
	ForSiteMerging *bool  `json:"for_site_merging,omitempty"`
}

type CreateResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type CreateForItemsResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type UpdateForItemsResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type ListCouponResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type ListResponse struct {
	List       []*ListCouponResponse `json:"list,omitempty"`
	NextOffset string                `json:"next_offset,omitempty"`
}

type RetrieveResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type UpdateResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type DeleteResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type CopyResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

type UnarchiveResponse struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}
