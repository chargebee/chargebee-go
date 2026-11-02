package customdata

import (
	"github.com/chargebee/chargebee-go/v3/enum"
)

type CustomData struct {
	EntityType   enum.EntityType `json:"entity_type"`
	Id           string          `json:"id"`
	SchemaId     string          `json:"schema_id"`
	SchemaValues string          `json:"schema_values"`
	CreatedAt    int64           `json:"created_at"`
	ModifiedAt   int64           `json:"modified_at"`
	UpdatedAt    int64           `json:"updated_at"`
	Object       string          `json:"object"`
}
