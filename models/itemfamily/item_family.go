package itemfamily

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
	itemFamilyEnum "github.com/chargebee/chargebee-go/v3/models/itemfamily/enum"
)

type ItemFamily struct {
	Id               string                 `json:"id"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Status           itemFamilyEnum.Status  `json:"status"`
	ResourceVersion  int64                  `json:"resource_version"`
	UpdatedAt        int64                  `json:"updated_at"`
	Channel          enum.Channel           `json:"channel"`
	BusinessEntityId string                 `json:"business_entity_id"`
	Deleted          bool                   `json:"deleted"`
	CustomField      map[string]interface{} `json:"custom_field"`
	Object           string                 `json:"object"`
}
type CreateRequestParams struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	BusinessEntityId string `json:"business_entity_id,omitempty"`
}
type ListRequestParams struct {
	Limit                     *int32                  `json:"limit,omitempty"`
	Offset                    string                  `json:"offset,omitempty"`
	Id                        *filter.StringFilter    `json:"id,omitempty"`
	Name                      *filter.StringFilter    `json:"name,omitempty"`
	UpdatedAt                 *filter.TimestampFilter `json:"updated_at,omitempty"`
	BusinessEntityId          *filter.StringFilter    `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter   `json:"include_site_level_resources,omitempty"`
}
type UpdateRequestParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
