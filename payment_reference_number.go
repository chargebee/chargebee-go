package chargebee

type Type string

const (
	TypeKid            Type = "kid"
	TypeOcr            Type = "ocr"
	TypeFrn            Type = "frn"
	TypeFik            Type = "fik"
	TypeSwissReference Type = "swiss_reference"
)

type PaymentReferenceNumber struct {
	Id        string `json:"id"`
	Type      Type   `json:"type"`
	Number    string `json:"number"`
	InvoiceId string `json:"invoice_id"`
	Object    string `json:"object"`
}
