package chargebee

type Status string

const (
	StatusScheduled Status = "scheduled"
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
	StatusDraft     Status = "draft"
)

type DiscountsToAddType string

const (
	DiscountsToAddTypeFixedAmount DiscountsToAddType = "fixed_amount"
	DiscountsToAddTypePercentage  DiscountsToAddType = "percentage"
)

type ContractTermActionAtTermEnd string

const (
	ContractTermActionAtTermEndRenew     ContractTermActionAtTermEnd = "renew"
	ContractTermActionAtTermEndEvergreen ContractTermActionAtTermEnd = "evergreen"
	ContractTermActionAtTermEndCancel    ContractTermActionAtTermEnd = "cancel"
	ContractTermActionAtTermEndRenewOnce ContractTermActionAtTermEnd = "renew_once"
)

type Ramp struct {
	Id                     string                  `json:"id"`
	Description            string                  `json:"description"`
	SubscriptionId         string                  `json:"subscription_id"`
	EffectiveFrom          int64                   `json:"effective_from"`
	Status                 Status                  `json:"status"`
	CreatedAt              int64                   `json:"created_at"`
	ResourceVersion        int64                   `json:"resource_version"`
	UpdatedAt              int64                   `json:"updated_at"`
	ItemsToAdd             []*ItemsToAdd           `json:"items_to_add"`
	ItemsToUpdate          []*ItemsToUpdate        `json:"items_to_update"`
	CouponsToAdd           []*CouponsToAdd         `json:"coupons_to_add"`
	DiscountsToAdd         []*DiscountsToAdd       `json:"discounts_to_add"`
	ItemTiers              []*ItemTier             `json:"item_tiers"`
	ItemsToRemove          []string                `json:"items_to_remove"`
	CouponsToRemove        []string                `json:"coupons_to_remove"`
	DiscountsToRemove      []string                `json:"discounts_to_remove"`
	ContractTerm           *ContractTerm           `json:"contract_term"`
	Deleted                bool                    `json:"deleted"`
	StatusTransitionReason *StatusTransitionReason `json:"status_transition_reason"`
	Object                 string                  `json:"object"`
}
type ItemsToAdd struct {
	ItemPriceId           string              `json:"item_price_id"`
	ItemType              enum.ItemType       `json:"item_type"`
	Quantity              int32               `json:"quantity"`
	QuantityInDecimal     string              `json:"quantity_in_decimal"`
	UnitPrice             int64               `json:"unit_price"`
	UnitPriceInDecimal    string              `json:"unit_price_in_decimal"`
	Amount                int64               `json:"amount"`
	AmountInDecimal       string              `json:"amount_in_decimal"`
	FreeQuantity          int32               `json:"free_quantity"`
	FreeQuantityInDecimal string              `json:"free_quantity_in_decimal"`
	BillingCycles         int32               `json:"billing_cycles"`
	ServicePeriodDays     int32               `json:"service_period_days"`
	MeteredQuantity       string              `json:"metered_quantity"`
	ChargeOnce            bool                `json:"charge_once"`
	ChargeOnOption        enum.ChargeOnOption `json:"charge_on_option"`
	ChargeOnEvent         enum.ChargeOnEvent  `json:"charge_on_event"`
	Object                string              `json:"object"`
}
type ItemsToUpdate struct {
	ItemPriceId           string              `json:"item_price_id"`
	ItemType              enum.ItemType       `json:"item_type"`
	Quantity              int32               `json:"quantity"`
	QuantityInDecimal     string              `json:"quantity_in_decimal"`
	UnitPrice             int64               `json:"unit_price"`
	UnitPriceInDecimal    string              `json:"unit_price_in_decimal"`
	Amount                int64               `json:"amount"`
	AmountInDecimal       string              `json:"amount_in_decimal"`
	FreeQuantity          int32               `json:"free_quantity"`
	FreeQuantityInDecimal string              `json:"free_quantity_in_decimal"`
	BillingCycles         int32               `json:"billing_cycles"`
	ServicePeriodDays     int32               `json:"service_period_days"`
	MeteredQuantity       string              `json:"metered_quantity"`
	ChargeOnce            bool                `json:"charge_once"`
	ChargeOnOption        enum.ChargeOnOption `json:"charge_on_option"`
	ChargeOnEvent         enum.ChargeOnEvent  `json:"charge_on_event"`
	Object                string              `json:"object"`
}
type CouponsToAdd struct {
	CouponId  string `json:"coupon_id"`
	ApplyTill int64  `json:"apply_till"`
	Object    string `json:"object"`
}
type DiscountsToAdd struct {
	Id            string             `json:"id"`
	InvoiceName   string             `json:"invoice_name"`
	Type          DiscountsToAddType `json:"type"`
	Percentage    float64            `json:"percentage"`
	Amount        int64              `json:"amount"`
	DurationType  enum.DurationType  `json:"duration_type"`
	Period        int32              `json:"period"`
	PeriodUnit    enum.PeriodUnit    `json:"period_unit"`
	IncludedInMrr bool               `json:"included_in_mrr"`
	ApplyOn       enum.ApplyOn       `json:"apply_on"`
	ItemPriceId   string             `json:"item_price_id"`
	CreatedAt     int64              `json:"created_at"`
	Object        string             `json:"object"`
}
type ItemTier struct {
	ItemPriceId           string           `json:"item_price_id"`
	StartingUnit          int32            `json:"starting_unit"`
	EndingUnit            int32            `json:"ending_unit"`
	Price                 int64            `json:"price"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal"`
	PriceInDecimal        string           `json:"price_in_decimal"`
	PricingType           enum.PricingType `json:"pricing_type"`
	PackageSize           int32            `json:"package_size"`
	Index                 int32            `json:"index"`
	Object                string           `json:"object"`
}
type ContractTerm struct {
	CancellationCutoffPeriod int32                       `json:"cancellation_cutoff_period"`
	RenewalBillingCycles     int32                       `json:"renewal_billing_cycles"`
	ActionAtTermEnd          ContractTermActionAtTermEnd `json:"action_at_term_end"`
	Object                   string                      `json:"object"`
}
type StatusTransitionReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Object  string `json:"object"`
}
type CreateForSubscriptionRequest struct {
	EffectiveFrom     *int64                                 `json:"effective_from"`
	Description       string                                 `json:"description,omitempty"`
	CouponsToRemove   []string                               `json:"coupons_to_remove,omitempty"`
	DiscountsToRemove []string                               `json:"discounts_to_remove,omitempty"`
	ItemsToRemove     []string                               `json:"items_to_remove,omitempty"`
	ItemsToAdd        []*CreateForSubscriptionItemsToAdd     `json:"items_to_add,omitempty"`
	ItemsToUpdate     []*CreateForSubscriptionItemsToUpdate  `json:"items_to_update,omitempty"`
	ItemTiers         []*CreateForSubscriptionItemTier       `json:"item_tiers,omitempty"`
	CouponsToAdd      []*CreateForSubscriptionCouponsToAdd   `json:"coupons_to_add,omitempty"`
	DiscountsToAdd    []*CreateForSubscriptionDiscountsToAdd `json:"discounts_to_add,omitempty"`
	ContractTerm      *CreateForSubscriptionContractTerm     `json:"contract_term,omitempty"`
}
type CreateForSubscriptionItemsToAdd struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type CreateForSubscriptionItemsToUpdate struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type CreateForSubscriptionItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CreateForSubscriptionCouponsToAdd struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type CreateForSubscriptionDiscountsToAdd struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type CreateForSubscriptionContractTerm struct {
	ActionAtTermEnd          ramp.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
	RenewalBillingCycles     *int32                           `json:"renewal_billing_cycles,omitempty"`
}
type UpdateRequest struct {
	EffectiveFrom     *int64                  `json:"effective_from"`
	Description       string                  `json:"description,omitempty"`
	CouponsToRemove   []string                `json:"coupons_to_remove,omitempty"`
	DiscountsToRemove []string                `json:"discounts_to_remove,omitempty"`
	ItemsToRemove     []string                `json:"items_to_remove,omitempty"`
	ItemsToAdd        []*UpdateItemsToAdd     `json:"items_to_add,omitempty"`
	ItemsToUpdate     []*UpdateItemsToUpdate  `json:"items_to_update,omitempty"`
	ItemTiers         []*UpdateItemTier       `json:"item_tiers,omitempty"`
	CouponsToAdd      []*UpdateCouponsToAdd   `json:"coupons_to_add,omitempty"`
	DiscountsToAdd    []*UpdateDiscountsToAdd `json:"discounts_to_add,omitempty"`
	ContractTerm      *UpdateContractTerm     `json:"contract_term,omitempty"`
}
type UpdateItemsToAdd struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type UpdateItemsToUpdate struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
}
type UpdateItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type UpdateCouponsToAdd struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}
type UpdateDiscountsToAdd struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
}
type UpdateContractTerm struct {
	ActionAtTermEnd          ramp.ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                           `json:"cancellation_cutoff_period,omitempty"`
	RenewalBillingCycles     *int32                           `json:"renewal_billing_cycles,omitempty"`
}
type ListRequest struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id"`
	EffectiveFrom  *filter.TimestampFilter `json:"effective_from,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}

type CreateForSubscriptionResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
}

type UpdateResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
}

type RetrieveResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
}

type DeleteResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
}

type ListRampResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
}

type ListResponse struct {
	List       []*ListRampResponse `json:"list,omitempty"`
	NextOffset string              `json:"next_offset,omitempty"`
}
