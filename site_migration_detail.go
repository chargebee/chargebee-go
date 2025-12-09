package chargebee

type Status string

const (
	StatusMovedIn   Status = "moved_in"
	StatusMovedOut  Status = "moved_out"
	StatusMovingOut Status = "moving_out"
)

type SiteMigrationDetail struct {
	EntityId            string          `json:"entity_id"`
	OtherSiteName       string          `json:"other_site_name"`
	EntityIdAtOtherSite string          `json:"entity_id_at_other_site"`
	MigratedAt          int64           `json:"migrated_at"`
	EntityType          enum.EntityType `json:"entity_type"`
	Status              Status          `json:"status"`
	Object              string          `json:"object"`
}
type ListRequest struct {
	Limit               *int32               `json:"limit,omitempty"`
	Offset              string               `json:"offset,omitempty"`
	EntityIdAtOtherSite *filter.StringFilter `json:"entity_id_at_other_site,omitempty"`
	OtherSiteName       *filter.StringFilter `json:"other_site_name,omitempty"`
	EntityId            *filter.StringFilter `json:"entity_id,omitempty"`
	EntityType          *filter.EnumFilter   `json:"entity_type,omitempty"`
	Status              *filter.EnumFilter   `json:"status,omitempty"`
}

type ListSiteMigrationDetailResponse struct {
	SiteMigrationDetail *SiteMigrationDetail `json:"site_migration_detail,omitempty"`
}

type ListResponse struct {
	List       []*ListSiteMigrationDetailResponse `json:"list,omitempty"`
	NextOffset string                             `json:"next_offset,omitempty"`
}
