package chargebee

import (
	"fmt"
)

type ResourceMigrationService struct {
	config *ClientConfig
}

func (s *ResourceMigrationService) RetrieveLatest(req *ResourceMigrationRetrieveLatestRequest) (*ResourceMigrationRetrieveLatestResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/resource_migrations/retrieve_latest")
	return send[*ResourceMigrationRetrieveLatestResponse](req, s.config)
}
