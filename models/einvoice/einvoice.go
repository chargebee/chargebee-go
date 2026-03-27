package einvoice

import (
	"encoding/json"
	einvoiceEnum "github.com/chargebee/chargebee-go/v3/models/einvoice/enum"
)

type Einvoice struct {
	Id                 string              `json:"id"`
	ReferenceId        string              `json:"reference_id"`
	ReferenceNumber    string              `json:"reference_number"`
	Status             einvoiceEnum.Status `json:"status"`
	Message            string              `json:"message"`
	ProviderReferences json.RawMessage     `json:"provider_references"`
	Object             string              `json:"object"`
}
