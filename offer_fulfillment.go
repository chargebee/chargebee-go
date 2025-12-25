package chargebee

type OfferFulfillmentProcessingType string

const (
	OfferFulfillmentProcessingTypeBillingUpdate OfferFulfillmentProcessingType = "billing_update"
	OfferFulfillmentProcessingTypeCheckout      OfferFulfillmentProcessingType = "checkout"
	OfferFulfillmentProcessingTypeUrlRedirect   OfferFulfillmentProcessingType = "url_redirect"
	OfferFulfillmentProcessingTypeWebhook       OfferFulfillmentProcessingType = "webhook"
	OfferFulfillmentProcessingTypeEmail         OfferFulfillmentProcessingType = "email"
)

type OfferFulfillmentStatus string

const (
	OfferFulfillmentStatusInProgress OfferFulfillmentStatus = "in_progress"
	OfferFulfillmentStatusCompleted  OfferFulfillmentStatus = "completed"
	OfferFulfillmentStatusFailed     OfferFulfillmentStatus = "failed"
)

type OfferFulfillmentErrorCode string

const (
	OfferFulfillmentErrorCodeBillingUpdateFailed       OfferFulfillmentErrorCode = "billing_update_failed"
	OfferFulfillmentErrorCodeCheckoutAbandoned         OfferFulfillmentErrorCode = "checkout_abandoned"
	OfferFulfillmentErrorCodeExternalFulfillmentFailed OfferFulfillmentErrorCode = "external_fulfillment_failed"
	OfferFulfillmentErrorCodeInternalError             OfferFulfillmentErrorCode = "internal_error"
	OfferFulfillmentErrorCodeFulfillmentExpired        OfferFulfillmentErrorCode = "fulfillment_expired"
)

// just struct
type OfferFulfillment struct {
	Id                  string                         `json:"id"`
	PersonalizedOfferId string                         `json:"personalized_offer_id"`
	OptionId            string                         `json:"option_id"`
	ProcessingType      OfferFulfillmentProcessingType `json:"processing_type"`
	Status              OfferFulfillmentStatus         `json:"status"`
	RedirectUrl         string                         `json:"redirect_url"`
	FailedAt            int64                          `json:"failed_at"`
	CreatedAt           int64                          `json:"created_at"`
	CompletedAt         int64                          `json:"completed_at"`
	Error               *Error                         `json:"error"`
	Object              string                         `json:"object"`
}

// sub resources
type OfferFulfillmentError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Object  string    `json:"object"`
}

// operations
// input params
type OfferFulfillmentOfferFulfillmentsRequest struct {
	PersonalizedOfferId string `json:"personalized_offer_id"`
	OptionId            string `json:"option_id"`
	apiRequest          `json:"-" form:"-"`
}

func (r *OfferFulfillmentOfferFulfillmentsRequest) payload() any { return r }

type OfferFulfillmentOfferFulfillmentsUpdateRequest struct {
	Id            string                 `json:"id"`
	Status        OfferFulfillmentStatus `json:"status"`
	FailureReason string                 `json:"failure_reason,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *OfferFulfillmentOfferFulfillmentsUpdateRequest) payload() any { return r }

// operation response
type OfferFulfillmentOfferFulfillmentsResponse struct {
	OfferFulfillment *OfferFulfillment `json:"offer_fulfillment,omitempty"`
	HostedPage       HostedPage        `json:"hosted_page,omitempty"`
	apiResponse
}

// operation response
type OfferFulfillmentOfferFulfillmentsGetResponse struct {
	OfferFulfillment *OfferFulfillment `json:"offer_fulfillment,omitempty"`
	apiResponse
}

// operation response
type OfferFulfillmentOfferFulfillmentsUpdateResponse struct {
	OfferFulfillment *OfferFulfillment `json:"offer_fulfillment,omitempty"`
	apiResponse
}
