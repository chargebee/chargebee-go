package quotedsubscription

import (
	"github.com/chargebee/chargebee-go/enum"
	quotedSubscriptionEnum "github.com/chargebee/chargebee-go/models/quotedsubscription/enum"
)

type QuotedSubscription struct {
	Id                     string                                   `json:"id"`
	PlanId                 string                                   `json:"plan_id"`
	PlanQuantity           int32                                    `json:"plan_quantity"`
	PlanUnitPrice          int32                                    `json:"plan_unit_price"`
	SetupFee               int32                                    `json:"setup_fee"`
	BillingPeriod          int32                                    `json:"billing_period"`
	BillingPeriodUnit      quotedSubscriptionEnum.BillingPeriodUnit `json:"billing_period_unit"`
	StartDate              int64                                    `json:"start_date"`
	TrialEnd               int64                                    `json:"trial_end"`
	RemainingBillingCycles int32                                    `json:"remaining_billing_cycles"`
	PoNumber               string                                   `json:"po_number"`
	AutoCollection         enum.AutoCollection                      `json:"auto_collection"`
	Addons                 []*Addon                                 `json:"addons"`
	EventBasedAddons       []*EventBasedAddon                       `json:"event_based_addons"`
	Coupons                []*Coupon                                `json:"coupons"`
	SubscriptionItems      []*SubscriptionItem                      `json:"subscription_items"`
	ItemTiers              []*ItemTier                              `json:"item_tiers"`
	Object                 string                                   `json:"object"`
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
	CouponId     string `json:"coupon_id"`
	ApplyTill    int64  `json:"apply_till"`
	AppliedCount int32  `json:"applied_count"`
	CouponCode   string `json:"coupon_code"`
	Object       string `json:"object"`
}
type SubscriptionItem struct {
	ItemPriceId       string              `json:"item_price_id"`
	ItemType          enum.ItemType       `json:"item_type"`
	Quantity          int32               `json:"quantity"`
	UnitPrice         int32               `json:"unit_price"`
	Amount            int32               `json:"amount"`
	FreeQuantity      int32               `json:"free_quantity"`
	TrialEnd          int64               `json:"trial_end"`
	BillingCycles     int32               `json:"billing_cycles"`
	ServicePeriodDays int32               `json:"service_period_days"`
	ChargeOnEvent     enum.ChargeOnEvent  `json:"charge_on_event"`
	ChargeOnce        bool                `json:"charge_once"`
	ChargeOnOption    enum.ChargeOnOption `json:"charge_on_option"`
	Object            string              `json:"object"`
}
type ItemTier struct {
	ItemPriceId  string `json:"item_price_id"`
	StartingUnit int32  `json:"starting_unit"`
	EndingUnit   int32  `json:"ending_unit"`
	Price        int32  `json:"price"`
	Object       string `json:"object"`
}
