package sitemigrationdetail

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	siteMigrationDetailEnum "github.com/chargebee/chargebee-go/models/sitemigrationdetail/enum"
)

type SiteMigrationDetail struct {
	EntityId            string                         `json:"entity_id"`
	OtherSiteName       string                         `json:"other_site_name"`
	EntityIdAtOtherSite string                         `json:"entity_id_at_other_site"`
	MigratedAt          int64                          `json:"migrated_at"`
	EntityType          enum.EntityType                `json:"entity_type"`
	Status              siteMigrationDetailEnum.Status `json:"status"`
	Object              string                         `json:"object"`
}
type ListRequestParams struct {
	Limit               *int32               `json:"limit,omitempty"`
	Offset              string               `json:"offset,omitempty"`
	EntityIdAtOtherSite *filter.StringFilter `json:"entity_id_at_other_site,omitempty"`
	OtherSiteName       *filter.StringFilter `json:"other_site_name,omitempty"`
	EntityId            *filter.StringFilter `json:"entity_id,omitempty"`
	EntityType          *filter.EnumFilter   `json:"entity_type,omitempty"`
	Status              *filter.EnumFilter   `json:"status,omitempty"`
}
