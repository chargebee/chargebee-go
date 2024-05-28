package businessentitytransfer

import (
	businessEntityTransferEnum "github.com/chargebee/chargebee-go/v3/models/businessentitytransfer/enum"
)

type BusinessEntityTransfer struct {
	Id                          string                                  `json:"id"`
	ResourceType                businessEntityTransferEnum.ResourceType `json:"resource_type"`
	ResourceId                  string                                  `json:"resource_id"`
	ActiveResourceId            string                                  `json:"active_resource_id"`
	DestinationBusinessEntityId string                                  `json:"destination_business_entity_id"`
	SourceBusinessEntityId      string                                  `json:"source_business_entity_id"`
	ReasonCode                  businessEntityTransferEnum.ReasonCode   `json:"reason_code"`
	CreatedAt                   int64                                   `json:"created_at"`
	Object                      string                                  `json:"object"`
}
