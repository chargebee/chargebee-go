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
	Object                 string                                   `json:"object"`
}
type Addon struct {
	Id                     string `json:"id"`
	Quantity               int32  `json:"quantity"`
	UnitPrice              int32  `json:"unit_price"`
	Amount                 int32  `json:"amount"`
	TrialEnd               int64  `json:"trial_end"`
	RemainingBillingCycles int32  `json:"remaining_billing_cycles"`
	Object                 string `json:"object"`
}
type EventBasedAddon struct {
	Id                  string       `json:"id"`
	Quantity            int32        `json:"quantity"`
	UnitPrice           int32        `json:"unit_price"`
	ServicePeriodInDays int32        `json:"service_period_in_days"`
	OnEvent             enum.OnEvent `json:"on_event"`
	ChargeOnce          bool         `json:"charge_once"`
	Object              string       `json:"object"`
}
type Coupon struct {
	CouponId     string `json:"coupon_id"`
	ApplyTill    int64  `json:"apply_till"`
	AppliedCount int32  `json:"applied_count"`
	CouponCode   string `json:"coupon_code"`
	Object       string `json:"object"`
}
