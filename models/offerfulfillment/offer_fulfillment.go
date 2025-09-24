package offerfulfillment

import (
	offerFulfillmentEnum "github.com/chargebee/chargebee-go/v3/models/offerfulfillment/enum"
)

type OfferFulfillment struct {
	Id                  string                              `json:"id"`
	PersonalizedOfferId string                              `json:"personalized_offer_id"`
	OptionId            string                              `json:"option_id"`
	ProcessingType      offerFulfillmentEnum.ProcessingType `json:"processing_type"`
	Status              offerFulfillmentEnum.Status         `json:"status"`
	RedirectUrl         string                              `json:"redirect_url"`
	FailedAt            int64                               `json:"failed_at"`
	CreatedAt           int64                               `json:"created_at"`
	CompletedAt         int64                               `json:"completed_at"`
	Error               *Error                              `json:"error"`
	Object              string                              `json:"object"`
}
type Error struct {
	Code    offerFulfillmentEnum.ErrorCode `json:"code"`
	Message string                         `json:"message"`
	Object  string                         `json:"object"`
}
type OfferFulfillmentsRequestParams struct {
	PersonalizedOfferId string `json:"personalized_offer_id"`
	OptionId            string `json:"option_id"`
}
type OfferFulfillmentsUpdateRequestParams struct {
	Id            string                      `json:"id"`
	Status        offerFulfillmentEnum.Status `json:"status"`
	FailureReason string                      `json:"failure_reason,omitempty"`
}
