package offerevent

import (
	offerEventEnum "github.com/chargebee/chargebee-go/v3/models/offerevent/enum"
)

type OfferEvent struct {
	Object string `json:"object"`
}
type OfferEventsRequestParams struct {
	PersonalizedOfferId string              `json:"personalized_offer_id"`
	Type                offerEventEnum.Type `json:"type"`
}
