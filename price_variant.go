package chargebee

type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
	StatusDeleted  Status = "deleted"
)

type PriceVariant struct {
	Id               string       `json:"id"`
	Name             string       `json:"name"`
	ExternalName     string       `json:"external_name"`
	VariantGroup     string       `json:"variant_group"`
	Description      string       `json:"description"`
	Status           Status       `json:"status"`
	CreatedAt        int64        `json:"created_at"`
	ResourceVersion  int64        `json:"resource_version"`
	UpdatedAt        int64        `json:"updated_at"`
	ArchivedAt       int64        `json:"archived_at"`
	Attributes       []*Attribute `json:"attributes"`
	BusinessEntityId string       `json:"business_entity_id"`
	Deleted          bool         `json:"deleted"`
	Object           string       `json:"object"`
}
type Attribute struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Object string `json:"object"`
}
type CreateRequest struct {
	Id               string             `json:"id"`
	Name             string             `json:"name"`
	ExternalName     string             `json:"external_name,omitempty"`
	Description      string             `json:"description,omitempty"`
	VariantGroup     string             `json:"variant_group,omitempty"`
	BusinessEntityId string             `json:"business_entity_id,omitempty"`
	Attributes       []*CreateAttribute `json:"attributes,omitempty"`
}
type CreateAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type UpdateRequest struct {
	Name         string             `json:"name,omitempty"`
	ExternalName string             `json:"external_name,omitempty"`
	Description  string             `json:"description,omitempty"`
	VariantGroup string             `json:"variant_group,omitempty"`
	Status       Status             `json:"status,omitempty"`
	Attributes   []*UpdateAttribute `json:"attributes,omitempty"`
}
type UpdateAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type ListRequest struct {
	Limit                     *int32                  `json:"limit,omitempty"`
	Offset                    string                  `json:"offset,omitempty"`
	Id                        *filter.StringFilter    `json:"id,omitempty"`
	Name                      *filter.StringFilter    `json:"name,omitempty"`
	Status                    *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt                 *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt                 *filter.TimestampFilter `json:"created_at,omitempty"`
	BusinessEntityId          *filter.StringFilter    `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter   `json:"include_site_level_resources,omitempty"`
	SortBy                    *filter.SortFilter      `json:"sort_by,omitempty"`
}

type CreateResponse struct {
	PriceVariant *PriceVariant `json:"price_variant,omitempty"`
}

type RetrieveResponse struct {
	PriceVariant *PriceVariant `json:"price_variant,omitempty"`
}

type UpdateResponse struct {
	PriceVariant *PriceVariant `json:"price_variant,omitempty"`
}

type DeleteResponse struct {
	PriceVariant *PriceVariant `json:"price_variant,omitempty"`
}

type ListPriceVariantResponse struct {
	PriceVariant *PriceVariant `json:"price_variant,omitempty"`
}

type ListResponse struct {
	List       []*ListPriceVariantResponse `json:"list,omitempty"`
	NextOffset string                      `json:"next_offset,omitempty"`
}
