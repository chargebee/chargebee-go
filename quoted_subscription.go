package chargebee

type QuotedSubscriptionChangeOption string

const (
	QuotedSubscriptionChangeOptionEndOfTerm    QuotedSubscriptionChangeOption = "end_of_term"
	QuotedSubscriptionChangeOptionSpecificDate QuotedSubscriptionChangeOption = "specific_date"
	QuotedSubscriptionChangeOptionImmediately  QuotedSubscriptionChangeOption = "immediately"
)

type QuotedSubscriptionAutoCollection string

const (
	QuotedSubscriptionAutoCollectionOn  QuotedSubscriptionAutoCollection = "on"
	QuotedSubscriptionAutoCollectionOff QuotedSubscriptionAutoCollection = "off"
)

type QuotedSubscriptionBillingPeriodUnit string

const (
	QuotedSubscriptionBillingPeriodUnitDay   QuotedSubscriptionBillingPeriodUnit = "day"
	QuotedSubscriptionBillingPeriodUnitWeek  QuotedSubscriptionBillingPeriodUnit = "week"
	QuotedSubscriptionBillingPeriodUnitMonth QuotedSubscriptionBillingPeriodUnit = "month"
	QuotedSubscriptionBillingPeriodUnitYear  QuotedSubscriptionBillingPeriodUnit = "year"
)

type QuotedSubscriptionSubscriptionItemItemType string

const (
	QuotedSubscriptionSubscriptionItemItemTypePlan   QuotedSubscriptionSubscriptionItemItemType = "plan"
	QuotedSubscriptionSubscriptionItemItemTypeAddon  QuotedSubscriptionSubscriptionItemItemType = "addon"
	QuotedSubscriptionSubscriptionItemItemTypeCharge QuotedSubscriptionSubscriptionItemItemType = "charge"
)

type QuotedSubscriptionSubscriptionItemBillingPeriodUnit string

const (
	QuotedSubscriptionSubscriptionItemBillingPeriodUnitDay   QuotedSubscriptionSubscriptionItemBillingPeriodUnit = "day"
	QuotedSubscriptionSubscriptionItemBillingPeriodUnitWeek  QuotedSubscriptionSubscriptionItemBillingPeriodUnit = "week"
	QuotedSubscriptionSubscriptionItemBillingPeriodUnitMonth QuotedSubscriptionSubscriptionItemBillingPeriodUnit = "month"
	QuotedSubscriptionSubscriptionItemBillingPeriodUnitYear  QuotedSubscriptionSubscriptionItemBillingPeriodUnit = "year"
)

type QuotedSubscriptionSubscriptionItemChargeOnEvent string

const (
	QuotedSubscriptionSubscriptionItemChargeOnEventSubscriptionCreation   QuotedSubscriptionSubscriptionItemChargeOnEvent = "subscription_creation"
	QuotedSubscriptionSubscriptionItemChargeOnEventSubscriptionTrialStart QuotedSubscriptionSubscriptionItemChargeOnEvent = "subscription_trial_start"
	QuotedSubscriptionSubscriptionItemChargeOnEventPlanActivation         QuotedSubscriptionSubscriptionItemChargeOnEvent = "plan_activation"
	QuotedSubscriptionSubscriptionItemChargeOnEventSubscriptionActivation QuotedSubscriptionSubscriptionItemChargeOnEvent = "subscription_activation"
	QuotedSubscriptionSubscriptionItemChargeOnEventContractTermination    QuotedSubscriptionSubscriptionItemChargeOnEvent = "contract_termination"
)

type QuotedSubscriptionSubscriptionItemChargeOnOption string

const (
	QuotedSubscriptionSubscriptionItemChargeOnOptionImmediately QuotedSubscriptionSubscriptionItemChargeOnOption = "immediately"
	QuotedSubscriptionSubscriptionItemChargeOnOptionOnEvent     QuotedSubscriptionSubscriptionItemChargeOnOption = "on_event"
)

type QuotedSubscriptionSubscriptionItemProrationType string

const (
	QuotedSubscriptionSubscriptionItemProrationTypeFullTerm    QuotedSubscriptionSubscriptionItemProrationType = "full_term"
	QuotedSubscriptionSubscriptionItemProrationTypePartialTerm QuotedSubscriptionSubscriptionItemProrationType = "partial_term"
	QuotedSubscriptionSubscriptionItemProrationTypeNone        QuotedSubscriptionSubscriptionItemProrationType = "none"
)

type QuotedSubscriptionSubscriptionItemUsageAccumulationResetFrequency string

const (
	QuotedSubscriptionSubscriptionItemUsageAccumulationResetFrequencyNever                        QuotedSubscriptionSubscriptionItemUsageAccumulationResetFrequency = "never"
	QuotedSubscriptionSubscriptionItemUsageAccumulationResetFrequencySubscriptionBillingFrequency QuotedSubscriptionSubscriptionItemUsageAccumulationResetFrequency = "subscription_billing_frequency"
)

type QuotedSubscriptionItemTierPricingType string

const (
	QuotedSubscriptionItemTierPricingTypePerUnit QuotedSubscriptionItemTierPricingType = "per_unit"
	QuotedSubscriptionItemTierPricingTypeFlatFee QuotedSubscriptionItemTierPricingType = "flat_fee"
	QuotedSubscriptionItemTierPricingTypePackage QuotedSubscriptionItemTierPricingType = "package"
)

type QuotedSubscriptionQuotedContractTermActionAtTermEnd string

const (
	QuotedSubscriptionQuotedContractTermActionAtTermEndRenew     QuotedSubscriptionQuotedContractTermActionAtTermEnd = "renew"
	QuotedSubscriptionQuotedContractTermActionAtTermEndEvergreen QuotedSubscriptionQuotedContractTermActionAtTermEnd = "evergreen"
	QuotedSubscriptionQuotedContractTermActionAtTermEndCancel    QuotedSubscriptionQuotedContractTermActionAtTermEnd = "cancel"
	QuotedSubscriptionQuotedContractTermActionAtTermEndRenewOnce QuotedSubscriptionQuotedContractTermActionAtTermEnd = "renew_once"
)

type QuotedSubscriptionEventBasedAddonOnEvent string

const (
	QuotedSubscriptionEventBasedAddonOnEventSubscriptionCreation   QuotedSubscriptionEventBasedAddonOnEvent = "subscription_creation"
	QuotedSubscriptionEventBasedAddonOnEventSubscriptionTrialStart QuotedSubscriptionEventBasedAddonOnEvent = "subscription_trial_start"
	QuotedSubscriptionEventBasedAddonOnEventPlanActivation         QuotedSubscriptionEventBasedAddonOnEvent = "plan_activation"
	QuotedSubscriptionEventBasedAddonOnEventSubscriptionActivation QuotedSubscriptionEventBasedAddonOnEvent = "subscription_activation"
	QuotedSubscriptionEventBasedAddonOnEventContractTermination    QuotedSubscriptionEventBasedAddonOnEvent = "contract_termination"
)

type QuotedSubscriptionAddonProrationType string

const (
	QuotedSubscriptionAddonProrationTypeFullTerm    QuotedSubscriptionAddonProrationType = "full_term"
	QuotedSubscriptionAddonProrationTypePartialTerm QuotedSubscriptionAddonProrationType = "partial_term"
	QuotedSubscriptionAddonProrationTypeNone        QuotedSubscriptionAddonProrationType = "none"
)

// just struct
type QuotedSubscription struct {
	Id                                string                                `json:"id"`
	PlanId                            string                                `json:"plan_id"`
	PlanQuantity                      int32                                 `json:"plan_quantity"`
	PlanUnitPrice                     int64                                 `json:"plan_unit_price"`
	SetupFee                          int64                                 `json:"setup_fee"`
	BillingPeriod                     int32                                 `json:"billing_period"`
	BillingPeriodUnit                 QuotedSubscriptionBillingPeriodUnit   `json:"billing_period_unit"`
	StartDate                         int64                                 `json:"start_date"`
	TrialEnd                          int64                                 `json:"trial_end"`
	RemainingBillingCycles            int32                                 `json:"remaining_billing_cycles"`
	PoNumber                          string                                `json:"po_number"`
	AutoCollection                    QuotedSubscriptionAutoCollection      `json:"auto_collection"`
	PlanQuantityInDecimal             string                                `json:"plan_quantity_in_decimal"`
	PlanUnitPriceInDecimal            string                                `json:"plan_unit_price_in_decimal"`
	ChangesScheduledAt                int64                                 `json:"changes_scheduled_at"`
	ChangeOption                      QuotedSubscriptionChangeOption        `json:"change_option"`
	ContractTermBillingCycleOnRenewal int32                                 `json:"contract_term_billing_cycle_on_renewal"`
	Addons                            []*QuotedSubscriptionAddon            `json:"addons"`
	EventBasedAddons                  []*QuotedSubscriptionEventBasedAddon  `json:"event_based_addons"`
	Coupons                           []*QuotedSubscriptionCoupon           `json:"coupons"`
	SubscriptionItems                 []*QuotedSubscriptionSubscriptionItem `json:"subscription_items"`
	ItemTiers                         []*QuotedSubscriptionItemTier         `json:"item_tiers"`
	QuotedContractTerm                *QuotedSubscriptionQuotedContractTerm `json:"quoted_contract_term"`
	Object                            string                                `json:"object"`
}

// sub resources
type QuotedSubscriptionCoupon struct {
	CouponId string `json:"coupon_id"`
	Object   string `json:"object"`
}
type QuotedSubscriptionSubscriptionItem struct {
	ItemPriceId                     string                                                            `json:"item_price_id"`
	ItemType                        QuotedSubscriptionSubscriptionItemItemType                        `json:"item_type"`
	Quantity                        int32                                                             `json:"quantity"`
	QuantityInDecimal               string                                                            `json:"quantity_in_decimal"`
	MeteredQuantity                 string                                                            `json:"metered_quantity"`
	LastCalculatedAt                int64                                                             `json:"last_calculated_at"`
	UnitPrice                       int64                                                             `json:"unit_price"`
	UnitPriceInDecimal              string                                                            `json:"unit_price_in_decimal"`
	Amount                          int64                                                             `json:"amount"`
	CurrentTermStart                int64                                                             `json:"current_term_start"`
	CurrentTermEnd                  int64                                                             `json:"current_term_end"`
	NextBillingAt                   int64                                                             `json:"next_billing_at"`
	AmountInDecimal                 string                                                            `json:"amount_in_decimal"`
	BillingPeriod                   int32                                                             `json:"billing_period"`
	BillingPeriodUnit               QuotedSubscriptionSubscriptionItemBillingPeriodUnit               `json:"billing_period_unit"`
	FreeQuantity                    int32                                                             `json:"free_quantity"`
	FreeQuantityInDecimal           string                                                            `json:"free_quantity_in_decimal"`
	TrialEnd                        int64                                                             `json:"trial_end"`
	BillingCycles                   int32                                                             `json:"billing_cycles"`
	ServicePeriodDays               int32                                                             `json:"service_period_days"`
	ChargeOnEvent                   QuotedSubscriptionSubscriptionItemChargeOnEvent                   `json:"charge_on_event"`
	ChargeOnce                      bool                                                              `json:"charge_once"`
	ChargeOnOption                  QuotedSubscriptionSubscriptionItemChargeOnOption                  `json:"charge_on_option"`
	ProrationType                   QuotedSubscriptionSubscriptionItemProrationType                   `json:"proration_type"`
	UsageAccumulationResetFrequency QuotedSubscriptionSubscriptionItemUsageAccumulationResetFrequency `json:"usage_accumulation_reset_frequency"`
	Object                          string                                                            `json:"object"`
}
type QuotedSubscriptionItemTier struct {
	ItemPriceId           string                                `json:"item_price_id"`
	StartingUnit          int32                                 `json:"starting_unit"`
	EndingUnit            int32                                 `json:"ending_unit"`
	Price                 int64                                 `json:"price"`
	StartingUnitInDecimal string                                `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string                                `json:"ending_unit_in_decimal"`
	PriceInDecimal        string                                `json:"price_in_decimal"`
	PricingType           QuotedSubscriptionItemTierPricingType `json:"pricing_type"`
	PackageSize           int32                                 `json:"package_size"`
	Index                 int32                                 `json:"index"`
	Object                string                                `json:"object"`
}
type QuotedSubscriptionQuotedContractTerm struct {
	ContractStart            int64                                               `json:"contract_start"`
	ContractEnd              int64                                               `json:"contract_end"`
	BillingCycle             int32                                               `json:"billing_cycle"`
	ActionAtTermEnd          QuotedSubscriptionQuotedContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue       int64                                               `json:"total_contract_value"`
	CancellationCutoffPeriod int32                                               `json:"cancellation_cutoff_period"`
	Object                   string                                              `json:"object"`
}
type QuotedSubscriptionEventBasedAddon struct {
	Id                  string                                   `json:"id"`
	Quantity            int32                                    `json:"quantity"`
	UnitPrice           int64                                    `json:"unit_price"`
	ServicePeriodInDays int32                                    `json:"service_period_in_days"`
	OnEvent             QuotedSubscriptionEventBasedAddonOnEvent `json:"on_event"`
	ChargeOnce          bool                                     `json:"charge_once"`
	QuantityInDecimal   string                                   `json:"quantity_in_decimal"`
	UnitPriceInDecimal  string                                   `json:"unit_price_in_decimal"`
	Object              string                                   `json:"object"`
}
type QuotedSubscriptionAddon struct {
	Id                     string                               `json:"id"`
	Quantity               int32                                `json:"quantity"`
	UnitPrice              int64                                `json:"unit_price"`
	Amount                 int64                                `json:"amount"`
	TrialEnd               int64                                `json:"trial_end"`
	RemainingBillingCycles int32                                `json:"remaining_billing_cycles"`
	QuantityInDecimal      string                               `json:"quantity_in_decimal"`
	UnitPriceInDecimal     string                               `json:"unit_price_in_decimal"`
	AmountInDecimal        string                               `json:"amount_in_decimal"`
	ProrationType          QuotedSubscriptionAddonProrationType `json:"proration_type"`
	Object                 string                               `json:"object"`
}

// operations
// input params
