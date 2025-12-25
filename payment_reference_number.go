package chargebee

type PaymentReferenceNumberType string

const (
	PaymentReferenceNumberTypeKid            PaymentReferenceNumberType = "kid"
	PaymentReferenceNumberTypeOcr            PaymentReferenceNumberType = "ocr"
	PaymentReferenceNumberTypeFrn            PaymentReferenceNumberType = "frn"
	PaymentReferenceNumberTypeFik            PaymentReferenceNumberType = "fik"
	PaymentReferenceNumberTypeSwissReference PaymentReferenceNumberType = "swiss_reference"
)

// just struct
type PaymentReferenceNumber struct {
	Id        string                     `json:"id"`
	Type      PaymentReferenceNumberType `json:"type"`
	Number    string                     `json:"number"`
	InvoiceId string                     `json:"invoice_id"`
	Object    string                     `json:"object"`
}

// sub resources
// operations
// input params
