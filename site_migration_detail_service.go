package chargebee

import (
	"fmt"
)

type SiteMigrationDetailService struct {
	config *ClientConfig
}

func (s *SiteMigrationDetailService) List(req *SiteMigrationDetailListRequest) (*SiteMigrationDetailListResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/site_migration_details")
	req.isListRequest = true
	return send[*SiteMigrationDetailListResponse](req, s.config)
}
