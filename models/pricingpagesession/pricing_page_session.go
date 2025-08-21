package pricingpagesession

import (
	"github.com/chargebee/chargebee-go/v3/enum"
)

type PricingPageSession struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	ExpiresAt int64  `json:"expires_at"`
	Object    string `json:"object"`
}
type CreateForNewSubscriptionRequestParams struct {
	RedirectUrl      string                                         `json:"redirect_url,omitempty"`
	PricingPage      *CreateForNewSubscriptionPricingPageParams     `json:"pricing_page,omitempty"`
	Subscription     *CreateForNewSubscriptionSubscriptionParams    `json:"subscription,omitempty"`
	BusinessEntityId string                                         `json:"business_entity_id,omitempty"`
	Customer         *CreateForNewSubscriptionCustomerParams        `json:"customer,omitempty"`
	Discounts        []*CreateForNewSubscriptionDiscountParams      `json:"discounts,omitempty"`
	BillingAddress   *CreateForNewSubscriptionBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress  *CreateForNewSubscriptionShippingAddressParams `json:"shipping_address,omitempty"`
}
type CreateForNewSubscriptionPricingPageParams struct {
	Id string `json:"id"`
}
type CreateForNewSubscriptionSubscriptionParams struct {
	Id string `json:"id,omitempty"`
}
type CreateForNewSubscriptionCustomerParams struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Locale    string `json:"locale,omitempty"`
}
type CreateForNewSubscriptionDiscountParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on,omitempty"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
	Quantity      *int32            `json:"quantity,omitempty"`
	Label         string            `json:"label,omitempty"`
}
type CreateForNewSubscriptionBillingAddressParams struct {
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
type CreateForNewSubscriptionShippingAddressParams struct {
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
type CreateForExistingSubscriptionRequestParams struct {
	RedirectUrl  string                                           `json:"redirect_url,omitempty"`
	PricingPage  *CreateForExistingSubscriptionPricingPageParams  `json:"pricing_page,omitempty"`
	Subscription *CreateForExistingSubscriptionSubscriptionParams `json:"subscription,omitempty"`
	Discounts    []*CreateForExistingSubscriptionDiscountParams   `json:"discounts,omitempty"`
}
type CreateForExistingSubscriptionPricingPageParams struct {
	Id string `json:"id"`
}
type CreateForExistingSubscriptionSubscriptionParams struct {
	Id string `json:"id"`
}
type CreateForExistingSubscriptionDiscountParams struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on,omitempty"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
	Quantity      *int32            `json:"quantity,omitempty"`
	Label         string            `json:"label,omitempty"`
}
