package gift

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	giftEnum "github.com/chargebee/chargebee-go/models/gift/enum"
)

type Gift struct {
	Id              string          `json:"id"`
	Status          giftEnum.Status `json:"status"`
	ScheduledAt     int64           `json:"scheduled_at"`
	AutoClaim       bool            `json:"auto_claim"`
	NoExpiry        bool            `json:"no_expiry"`
	ClaimExpiryDate int64           `json:"claim_expiry_date"`
	ResourceVersion int64           `json:"resource_version"`
	UpdatedAt       int64           `json:"updated_at"`
	Gifter          *Gifter         `json:"gifter"`
	GiftReceiver    *GiftReceiver   `json:"gift_receiver"`
	GiftTimelines   []*GiftTimeline `json:"gift_timelines"`
	Object          string          `json:"object"`
}
type Gifter struct {
	CustomerId string `json:"customer_id"`
	InvoiceId  string `json:"invoice_id"`
	Signature  string `json:"signature"`
	Note       string `json:"note"`
	Object     string `json:"object"`
}
type GiftReceiver struct {
	CustomerId     string `json:"customer_id"`
	SubscriptionId string `json:"subscription_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Object         string `json:"object"`
}
type GiftTimeline struct {
	Status     enum.Status `json:"status"`
	OccurredAt int64       `json:"occurred_at"`
	Object     string      `json:"object"`
}
type CreateRequestParams struct {
	ScheduledAt     *int64                       `json:"scheduled_at,omitempty"`
	AutoClaim       *bool                        `json:"auto_claim,omitempty"`
	NoExpiry        *bool                        `json:"no_expiry,omitempty"`
	ClaimExpiryDate *int64                       `json:"claim_expiry_date,omitempty"`
	Gifter          *CreateGifterParams          `json:"gifter,omitempty"`
	GiftReceiver    *CreateGiftReceiverParams    `json:"gift_receiver,omitempty"`
	Subscription    *CreateSubscriptionParams    `json:"subscription,omitempty"`
	Addons          []*CreateAddonParams         `json:"addons,omitempty"`
	PaymentIntent   *CreatePaymentIntentParams   `json:"payment_intent,omitempty"`
	ShippingAddress *CreateShippingAddressParams `json:"shipping_address,omitempty"`
	CouponIds       []string                     `json:"coupon_ids,omitempty"`
}
type CreateGifterParams struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type CreateGiftReceiverParams struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}
type CreateSubscriptionParams struct {
	PlanId       string `json:"plan_id"`
	PlanQuantity *int32 `json:"plan_quantity,omitempty"`
}
type CreateAddonParams struct {
	Id       string `json:"id,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
}
type CreatePaymentIntentParams struct {
	Id                string `json:"id,omitempty"`
	GatewayAccountId  string `json:"gateway_account_id,omitempty"`
	GwToken           string `json:"gw_token,omitempty"`
	ReferenceId       string `json:"reference_id,omitempty"`
	GwPaymentMethodId string `json:"gw_payment_method_id,omitempty"`
}
type CreateShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CartGiftCreateSubscriptionParams struct {
	Id            string `json:"id,omitempty"`
	PlanId        string `json:"plan_id"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
	Coupon        string `json:"coupon,omitempty"`
}
type CartGiftCreateAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CartGiftCreateEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CartGiftCreateBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CartGiftCreateShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CartGiftCreateCustomerParams struct {
	VatNumber        string          `json:"vat_number,omitempty"`
	RegisteredForGst *bool           `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability `json:"taxability,omitempty"`
	EntityCode       enum.EntityCode `json:"entity_code,omitempty"`
	ExemptNumber     string          `json:"exempt_number,omitempty"`
}
type CartGiftCreateForCustomerSubscriptionParams struct {
	Id            string `json:"id,omitempty"`
	PlanId        string `json:"plan_id"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CartGiftCreateForCustomerAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CartGiftCreateForCustomerEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CartGiftCreateForCustomerShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type ListRequestParams struct {
	Limit        *int32                  `json:"limit,omitempty"`
	Offset       string                  `json:"offset,omitempty"`
	Status       *filter.EnumFilter      `json:"status,omitempty"`
	GiftReceiver *ListGiftReceiverParams `json:"gift_receiver,omitempty"`
	Gifter       *ListGifterParams       `json:"gifter,omitempty"`
}
type ListGiftReceiverParams struct {
	Email      *filter.StringFilter `json:"email,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
type ListGifterParams struct {
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
}
type UpdateGiftRequestParams struct {
	ScheduledAt *int64 `json:"scheduled_at"`
	Comment     string `json:"comment,omitempty"`
}
