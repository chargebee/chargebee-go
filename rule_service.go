package chargebee

import (
	"fmt"
	"net/url"
)

type RuleService struct {
	config *ClientConfig
}

func (s *RuleService) Retrieve(id string) (*RuleRetrieveResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/rules/%v", url.PathEscape(id))
	return send[*RuleRetrieveResponse](req, s.config)
}
