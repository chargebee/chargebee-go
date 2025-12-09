package chargebee

type ProcessingType string

const (
	ProcessingTypeBillingUpdate ProcessingType = "billing_update"
	ProcessingTypeCheckout      ProcessingType = "checkout"
	ProcessingTypeUrlRedirect   ProcessingType = "url_redirect"
	ProcessingTypeWebhook       ProcessingType = "webhook"
	ProcessingTypeEmail         ProcessingType = "email"
)

type Status string

const (
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

type ErrorCode string

const (
	ErrorCodeBillingUpdateFailed       ErrorCode = "billing_update_failed"
	ErrorCodeCheckoutAbandoned         ErrorCode = "checkout_abandoned"
	ErrorCodeExternalFulfillmentFailed ErrorCode = "external_fulfillment_failed"
	ErrorCodeInternalError             ErrorCode = "internal_error"
	ErrorCodeFulfillmentExpired        ErrorCode = "fulfillment_expired"
)

type OfferFulfillment struct {
	Id                  string         `json:"id"`
	PersonalizedOfferId string         `json:"personalized_offer_id"`
	OptionId            string         `json:"option_id"`
	ProcessingType      ProcessingType `json:"processing_type"`
	Status              Status         `json:"status"`
	RedirectUrl         string         `json:"redirect_url"`
	FailedAt            int64          `json:"failed_at"`
	CreatedAt           int64          `json:"created_at"`
	CompletedAt         int64          `json:"completed_at"`
	Error               *Error         `json:"error"`
	Object              string         `json:"object"`
}
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Object  string    `json:"object"`
}
type OfferFulfillmentsRequest struct {
	PersonalizedOfferId string `json:"personalized_offer_id"`
	OptionId            string `json:"option_id"`
}
type OfferFulfillmentsUpdateRequest struct {
	Id            string `json:"id"`
	Status        Status `json:"status"`
	FailureReason string `json:"failure_reason,omitempty"`
}

type OfferFulfillmentsResponse struct {
	OfferFulfillment *OfferFulfillment      `json:"offer_fulfillment,omitempty"`
	HostedPage       *hostedpage.HostedPage `json:"hosted_page,omitempty"`
}

type OfferFulfillmentsGetResponse struct {
	OfferFulfillment *OfferFulfillment `json:"offer_fulfillment,omitempty"`
}

type OfferFulfillmentsUpdateResponse struct {
	OfferFulfillment *OfferFulfillment `json:"offer_fulfillment,omitempty"`
}
