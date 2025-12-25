package chargebee

import (
	"fmt"
	"net/url"
)

type OmnichannelOneTimeOrderService struct {
	config *ClientConfig
}

func (s *OmnichannelOneTimeOrderService) Retrieve(id string) (*OmnichannelOneTimeOrderRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/omnichannel_one_time_orders/%v", url.PathEscape(id))
	return send[*OmnichannelOneTimeOrderRetrieveResponse](req, s.config)
}

func (s *OmnichannelOneTimeOrderService) List(req *OmnichannelOneTimeOrderListRequest) (*OmnichannelOneTimeOrderListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/omnichannel_one_time_orders")
	req.isListRequest = true
	return send[*OmnichannelOneTimeOrderListResponse](req, s.config)
}
