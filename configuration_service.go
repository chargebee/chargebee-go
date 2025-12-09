package chargebee

import (
	"fmt"
)

type ConfigurationService struct {
	transport *Transport
}

func (s *ConfigurationService) List() Request {
	return s.transport.Send("GET", fmt.Sprintf("/configurations"), nil)
}
