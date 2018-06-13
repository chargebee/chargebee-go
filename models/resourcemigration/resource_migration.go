package resourcemigration

import (
	"github.com/chargebee/chargebee-go/enum"
	resourceMigrationEnum "github.com/chargebee/chargebee-go/models/resourcemigration/enum"
)

type ResourceMigration struct {
	FromSite   string                       `json:"from_site"`
	EntityType enum.EntityType              `json:"entity_type"`
	EntityId   string                       `json:"entity_id"`
	Status     resourceMigrationEnum.Status `json:"status"`
	Errors     string                       `json:"errors"`
	CreatedAt  int64                        `json:"created_at"`
	UpdatedAt  int64                        `json:"updated_at"`
	Object     string                       `json:"object"`
}
type RetrieveLatestRequestParams struct {
	FromSite   string          `json:"from_site"`
	EntityType enum.EntityType `json:"entity_type"`
	EntityId   string          `json:"entity_id"`
}
