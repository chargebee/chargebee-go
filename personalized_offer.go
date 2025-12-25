package chargebee

type PersonalizedOfferOptionProcessingType string

const (
	PersonalizedOfferOptionProcessingTypeBillingUpdate PersonalizedOfferOptionProcessingType = "billing_update"
	PersonalizedOfferOptionProcessingTypeCheckout      PersonalizedOfferOptionProcessingType = "checkout"
	PersonalizedOfferOptionProcessingTypeUrlRedirect   PersonalizedOfferOptionProcessingType = "url_redirect"
	PersonalizedOfferOptionProcessingTypeWebhook       PersonalizedOfferOptionProcessingType = "webhook"
	PersonalizedOfferOptionProcessingTypeEmail         PersonalizedOfferOptionProcessingType = "email"
)

type PersonalizedOfferOptionProcessingLayout string

const (
	PersonalizedOfferOptionProcessingLayoutInApp    PersonalizedOfferOptionProcessingLayout = "in_app"
	PersonalizedOfferOptionProcessingLayoutFullPage PersonalizedOfferOptionProcessingLayout = "full_page"
)

// just struct
type PersonalizedOffer struct {
	Id      string                     `json:"id"`
	OfferId string                     `json:"offer_id"`
	Content *Content                   `json:"content"`
	Options []*PersonalizedOfferOption `json:"options"`
	Object  string                     `json:"object"`
}

// sub resources
type PersonalizedOfferContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type PersonalizedOfferOption struct {
	Id               string                 `json:"id"`
	Label            string                 `json:"label"`
	ProcessingType   OptionProcessingType   `json:"processing_type"`
	ProcessingLayout OptionProcessingLayout `json:"processing_layout"`
	RedirectUrl      string                 `json:"redirect_url"`
	Object           string                 `json:"object"`
}

// operations
// input params
type PersonalizedOfferPersonalizedOffersRequest struct {
	FirstName      string                                             `json:"first_name,omitempty"`
	LastName       string                                             `json:"last_name,omitempty"`
	Email          string                                             `json:"email,omitempty"`
	Roles          []string                                           `json:"roles,omitempty"`
	ExternalUserId string                                             `json:"external_user_id,omitempty"`
	SubscriptionId string                                             `json:"subscription_id,omitempty"`
	CustomerId     string                                             `json:"customer_id"`
	Custom         map[string]interface{}                             `json:"custom,omitempty"`
	RequestContext *PersonalizedOfferPersonalizedOffersRequestContext `json:"request_context,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *PersonalizedOfferPersonalizedOffersRequest) payload() any { return r }

// input sub resource params single
type PersonalizedOfferPersonalizedOffersRequestContext struct {
	UserAgent   string `json:"user_agent,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	Url         string `json:"url,omitempty"`
	ReferrerUrl string `json:"referrer_url,omitempty"`
}

// operation response
type PersonalizedOfferPersonalizedOffersResponse struct {
	PersonalizedOffers []*PersonalizedOffer `json:"personalized_offers,omitempty"`
	ExpiresAt          int64                `json:"expires_at,omitempty"`
	Brand              Brand                `json:"brand,omitempty"`
	apiResponse
}
