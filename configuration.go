package chargebee

// just struct
type Configuration struct {
	Domain                      string                      `json:"domain"`
	ProductCatalogVersion       ProductCatalogVersion       `json:"product_catalog_version"`
	ChargebeeResponseSchemaType ChargebeeResponseSchemaType `json:"chargebee_response_schema_type"`
	Object                      string                      `json:"object"`
}

// sub resources
// operations
// input params

// operation response
type ConfigurationListResponse struct {
	Configurations []*Configuration `json:"configurations,omitempty"`
	apiResponse
}
