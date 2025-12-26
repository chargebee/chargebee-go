package chargebee

type SiteMigrationDetailStatus string

const (
	SiteMigrationDetailStatusMovedIn   SiteMigrationDetailStatus = "moved_in"
	SiteMigrationDetailStatusMovedOut  SiteMigrationDetailStatus = "moved_out"
	SiteMigrationDetailStatusMovingOut SiteMigrationDetailStatus = "moving_out"
)

// just struct
type SiteMigrationDetail struct {
	EntityId            string                    `json:"entity_id"`
	OtherSiteName       string                    `json:"other_site_name"`
	EntityIdAtOtherSite string                    `json:"entity_id_at_other_site"`
	MigratedAt          int64                     `json:"migrated_at"`
	EntityType          EntityType                `json:"entity_type"`
	Status              SiteMigrationDetailStatus `json:"status"`
	Object              string                    `json:"object"`
}

// sub resources
// operations
// input params
type SiteMigrationDetailListRequest struct {
	Limit               *int32        `json:"limit,omitempty"`
	Offset              string        `json:"offset,omitempty"`
	EntityIdAtOtherSite *StringFilter `json:"entity_id_at_other_site,omitempty"`
	OtherSiteName       *StringFilter `json:"other_site_name,omitempty"`
	EntityId            *StringFilter `json:"entity_id,omitempty"`
	EntityType          *EnumFilter   `json:"entity_type,omitempty"`
	Status              *EnumFilter   `json:"status,omitempty"`
	apiRequest          `json:"-" form:"-"`
}

func (r *SiteMigrationDetailListRequest) payload() any { return r }

// operation sub response
type SiteMigrationDetailListSiteMigrationDetailResponse struct {
	SiteMigrationDetail *SiteMigrationDetail `json:"site_migration_detail,omitempty"`
}

// operation response
type SiteMigrationDetailListResponse struct {
	List       []*SiteMigrationDetailListSiteMigrationDetailResponse `json:"list,omitempty"`
	NextOffset string                                                `json:"next_offset,omitempty"`
	apiResponse
}
