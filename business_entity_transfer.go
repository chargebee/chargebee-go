package chargebee

type BusinessEntityTransferResourceType string

const (
	BusinessEntityTransferResourceTypeCustomer     BusinessEntityTransferResourceType = "customer"
	BusinessEntityTransferResourceTypeSubscription BusinessEntityTransferResourceType = "subscription"
)

type BusinessEntityTransferReasonCode string

const (
	BusinessEntityTransferReasonCodeCorrection BusinessEntityTransferReasonCode = "correction"
)

// just struct
type BusinessEntityTransfer struct {
	Id                          string                             `json:"id"`
	ResourceType                BusinessEntityTransferResourceType `json:"resource_type"`
	ResourceId                  string                             `json:"resource_id"`
	ActiveResourceId            string                             `json:"active_resource_id"`
	DestinationBusinessEntityId string                             `json:"destination_business_entity_id"`
	SourceBusinessEntityId      string                             `json:"source_business_entity_id"`
	ReasonCode                  BusinessEntityTransferReasonCode   `json:"reason_code"`
	CreatedAt                   int64                              `json:"created_at"`
	Object                      string                             `json:"object"`
}

// sub resources
// operations
// input params
