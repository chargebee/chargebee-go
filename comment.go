package chargebee

type Type string

const (
	TypeUser   Type = "user"
	TypeSystem Type = "system"
)

type Comment struct {
	Id               string          `json:"id"`
	EntityType       enum.EntityType `json:"entity_type"`
	AddedBy          string          `json:"added_by"`
	Notes            string          `json:"notes"`
	CreatedAt        int64           `json:"created_at"`
	Type             Type            `json:"type"`
	EntityId         string          `json:"entity_id"`
	BusinessEntityId string          `json:"business_entity_id"`
	Object           string          `json:"object"`
}
type CreateRequest struct {
	EntityType enum.EntityType `json:"entity_type"`
	EntityId   string          `json:"entity_id"`
	Notes      string          `json:"notes"`
	AddedBy    string          `json:"added_by,omitempty"`
}
type ListRequest struct {
	Limit      *int32                  `json:"limit,omitempty"`
	Offset     string                  `json:"offset,omitempty"`
	EntityType enum.EntityType         `json:"entity_type,omitempty"`
	EntityId   string                  `json:"entity_id,omitempty"`
	CreatedAt  *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy     *filter.SortFilter      `json:"sort_by,omitempty"`
}

type CreateResponse struct {
	Comment *Comment `json:"comment,omitempty"`
}

type RetrieveResponse struct {
	Comment *Comment `json:"comment,omitempty"`
}

type ListCommentResponse struct {
	Comment *Comment `json:"comment,omitempty"`
}

type ListResponse struct {
	List       []*ListCommentResponse `json:"list,omitempty"`
	NextOffset string                 `json:"next_offset,omitempty"`
}

type DeleteResponse struct {
	Comment *Comment `json:"comment,omitempty"`
}
