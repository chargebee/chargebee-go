package itemfamily

import (
	"github.com/chargebee/chargebee-go/filter"
	itemFamilyEnum "github.com/chargebee/chargebee-go/models/itemfamily/enum"
)

type ItemFamily struct {
	Id              string                 `json:"id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Status          itemFamilyEnum.Status  `json:"status"`
	ResourceVersion int64                  `json:"resource_version"`
	UpdatedAt       int64                  `json:"updated_at"`
	CustomField     map[string]interface{} `json:"custom_field"`
	Object          string                 `json:"object"`
}
type CreateRequestParams struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
type ListRequestParams struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Offset    string                  `json:"offset,omitempty"`
	Id        *filter.StringFilter    `json:"id,omitempty"`
	Name      *filter.StringFilter    `json:"name,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type UpdateRequestParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
