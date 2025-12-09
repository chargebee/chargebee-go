package chargebee

type Status string

const (
	StatusActive  Status = "active"
	StatusDeleted Status = "deleted"
)

type ItemFamily struct {
	Id               string                 `json:"id"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Status           Status                 `json:"status"`
	ResourceVersion  int64                  `json:"resource_version"`
	UpdatedAt        int64                  `json:"updated_at"`
	Channel          enum.Channel           `json:"channel"`
	BusinessEntityId string                 `json:"business_entity_id"`
	Deleted          bool                   `json:"deleted"`
	CustomField      map[string]interface{} `json:"custom_field"`
	Object           string                 `json:"object"`
}
type CreateRequest struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	BusinessEntityId string `json:"business_entity_id,omitempty"`
}
type ListRequest struct {
	Limit                     *int32                  `json:"limit,omitempty"`
	Offset                    string                  `json:"offset,omitempty"`
	Id                        *filter.StringFilter    `json:"id,omitempty"`
	Name                      *filter.StringFilter    `json:"name,omitempty"`
	UpdatedAt                 *filter.TimestampFilter `json:"updated_at,omitempty"`
	BusinessEntityId          *filter.StringFilter    `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter   `json:"include_site_level_resources,omitempty"`
}
type UpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateResponse struct {
	ItemFamily *ItemFamily `json:"item_family,omitempty"`
}

type RetrieveResponse struct {
	ItemFamily *ItemFamily `json:"item_family,omitempty"`
}

type ListItemFamilyResponse struct {
	ItemFamily *ItemFamily `json:"item_family,omitempty"`
}

type ListResponse struct {
	List       []*ListItemFamilyResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
}

type UpdateResponse struct {
	ItemFamily *ItemFamily `json:"item_family,omitempty"`
}

type DeleteResponse struct {
	ItemFamily *ItemFamily `json:"item_family,omitempty"`
}
