package chargebee

import (
	"fmt"
)

type ResourceMigrationService struct {
	transport *Transport
}

func (s *ResourceMigrationService) RetrieveLatest(req *RetrieveLatestRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/resource_migrations/retrieve_latest"), req)
}
