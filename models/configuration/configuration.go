package configuration

import (
	"github.com/chargebee/chargebee-go/v3/enum"
)

type Configuration struct {
	Domain                      string                           `json:"domain"`
	ProductCatalogVersion       enum.ProductCatalogVersion       `json:"product_catalog_version"`
	ChargebeeResponseSchemaType enum.ChargebeeResponseSchemaType `json:"chargebee_response_schema_type"`
	Object                      string                           `json:"object"`
}
