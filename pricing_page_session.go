package chargebee

type PricingPageSession struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	ExpiresAt int64  `json:"expires_at"`
	Object    string `json:"object"`
}
type CreateForNewSubscriptionRequest struct {
	RedirectUrl             string                                   `json:"redirect_url,omitempty"`
	PricingPage             *CreateForNewSubscriptionPricingPage     `json:"pricing_page,omitempty"`
	Subscription            *CreateForNewSubscriptionSubscription    `json:"subscription,omitempty"`
	BusinessEntityId        string                                   `json:"business_entity_id,omitempty"`
	AutoSelectLocalCurrency *bool                                    `json:"auto_select_local_currency,omitempty"`
	Customer                *CreateForNewSubscriptionCustomer        `json:"customer,omitempty"`
	Discounts               []*CreateForNewSubscriptionDiscount      `json:"discounts,omitempty"`
	BillingAddress          *CreateForNewSubscriptionBillingAddress  `json:"billing_address,omitempty"`
	ShippingAddress         *CreateForNewSubscriptionShippingAddress `json:"shipping_address,omitempty"`
}
type CreateForNewSubscriptionPricingPage struct {
	Id string `json:"id"`
}
type CreateForNewSubscriptionSubscription struct {
	Id string `json:"id,omitempty"`
}
type CreateForNewSubscriptionCustomer struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Locale    string `json:"locale,omitempty"`
}
type CreateForNewSubscriptionDiscount struct {
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
type CreateForNewSubscriptionBillingAddress struct {
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
type CreateForNewSubscriptionShippingAddress struct {
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
type CreateForExistingSubscriptionRequest struct {
	RedirectUrl  string                                     `json:"redirect_url,omitempty"`
	PricingPage  *CreateForExistingSubscriptionPricingPage  `json:"pricing_page,omitempty"`
	Subscription *CreateForExistingSubscriptionSubscription `json:"subscription,omitempty"`
	Discounts    []*CreateForExistingSubscriptionDiscount   `json:"discounts,omitempty"`
}
type CreateForExistingSubscriptionPricingPage struct {
	Id string `json:"id"`
}
type CreateForExistingSubscriptionSubscription struct {
	Id string `json:"id"`
}
type CreateForExistingSubscriptionDiscount struct {
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

type CreateForNewSubscriptionResponse struct {
	PricingPageSession *PricingPageSession `json:"pricing_page_session,omitempty"`
}

type CreateForExistingSubscriptionResponse struct {
	PricingPageSession *PricingPageSession `json:"pricing_page_session,omitempty"`
}
