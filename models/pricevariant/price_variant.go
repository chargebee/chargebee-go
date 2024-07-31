package pricevariant

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	priceVariantEnum "github.com/chargebee/chargebee-go/v3/models/pricevariant/enum"
)

type PriceVariant struct {
	Id              string                  `json:"id"`
	Name            string                  `json:"name"`
	ExternalName    string                  `json:"external_name"`
	VariantGroup    string                  `json:"variant_group"`
	Description     string                  `json:"description"`
	Status          priceVariantEnum.Status `json:"status"`
	CreatedAt       int64                   `json:"created_at"`
	ResourceVersion int64                   `json:"resource_version"`
	UpdatedAt       int64                   `json:"updated_at"`
	ArchivedAt      int64                   `json:"archived_at"`
	Attributes      []*Attribute            `json:"attributes"`
	Object          string                  `json:"object"`
}
type Attribute struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Object string `json:"object"`
}
type CreateRequestParams struct {
	Id           string                   `json:"id"`
	Name         string                   `json:"name"`
	ExternalName string                   `json:"external_name,omitempty"`
	Description  string                   `json:"description,omitempty"`
	VariantGroup string                   `json:"variant_group,omitempty"`
	Attributes   []*CreateAttributeParams `json:"attributes,omitempty"`
}
type CreateAttributeParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type UpdateRequestParams struct {
	Name         string                   `json:"name,omitempty"`
	ExternalName string                   `json:"external_name,omitempty"`
	Description  string                   `json:"description,omitempty"`
	VariantGroup string                   `json:"variant_group,omitempty"`
	Status       priceVariantEnum.Status  `json:"status,omitempty"`
	Attributes   []*UpdateAttributeParams `json:"attributes,omitempty"`
}
type UpdateAttributeParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type ListRequestParams struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Offset    string                  `json:"offset,omitempty"`
	Id        *filter.StringFilter    `json:"id,omitempty"`
	Name      *filter.StringFilter    `json:"name,omitempty"`
	Status    *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy    *filter.SortFilter      `json:"sort_by,omitempty"`
}
