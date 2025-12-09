package chargebee

import (
	"fmt"
	"net/url"
)

type OmnichannelSubscriptionService struct {
	transport *Transport
}

func (s *OmnichannelSubscriptionService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/omnichannel_subscriptions/%v", url.PathEscape(id)), nil)
}

func (s *OmnichannelSubscriptionService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/omnichannel_subscriptions"), req)
}

func (s *OmnichannelSubscriptionService) OmnichannelTransactionsForOmnichannelSubscription(id string, req *OmnichannelTransactionsForOmnichannelSubscriptionRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/omnichannel_subscriptions/%v/omnichannel_transactions", url.PathEscape(id)), req)
}

func (s *OmnichannelSubscriptionService) Move(id string, req *MoveRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/omnichannel_subscriptions/%v/move", url.PathEscape(id)), req).SetIdempotency(true)
}
