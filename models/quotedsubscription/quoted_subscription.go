package quotedsubscription

import (
	"github.com/chargebee/chargebee-go/enum"
	quotedSubscriptionEnum "github.com/chargebee/chargebee-go/models/quotedsubscription/enum"
)

type QuotedSubscription struct {
	Id                                string                                   `json:"id"`
	PlanId                            string                                   `json:"plan_id"`
	PlanQuantity                      int32                                    `json:"plan_quantity"`
	PlanUnitPrice                     int32                                    `json:"plan_unit_price"`
	SetupFee                          int32                                    `json:"setup_fee"`
	BillingPeriod                     int32                                    `json:"billing_period"`
	BillingPeriodUnit                 quotedSubscriptionEnum.BillingPeriodUnit `json:"billing_period_unit"`
	StartDate                         int64                                    `json:"start_date"`
	TrialEnd                          int64                                    `json:"trial_end"`
	RemainingBillingCycles            int32                                    `json:"remaining_billing_cycles"`
	PoNumber                          string                                   `json:"po_number"`
	AutoCollection                    enum.AutoCollection                      `json:"auto_collection"`
	PlanQuantityInDecimal             string                                   `json:"plan_quantity_in_decimal"`
	PlanUnitPriceInDecimal            string                                   `json:"plan_unit_price_in_decimal"`
	ChangesScheduledAt                int64                                    `json:"changes_scheduled_at"`
	ChangeOption                      quotedSubscriptionEnum.ChangeOption      `json:"change_option"`
	ContractTermBillingCycleOnRenewal int32                                    `json:"contract_term_billing_cycle_on_renewal"`
	Addons                            []*Addon                                 `json:"addons"`
	EventBasedAddons                  []*EventBasedAddon                       `json:"event_based_addons"`
	Coupons                           []*Coupon                                `json:"coupons"`
	SubscriptionItems                 []*SubscriptionItem                      `json:"subscription_items"`
	ItemTiers                         []*ItemTier                              `json:"item_tiers"`
	QuotedContractTerm                *QuotedContractTerm                      `json:"quoted_contract_term"`
	Object                            string                                   `json:"object"`
}
type Addon struct {
	Id                     string `json:"id"`
	Quantity               int32  `json:"quantity"`
	UnitPrice              int32  `json:"unit_price"`
	Amount                 int32  `json:"amount"`
	TrialEnd               int64  `json:"trial_end"`
	RemainingBillingCycles int32  `json:"remaining_billing_cycles"`
	QuantityInDecimal      string `json:"quantity_in_decimal"`
	UnitPriceInDecimal     string `json:"unit_price_in_decimal"`
	AmountInDecimal        string `json:"amount_in_decimal"`
	Object                 string `json:"object"`
}
type EventBasedAddon struct {
	Id                  string       `json:"id"`
	Quantity            int32        `json:"quantity"`
	UnitPrice           int32        `json:"unit_price"`
	ServicePeriodInDays int32        `json:"service_period_in_days"`
	OnEvent             enum.OnEvent `json:"on_event"`
	ChargeOnce          bool         `json:"charge_once"`
	QuantityInDecimal   string       `json:"quantity_in_decimal"`
	UnitPriceInDecimal  string       `json:"unit_price_in_decimal"`
	Object              string       `json:"object"`
}
type Coupon struct {
	CouponId string `json:"coupon_id"`
	Object   string `json:"object"`
}
type SubscriptionItem struct {
	ItemPriceId           string              `json:"item_price_id"`
	ItemType              enum.ItemType       `json:"item_type"`
	Quantity              int32               `json:"quantity"`
	QuantityInDecimal     string              `json:"quantity_in_decimal"`
	MeteredQuantity       string              `json:"metered_quantity"`
	LastCalculatedAt      int64               `json:"last_calculated_at"`
	UnitPrice             int32               `json:"unit_price"`
	UnitPriceInDecimal    string              `json:"unit_price_in_decimal"`
	Amount                int32               `json:"amount"`
	AmountInDecimal       string              `json:"amount_in_decimal"`
	FreeQuantity          int32               `json:"free_quantity"`
	FreeQuantityInDecimal string              `json:"free_quantity_in_decimal"`
	TrialEnd              int64               `json:"trial_end"`
	BillingCycles         int32               `json:"billing_cycles"`
	ServicePeriodDays     int32               `json:"service_period_days"`
	ChargeOnEvent         enum.ChargeOnEvent  `json:"charge_on_event"`
	ChargeOnce            bool                `json:"charge_once"`
	ChargeOnOption        enum.ChargeOnOption `json:"charge_on_option"`
	Object                string              `json:"object"`
}
type ItemTier struct {
	ItemPriceId           string `json:"item_price_id"`
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	Price                 int32  `json:"price"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	PriceInDecimal        string `json:"price_in_decimal"`
	Object                string `json:"object"`
}
type QuotedContractTerm struct {
	ContractStart            int64                                                    `json:"contract_start"`
	ContractEnd              int64                                                    `json:"contract_end"`
	BillingCycle             int32                                                    `json:"billing_cycle"`
	ActionAtTermEnd          quotedSubscriptionEnum.QuotedContractTermActionAtTermEnd `json:"action_at_term_end"`
	TotalContractValue       int64                                                    `json:"total_contract_value"`
	CancellationCutoffPeriod int32                                                    `json:"cancellation_cutoff_period"`
	Object                   string                                                   `json:"object"`
}
