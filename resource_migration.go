package chargebee

type ResourceMigrationEntityType string

const (
	ResourceMigrationEntityTypeCustomer ResourceMigrationEntityType = "customer"
)

type ResourceMigrationStatus string

const (
	ResourceMigrationStatusScheduled ResourceMigrationStatus = "scheduled"
	ResourceMigrationStatusFailed    ResourceMigrationStatus = "failed"
	ResourceMigrationStatusSucceeded ResourceMigrationStatus = "succeeded"
)

// just struct
type ResourceMigration struct {
	FromSite   string                      `json:"from_site"`
	EntityType ResourceMigrationEntityType `json:"entity_type"`
	EntityId   string                      `json:"entity_id"`
	Status     ResourceMigrationStatus     `json:"status"`
	Errors     string                      `json:"errors"`
	CreatedAt  int64                       `json:"created_at"`
	UpdatedAt  int64                       `json:"updated_at"`
	Object     string                      `json:"object"`
}

// sub resources
// operations
// input params
type ResourceMigrationRetrieveLatestRequest struct {
	FromSite   string                      `json:"from_site"`
	EntityType ResourceMigrationEntityType `json:"entity_type"`
	EntityId   string                      `json:"entity_id"`
	apiRequest `json:"-" form:"-"`
}

func (r *ResourceMigrationRetrieveLatestRequest) payload() any { return r }

// operation response
type ResourceMigrationRetrieveLatestResponse struct {
	ResourceMigration *ResourceMigration `json:"resource_migration,omitempty"`
	apiResponse
}
