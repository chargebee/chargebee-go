package chargebee

import (
	"fmt"
)

type SiteMigrationDetailService struct {
	transport *Transport
}

func (s *SiteMigrationDetailService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/site_migration_details"), req)
}
