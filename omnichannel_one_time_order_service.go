package chargebee

import (
	"fmt"
	"net/url"
)

type OmnichannelOneTimeOrderService struct {
	transport *Transport
}

func (s *OmnichannelOneTimeOrderService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/omnichannel_one_time_orders/%v", url.PathEscape(id)), nil)
}

func (s *OmnichannelOneTimeOrderService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/omnichannel_one_time_orders"), req)
}
