package chargebee

type RampStatus string

const (
	RampStatusScheduled RampStatus = "scheduled"
	RampStatusSucceeded RampStatus = "succeeded"
	RampStatusFailed    RampStatus = "failed"
	RampStatusDraft     RampStatus = "draft"
)

type RampDiscountsToAddType string

const (
	RampDiscountsToAddTypeFixedAmount RampDiscountsToAddType = "fixed_amount"
	RampDiscountsToAddTypePercentage  RampDiscountsToAddType = "percentage"
)

type RampContractTermActionAtTermEnd string

const (
	RampContractTermActionAtTermEndRenew     RampContractTermActionAtTermEnd = "renew"
	RampContractTermActionAtTermEndEvergreen RampContractTermActionAtTermEnd = "evergreen"
	RampContractTermActionAtTermEndCancel    RampContractTermActionAtTermEnd = "cancel"
	RampContractTermActionAtTermEndRenewOnce RampContractTermActionAtTermEnd = "renew_once"
)

// just struct
type Ramp struct {
	Id                     string                      `json:"id"`
	Description            string                      `json:"description"`
	SubscriptionId         string                      `json:"subscription_id"`
	EffectiveFrom          int64                       `json:"effective_from"`
	Status                 RampStatus                  `json:"status"`
	CreatedAt              int64                       `json:"created_at"`
	ResourceVersion        int64                       `json:"resource_version"`
	UpdatedAt              int64                       `json:"updated_at"`
	ItemsToAdd             []*RampItemsToAdd           `json:"items_to_add"`
	ItemsToUpdate          []*RampItemsToUpdate        `json:"items_to_update"`
	CouponsToAdd           []*RampCouponsToAdd         `json:"coupons_to_add"`
	DiscountsToAdd         []*RampDiscountsToAdd       `json:"discounts_to_add"`
	ItemTiers              []*RampItemTier             `json:"item_tiers"`
	ItemsToRemove          []string                    `json:"items_to_remove"`
	CouponsToRemove        []string                    `json:"coupons_to_remove"`
	DiscountsToRemove      []string                    `json:"discounts_to_remove"`
	ContractTerm           *RampContractTerm           `json:"contract_term"`
	Deleted                bool                        `json:"deleted"`
	StatusTransitionReason *RampStatusTransitionReason `json:"status_transition_reason"`
	Object                 string                      `json:"object"`
}

// sub resources
type RampItemsToAdd struct {
	ItemPriceId           string         `json:"item_price_id"`
	ItemType              ItemType       `json:"item_type"`
	Quantity              int32          `json:"quantity"`
	QuantityInDecimal     string         `json:"quantity_in_decimal"`
	UnitPrice             int64          `json:"unit_price"`
	UnitPriceInDecimal    string         `json:"unit_price_in_decimal"`
	Amount                int64          `json:"amount"`
	AmountInDecimal       string         `json:"amount_in_decimal"`
	FreeQuantity          int32          `json:"free_quantity"`
	FreeQuantityInDecimal string         `json:"free_quantity_in_decimal"`
	BillingCycles         int32          `json:"billing_cycles"`
	ServicePeriodDays     int32          `json:"service_period_days"`
	MeteredQuantity       string         `json:"metered_quantity"`
	ChargeOnce            bool           `json:"charge_once"`
	ChargeOnOption        ChargeOnOption `json:"charge_on_option"`
	ChargeOnEvent         ChargeOnEvent  `json:"charge_on_event"`
	Object                string         `json:"object"`
}

type RampItemsToUpdate struct {
	ItemPriceId           string         `json:"item_price_id"`
	ItemType              ItemType       `json:"item_type"`
	Quantity              int32          `json:"quantity"`
	QuantityInDecimal     string         `json:"quantity_in_decimal"`
	UnitPrice             int64          `json:"unit_price"`
	UnitPriceInDecimal    string         `json:"unit_price_in_decimal"`
	Amount                int64          `json:"amount"`
	AmountInDecimal       string         `json:"amount_in_decimal"`
	FreeQuantity          int32          `json:"free_quantity"`
	FreeQuantityInDecimal string         `json:"free_quantity_in_decimal"`
	BillingCycles         int32          `json:"billing_cycles"`
	ServicePeriodDays     int32          `json:"service_period_days"`
	MeteredQuantity       string         `json:"metered_quantity"`
	ChargeOnce            bool           `json:"charge_once"`
	ChargeOnOption        ChargeOnOption `json:"charge_on_option"`
	ChargeOnEvent         ChargeOnEvent  `json:"charge_on_event"`
	Object                string         `json:"object"`
}

type RampCouponsToAdd struct {
	CouponId  string `json:"coupon_id"`
	ApplyTill int64  `json:"apply_till"`
	Object    string `json:"object"`
}

type RampDiscountsToAdd struct {
	Id            string                 `json:"id"`
	InvoiceName   string                 `json:"invoice_name"`
	Type          RampDiscountsToAddType `json:"type"`
	Percentage    float64                `json:"percentage"`
	Amount        int64                  `json:"amount"`
	DurationType  DurationType           `json:"duration_type"`
	Period        int32                  `json:"period"`
	PeriodUnit    PeriodUnit             `json:"period_unit"`
	IncludedInMrr bool                   `json:"included_in_mrr"`
	ApplyOn       ApplyOn                `json:"apply_on"`
	ItemPriceId   string                 `json:"item_price_id"`
	CreatedAt     int64                  `json:"created_at"`
	Object        string                 `json:"object"`
}

type RampItemTier struct {
	ItemPriceId           string      `json:"item_price_id"`
	StartingUnit          int32       `json:"starting_unit"`
	EndingUnit            int32       `json:"ending_unit"`
	Price                 int64       `json:"price"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal"`
	PriceInDecimal        string      `json:"price_in_decimal"`
	PricingType           PricingType `json:"pricing_type"`
	PackageSize           int32       `json:"package_size"`
	Index                 int32       `json:"index"`
	Object                string      `json:"object"`
}

type RampContractTerm struct {
	CancellationCutoffPeriod int32                           `json:"cancellation_cutoff_period"`
	RenewalBillingCycles     int32                           `json:"renewal_billing_cycles"`
	ActionAtTermEnd          RampContractTermActionAtTermEnd `json:"action_at_term_end"`
	Object                   string                          `json:"object"`
}

type RampStatusTransitionReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Object  string `json:"object"`
}

// operations
// input params
type RampCreateForSubscriptionRequest struct {
	EffectiveFrom     *int64                                     `json:"effective_from"`
	Description       string                                     `json:"description,omitempty"`
	CouponsToRemove   []string                                   `json:"coupons_to_remove,omitempty"`
	DiscountsToRemove []string                                   `json:"discounts_to_remove,omitempty"`
	ItemsToRemove     []string                                   `json:"items_to_remove,omitempty"`
	ItemsToAdd        []*RampCreateForSubscriptionItemsToAdd     `json:"items_to_add,omitempty"`
	ItemsToUpdate     []*RampCreateForSubscriptionItemsToUpdate  `json:"items_to_update,omitempty"`
	ItemTiers         []*RampCreateForSubscriptionItemTier       `json:"item_tiers,omitempty"`
	CouponsToAdd      []*RampCreateForSubscriptionCouponsToAdd   `json:"coupons_to_add,omitempty"`
	DiscountsToAdd    []*RampCreateForSubscriptionDiscountsToAdd `json:"discounts_to_add,omitempty"`
	ContractTerm      *RampCreateForSubscriptionContractTerm     `json:"contract_term,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *RampCreateForSubscriptionRequest) payload() any { return r }

// input sub resource params multi
type RampCreateForSubscriptionItemsToAdd struct {
	ItemPriceId        string         `json:"item_price_id"`
	Quantity           *int32         `json:"quantity,omitempty"`
	QuantityInDecimal  string         `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64         `json:"unit_price,omitempty"`
	UnitPriceInDecimal string         `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32         `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32         `json:"service_period_days,omitempty"`
	ChargeOnEvent      ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool          `json:"charge_once,omitempty"`
	ChargeOnOption     ChargeOnOption `json:"charge_on_option,omitempty"`
}

// input sub resource params multi
type RampCreateForSubscriptionItemsToUpdate struct {
	ItemPriceId        string         `json:"item_price_id"`
	Quantity           *int32         `json:"quantity,omitempty"`
	QuantityInDecimal  string         `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64         `json:"unit_price,omitempty"`
	UnitPriceInDecimal string         `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32         `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32         `json:"service_period_days,omitempty"`
	ChargeOnEvent      ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool          `json:"charge_once,omitempty"`
	ChargeOnOption     ChargeOnOption `json:"charge_on_option,omitempty"`
}

// input sub resource params multi
type RampCreateForSubscriptionItemTier struct {
	ItemPriceId           string      `json:"item_price_id,omitempty"`
	StartingUnit          *int32      `json:"starting_unit,omitempty"`
	EndingUnit            *int32      `json:"ending_unit,omitempty"`
	Price                 *int64      `json:"price,omitempty"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string      `json:"price_in_decimal,omitempty"`
	PricingType           PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32      `json:"package_size,omitempty"`
}

// input sub resource params multi
type RampCreateForSubscriptionCouponsToAdd struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

// input sub resource params multi
type RampCreateForSubscriptionDiscountsToAdd struct {
	ApplyOn       ApplyOn      `json:"apply_on"`
	DurationType  DurationType `json:"duration_type"`
	Percentage    *float64     `json:"percentage,omitempty"`
	Amount        *int64       `json:"amount,omitempty"`
	Period        *int32       `json:"period,omitempty"`
	PeriodUnit    PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool        `json:"included_in_mrr,omitempty"`
	ItemPriceId   string       `json:"item_price_id,omitempty"`
}

// input sub resource params single
type RampCreateForSubscriptionContractTerm struct {
	ActionAtTermEnd          ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                      `json:"cancellation_cutoff_period,omitempty"`
	RenewalBillingCycles     *int32                      `json:"renewal_billing_cycles,omitempty"`
}

type RampUpdateRequest struct {
	EffectiveFrom     *int64                      `json:"effective_from"`
	Description       string                      `json:"description,omitempty"`
	CouponsToRemove   []string                    `json:"coupons_to_remove,omitempty"`
	DiscountsToRemove []string                    `json:"discounts_to_remove,omitempty"`
	ItemsToRemove     []string                    `json:"items_to_remove,omitempty"`
	ItemsToAdd        []*RampUpdateItemsToAdd     `json:"items_to_add,omitempty"`
	ItemsToUpdate     []*RampUpdateItemsToUpdate  `json:"items_to_update,omitempty"`
	ItemTiers         []*RampUpdateItemTier       `json:"item_tiers,omitempty"`
	CouponsToAdd      []*RampUpdateCouponsToAdd   `json:"coupons_to_add,omitempty"`
	DiscountsToAdd    []*RampUpdateDiscountsToAdd `json:"discounts_to_add,omitempty"`
	ContractTerm      *RampUpdateContractTerm     `json:"contract_term,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *RampUpdateRequest) payload() any { return r }

// input sub resource params multi
type RampUpdateItemsToAdd struct {
	ItemPriceId        string         `json:"item_price_id"`
	Quantity           *int32         `json:"quantity,omitempty"`
	QuantityInDecimal  string         `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64         `json:"unit_price,omitempty"`
	UnitPriceInDecimal string         `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32         `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32         `json:"service_period_days,omitempty"`
	ChargeOnEvent      ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool          `json:"charge_once,omitempty"`
	ChargeOnOption     ChargeOnOption `json:"charge_on_option,omitempty"`
}

// input sub resource params multi
type RampUpdateItemsToUpdate struct {
	ItemPriceId        string         `json:"item_price_id"`
	Quantity           *int32         `json:"quantity,omitempty"`
	QuantityInDecimal  string         `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64         `json:"unit_price,omitempty"`
	UnitPriceInDecimal string         `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32         `json:"billing_cycles,omitempty"`
	ServicePeriodDays  *int32         `json:"service_period_days,omitempty"`
	ChargeOnEvent      ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool          `json:"charge_once,omitempty"`
	ChargeOnOption     ChargeOnOption `json:"charge_on_option,omitempty"`
}

// input sub resource params multi
type RampUpdateItemTier struct {
	ItemPriceId           string      `json:"item_price_id,omitempty"`
	StartingUnit          *int32      `json:"starting_unit,omitempty"`
	EndingUnit            *int32      `json:"ending_unit,omitempty"`
	Price                 *int64      `json:"price,omitempty"`
	StartingUnitInDecimal string      `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string      `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string      `json:"price_in_decimal,omitempty"`
	PricingType           PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32      `json:"package_size,omitempty"`
}

// input sub resource params multi
type RampUpdateCouponsToAdd struct {
	CouponId  string `json:"coupon_id,omitempty"`
	ApplyTill *int64 `json:"apply_till,omitempty"`
}

// input sub resource params multi
type RampUpdateDiscountsToAdd struct {
	ApplyOn       ApplyOn      `json:"apply_on"`
	DurationType  DurationType `json:"duration_type"`
	Percentage    *float64     `json:"percentage,omitempty"`
	Amount        *int64       `json:"amount,omitempty"`
	Period        *int32       `json:"period,omitempty"`
	PeriodUnit    PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool        `json:"included_in_mrr,omitempty"`
	ItemPriceId   string       `json:"item_price_id,omitempty"`
}

// input sub resource params single
type RampUpdateContractTerm struct {
	ActionAtTermEnd          ContractTermActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                      `json:"cancellation_cutoff_period,omitempty"`
	RenewalBillingCycles     *int32                      `json:"renewal_billing_cycles,omitempty"`
}

type RampListRequest struct {
	Limit          *int32           `json:"limit,omitempty"`
	Offset         string           `json:"offset,omitempty"`
	IncludeDeleted *bool            `json:"include_deleted,omitempty"`
	Status         *EnumFilter      `json:"status,omitempty"`
	SubscriptionId *StringFilter    `json:"subscription_id"`
	EffectiveFrom  *TimestampFilter `json:"effective_from,omitempty"`
	UpdatedAt      *TimestampFilter `json:"updated_at,omitempty"`
	SortBy         *SortFilter      `json:"sort_by,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *RampListRequest) payload() any { return r }

// operation response
type RampCreateForSubscriptionResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
	apiResponse
}

// operation response
type RampUpdateResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
	apiResponse
}

// operation response
type RampRetrieveResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
	apiResponse
}

// operation response
type RampDeleteResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
	apiResponse
}

// operation sub response
type RampListRampResponse struct {
	Ramp *Ramp `json:"ramp,omitempty"`
}

// operation response
type RampListResponse struct {
	List       []*RampListRampResponse `json:"list,omitempty"`
	NextOffset string                  `json:"next_offset,omitempty"`
	apiResponse
}
