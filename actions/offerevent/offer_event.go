package offerevent

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/offerevent"
)

func OfferEvents(params *offerevent.OfferEventsRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/offer_events"), params).SetSubDomain("grow")
}
