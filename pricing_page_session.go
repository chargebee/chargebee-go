package chargebee

// just struct
type PricingPageSession struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	ExpiresAt int64  `json:"expires_at"`
	Object    string `json:"object"`
}

// sub resources
// operations
// input params
type PricingPageSessionCreateForNewSubscriptionRequest struct {
	RedirectUrl             string                                                     `json:"redirect_url,omitempty"`
	PricingPage             *PricingPageSessionCreateForNewSubscriptionPricingPage     `json:"pricing_page,omitempty"`
	Subscription            *PricingPageSessionCreateForNewSubscriptionSubscription    `json:"subscription,omitempty"`
	BusinessEntityId        string                                                     `json:"business_entity_id,omitempty"`
	AutoSelectLocalCurrency *bool                                                      `json:"auto_select_local_currency,omitempty"`
	Customer                *PricingPageSessionCreateForNewSubscriptionCustomer        `json:"customer,omitempty"`
	Discounts               []*PricingPageSessionCreateForNewSubscriptionDiscount      `json:"discounts,omitempty"`
	BillingAddress          *PricingPageSessionCreateForNewSubscriptionBillingAddress  `json:"billing_address,omitempty"`
	ShippingAddress         *PricingPageSessionCreateForNewSubscriptionShippingAddress `json:"shipping_address,omitempty"`
	apiRequest              `json:"-" form:"-"`
}

func (r *PricingPageSessionCreateForNewSubscriptionRequest) payload() any { return r }

// input sub resource params single
type PricingPageSessionCreateForNewSubscriptionPricingPage struct {
	Id string `json:"id"`
}

// input sub resource params single
type PricingPageSessionCreateForNewSubscriptionSubscription struct {
	Id string `json:"id,omitempty"`
}

// input sub resource params single
type PricingPageSessionCreateForNewSubscriptionCustomer struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Locale    string `json:"locale,omitempty"`
}

// input sub resource params multi
type PricingPageSessionCreateForNewSubscriptionDiscount struct {
	ApplyOn       ApplyOn      `json:"apply_on,omitempty"`
	DurationType  DurationType `json:"duration_type"`
	Percentage    *float64     `json:"percentage,omitempty"`
	Amount        *int64       `json:"amount,omitempty"`
	Period        *int32       `json:"period,omitempty"`
	PeriodUnit    PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool        `json:"included_in_mrr,omitempty"`
	ItemPriceId   string       `json:"item_price_id,omitempty"`
	Quantity      *int32       `json:"quantity,omitempty"`
	Label         string       `json:"label,omitempty"`
}

// input sub resource params single
type PricingPageSessionCreateForNewSubscriptionBillingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}

// input sub resource params single
type PricingPageSessionCreateForNewSubscriptionShippingAddress struct {
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Email            string           `json:"email,omitempty"`
	Company          string           `json:"company,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Line1            string           `json:"line1,omitempty"`
	Line2            string           `json:"line2,omitempty"`
	Line3            string           `json:"line3,omitempty"`
	City             string           `json:"city,omitempty"`
	StateCode        string           `json:"state_code,omitempty"`
	State            string           `json:"state,omitempty"`
	Zip              string           `json:"zip,omitempty"`
	Country          string           `json:"country,omitempty"`
	ValidationStatus ValidationStatus `json:"validation_status,omitempty"`
}
type PricingPageSessionCreateForExistingSubscriptionRequest struct {
	RedirectUrl  string                                                       `json:"redirect_url,omitempty"`
	PricingPage  *PricingPageSessionCreateForExistingSubscriptionPricingPage  `json:"pricing_page,omitempty"`
	Subscription *PricingPageSessionCreateForExistingSubscriptionSubscription `json:"subscription,omitempty"`
	Discounts    []*PricingPageSessionCreateForExistingSubscriptionDiscount   `json:"discounts,omitempty"`
	apiRequest   `json:"-" form:"-"`
}

func (r *PricingPageSessionCreateForExistingSubscriptionRequest) payload() any { return r }

// input sub resource params single
type PricingPageSessionCreateForExistingSubscriptionPricingPage struct {
	Id string `json:"id"`
}

// input sub resource params single
type PricingPageSessionCreateForExistingSubscriptionSubscription struct {
	Id string `json:"id"`
}

// input sub resource params multi
type PricingPageSessionCreateForExistingSubscriptionDiscount struct {
	ApplyOn       ApplyOn      `json:"apply_on,omitempty"`
	DurationType  DurationType `json:"duration_type"`
	Percentage    *float64     `json:"percentage,omitempty"`
	Amount        *int64       `json:"amount,omitempty"`
	Period        *int32       `json:"period,omitempty"`
	PeriodUnit    PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool        `json:"included_in_mrr,omitempty"`
	ItemPriceId   string       `json:"item_price_id,omitempty"`
	Quantity      *int32       `json:"quantity,omitempty"`
	Label         string       `json:"label,omitempty"`
}

// operation response
type PricingPageSessionCreateForNewSubscriptionResponse struct {
	PricingPageSession *PricingPageSession `json:"pricing_page_session,omitempty"`
	apiResponse
}

// operation response
type PricingPageSessionCreateForExistingSubscriptionResponse struct {
	PricingPageSession *PricingPageSession `json:"pricing_page_session,omitempty"`
	apiResponse
}
