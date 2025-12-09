package chargebee

type OptionProcessingType string

const (
	OptionProcessingTypeBillingUpdate OptionProcessingType = "billing_update"
	OptionProcessingTypeCheckout      OptionProcessingType = "checkout"
	OptionProcessingTypeUrlRedirect   OptionProcessingType = "url_redirect"
	OptionProcessingTypeWebhook       OptionProcessingType = "webhook"
	OptionProcessingTypeEmail         OptionProcessingType = "email"
)

type OptionProcessingLayout string

const (
	OptionProcessingLayoutInApp    OptionProcessingLayout = "in_app"
	OptionProcessingLayoutFullPage OptionProcessingLayout = "full_page"
)

type PersonalizedOffer struct {
	Id      string    `json:"id"`
	OfferId string    `json:"offer_id"`
	Content *Content  `json:"content"`
	Options []*Option `json:"options"`
	Object  string    `json:"object"`
}
type Content struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type Option struct {
	Id               string                 `json:"id"`
	Label            string                 `json:"label"`
	ProcessingType   OptionProcessingType   `json:"processing_type"`
	ProcessingLayout OptionProcessingLayout `json:"processing_layout"`
	RedirectUrl      string                 `json:"redirect_url"`
	Object           string                 `json:"object"`
}
type PersonalizedOffersRequest struct {
	FirstName      string                            `json:"first_name,omitempty"`
	LastName       string                            `json:"last_name,omitempty"`
	Email          string                            `json:"email,omitempty"`
	Roles          []string                          `json:"roles,omitempty"`
	ExternalUserId string                            `json:"external_user_id,omitempty"`
	SubscriptionId string                            `json:"subscription_id,omitempty"`
	CustomerId     string                            `json:"customer_id"`
	Custom         map[string]interface{}            `json:"custom,omitempty"`
	RequestContext *PersonalizedOffersRequestContext `json:"request_context,omitempty"`
}
type PersonalizedOffersRequestContext struct {
	UserAgent   string `json:"user_agent,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	Url         string `json:"url,omitempty"`
	ReferrerUrl string `json:"referrer_url,omitempty"`
}

type PersonalizedOffersResponse struct {
	PersonalizedOffers []*PersonalizedOffer `json:"personalized_offers,omitempty"`
	ExpiresAt          int64                `json:"expires_at,omitempty"`
	Brand              *brand.Brand         `json:"brand,omitempty"`
}
