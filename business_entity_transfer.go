package chargebee

type ResourceType string

const (
	ResourceTypeCustomer     ResourceType = "customer"
	ResourceTypeSubscription ResourceType = "subscription"
)

type ReasonCode string

const (
	ReasonCodeCorrection ReasonCode = "correction"
)

type BusinessEntityTransfer struct {
	Id                          string       `json:"id"`
	ResourceType                ResourceType `json:"resource_type"`
	ResourceId                  string       `json:"resource_id"`
	ActiveResourceId            string       `json:"active_resource_id"`
	DestinationBusinessEntityId string       `json:"destination_business_entity_id"`
	SourceBusinessEntityId      string       `json:"source_business_entity_id"`
	ReasonCode                  ReasonCode   `json:"reason_code"`
	CreatedAt                   int64        `json:"created_at"`
	Object                      string       `json:"object"`
}
