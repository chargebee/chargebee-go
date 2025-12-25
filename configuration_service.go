package chargebee

import (
	"fmt"
)

type ConfigurationService struct {
	config *ClientConfig
}

func (s *ConfigurationService) List() (*ConfigurationListResponse, error) {
	req := &BlankRequest{}
	req.method = "GET"
	req.path = fmt.Sprintf("/configurations")
	return send[*ConfigurationListResponse](req, s.config)
}
