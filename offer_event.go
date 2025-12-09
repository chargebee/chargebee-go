package chargebee

type Type string

const (
	TypeViewed    Type = "viewed"
	TypeDismissed Type = "dismissed"
)

type OfferEvent struct {
	Object string `json:"object"`
}
type OfferEventsRequest struct {
	PersonalizedOfferId string `json:"personalized_offer_id"`
	Type                Type   `json:"type"`
}

type OfferEventsResponse struct {
}
