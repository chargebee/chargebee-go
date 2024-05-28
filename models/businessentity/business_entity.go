package businessentity

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	businessEntityEnum "github.com/chargebee/chargebee-go/v3/models/businessentity/enum"
)

type BusinessEntity struct {
	Id              string                    `json:"id"`
	Name            string                    `json:"name"`
	Status          businessEntityEnum.Status `json:"status"`
	Deleted         bool                      `json:"deleted"`
	CreatedAt       int64                     `json:"created_at"`
	ResourceVersion int64                     `json:"resource_version"`
	UpdatedAt       int64                     `json:"updated_at"`
	Object          string                    `json:"object"`
}
type CreateTransfersRequestParams struct {
	ActiveResourceIds            []string `json:"active_resource_ids"`
	DestinationBusinessEntityIds []string `json:"destination_business_entity_ids"`
	SourceBusinessEntityIds      []string `json:"source_business_entity_ids,omitempty"`
	ResourceTypes                []string `json:"resource_types"`
	ReasonCodes                  []string `json:"reason_codes"`
}
type GetTransfersRequestParams struct {
	Limit            *int32                  `json:"limit,omitempty"`
	Offset           string                  `json:"offset,omitempty"`
	ResourceType     *filter.StringFilter    `json:"resource_type,omitempty"`
	ResourceId       *filter.StringFilter    `json:"resource_id,omitempty"`
	ActiveResourceId *filter.StringFilter    `json:"active_resource_id,omitempty"`
	CreatedAt        *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy           *filter.SortFilter      `json:"sort_by,omitempty"`
}
