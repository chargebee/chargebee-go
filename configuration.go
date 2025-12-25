package chargebee

type ConfigurationProductCatalogVersion string

const (
	ConfigurationProductCatalogVersionV1 ConfigurationProductCatalogVersion = "v1"
	ConfigurationProductCatalogVersionV2 ConfigurationProductCatalogVersion = "v2"
)

type ConfigurationChargebeeResponseSchemaType string

const (
	ConfigurationChargebeeResponseSchemaTypePlansAddons ConfigurationChargebeeResponseSchemaType = "plans_addons"
	ConfigurationChargebeeResponseSchemaTypeItems       ConfigurationChargebeeResponseSchemaType = "items"
	ConfigurationChargebeeResponseSchemaTypeCompat      ConfigurationChargebeeResponseSchemaType = "compat"
)

// just struct
type Configuration struct {
	Domain                      string                                   `json:"domain"`
	ProductCatalogVersion       ConfigurationProductCatalogVersion       `json:"product_catalog_version"`
	ChargebeeResponseSchemaType ConfigurationChargebeeResponseSchemaType `json:"chargebee_response_schema_type"`
	Object                      string                                   `json:"object"`
}

// sub resources
// operations
// input params

// operation response
type ConfigurationListResponse struct {
	Configurations []*Configuration `json:"configurations,omitempty"`
	apiResponse
}
