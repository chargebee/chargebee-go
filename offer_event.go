package chargebee

type OfferEventType string

const (
	OfferEventTypeViewed    OfferEventType = "viewed"
	OfferEventTypeDismissed OfferEventType = "dismissed"
)

// just struct
type OfferEvent struct {
	Object string `json:"object"`
}

// sub resources
// operations
// input params
type OfferEventOfferEventsRequest struct {
	PersonalizedOfferId string         `json:"personalized_offer_id"`
	Type                OfferEventType `json:"type"`
	apiRequest          `json:"-" form:"-"`
}

func (r *OfferEventOfferEventsRequest) payload() any { return r }

// operation response
type OfferEventOfferEventsResponse struct {
	apiResponse
}
