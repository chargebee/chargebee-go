package chargebee

import (
	"fmt"
	"net/url"
)

type OmnichannelSubscriptionService struct {
	config *ClientConfig
}

func (s *OmnichannelSubscriptionService) Retrieve(id string) (*OmnichannelSubscriptionRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/omnichannel_subscriptions/%v", url.PathEscape(id))
	return send[*OmnichannelSubscriptionRetrieveResponse](req, s.config)
}

func (s *OmnichannelSubscriptionService) List(req *OmnichannelSubscriptionListRequest) (*OmnichannelSubscriptionListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/omnichannel_subscriptions")
	req.isListRequest = true
	return send[*OmnichannelSubscriptionListResponse](req, s.config)
}

func (s *OmnichannelSubscriptionService) OmnichannelTransactionsForOmnichannelSubscription(id string, req *OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionRequest) (*OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/omnichannel_subscriptions/%v/omnichannel_transactions", url.PathEscape(id))
	req.isListRequest = true
	return send[*OmnichannelSubscriptionOmnichannelTransactionsForOmnichannelSubscriptionResponse](req, s.config)
}

func (s *OmnichannelSubscriptionService) Move(id string, req *OmnichannelSubscriptionMoveRequest) (*OmnichannelSubscriptionMoveResponse, error) {
	req.method = "POST"
	req.path = fmt.Sprintf("/omnichannel_subscriptions/%v/move", url.PathEscape(id))
	req.isIdempotent = true
	return send[*OmnichannelSubscriptionMoveResponse](req, s.config)
}
