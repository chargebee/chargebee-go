package comment

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	commentEnum "github.com/chargebee/chargebee-go/models/comment/enum"
)

type Comment struct {
	Id         string           `json:"id"`
	EntityType enum.EntityType  `json:"entity_type"`
	AddedBy    string           `json:"added_by"`
	Notes      string           `json:"notes"`
	CreatedAt  int64            `json:"created_at"`
	Type       commentEnum.Type `json:"type"`
	EntityId   string           `json:"entity_id"`
	Object     string           `json:"object"`
}
type CreateRequestParams struct {
	EntityType enum.EntityType `json:"entity_type"`
	EntityId   string          `json:"entity_id"`
	Notes      string          `json:"notes"`
	AddedBy    string          `json:"added_by,omitempty"`
}
type ListRequestParams struct {
	Limit      *int32                  `json:"limit,omitempty"`
	Offset     string                  `json:"offset,omitempty"`
	EntityType enum.EntityType         `json:"entity_type,omitempty"`
	EntityId   string                  `json:"entity_id,omitempty"`
	CreatedAt  *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy     *filter.SortFilter      `json:"sort_by,omitempty"`
}
