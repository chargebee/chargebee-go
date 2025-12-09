package chargebee

type Status string

const (
	StatusScheduled Status = "scheduled"
	StatusFailed    Status = "failed"
	StatusSucceeded Status = "succeeded"
)

type ResourceMigration struct {
	FromSite   string          `json:"from_site"`
	EntityType enum.EntityType `json:"entity_type"`
	EntityId   string          `json:"entity_id"`
	Status     Status          `json:"status"`
	Errors     string          `json:"errors"`
	CreatedAt  int64           `json:"created_at"`
	UpdatedAt  int64           `json:"updated_at"`
	Object     string          `json:"object"`
}
type RetrieveLatestRequest struct {
	FromSite   string          `json:"from_site"`
	EntityType enum.EntityType `json:"entity_type"`
	EntityId   string          `json:"entity_id"`
}

type RetrieveLatestResponse struct {
	ResourceMigration *ResourceMigration `json:"resource_migration,omitempty"`
}
