package customdataschema

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	customDataSchemaEnum "github.com/chargebee/chargebee-go/v3/models/customdataschema/enum"
)

type CustomDataSchema struct {
	Id               string                      `json:"id"`
	DisplayName      string                      `json:"display_name"`
	EntityType       enum.EntityType             `json:"entity_type"`
	SchemaDefinition string                      `json:"schema_definition"`
	Status           customDataSchemaEnum.Status `json:"status"`
	CreatedAt        int64                       `json:"created_at"`
	ModifiedAt       int64                       `json:"modified_at"`
	UpdatedAt        int64                       `json:"updated_at"`
	Object           string                      `json:"object"`
}
