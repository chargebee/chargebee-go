package einvoice

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/v3/filter"
	einvoiceEnum "github.com/chargebee/chargebee-go/v3/models/einvoice/enum"
)

type Einvoice struct {
	Id                 string                  `json:"id"`
	EntityType         einvoiceEnum.EntityType `json:"entity_type"`
	EntityId           string                  `json:"entity_id"`
	ReferenceId        string                  `json:"reference_id"`
	ReferenceNumber    string                  `json:"reference_number"`
	Status             einvoiceEnum.Status     `json:"status"`
	Message            string                  `json:"message"`
	CreatedAt          int64                   `json:"created_at"`
	ResourceVersion    int64                   `json:"resource_version"`
	UpdatedAt          int64                   `json:"updated_at"`
	Deleted            bool                    `json:"deleted"`
	ProviderReferences json.RawMessage         `json:"provider_references"`
	BusinessEntityId   string                  `json:"business_entity_id"`
	Artifacts          []*Artifact             `json:"artifacts"`
	Object             string                  `json:"object"`
}
type Artifact struct {
	ArtifactType       string                                 `json:"artifact_type"`
	Direction          einvoiceEnum.EinvoiceArtifactDirection `json:"direction"`
	Status             einvoiceEnum.EinvoiceArtifactStatus    `json:"status"`
	Code               string                                 `json:"code"`
	ExternalArtifactId string                                 `json:"external_artifact_id"`
	CreatedAt          int64                                  `json:"created_at"`
	ResourceVersion    int64                                  `json:"resource_version"`
	UpdatedAt          int64                                  `json:"updated_at"`
	Deleted            bool                                   `json:"deleted"`
	Object             string                                 `json:"object"`
}
type ListEinvoicesRequestParams struct {
	Limit        *int32                  `json:"limit,omitempty"`
	Offset       string                  `json:"offset,omitempty"`
	Id           *filter.StringFilter    `json:"id,omitempty"`
	ReferenceId  *filter.StringFilter    `json:"reference_id,omitempty"`
	UpdatedAt    *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy       *filter.SortFilter      `json:"sort_by,omitempty"`
	InvoiceId    string                  `json:"invoice_id,omitempty"`
	CreditNoteId string                  `json:"credit_note_id,omitempty"`
}
