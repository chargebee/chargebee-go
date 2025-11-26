package einvoice

import (
	einvoiceEnum "github.com/chargebee/chargebee-go/v3/models/einvoice/enum"
)

type Einvoice struct {
	Id              string              `json:"id"`
	ReferenceNumber string              `json:"reference_number"`
	Status          einvoiceEnum.Status `json:"status"`
	Message         string              `json:"message"`
	Object          string              `json:"object"`
}
