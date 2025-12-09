package chargebee

import (
	"fmt"
	"net/url"
)

type RuleService struct {
	transport *Transport
}

func (s *RuleService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/rules/%v", url.PathEscape(id)), nil)
}
