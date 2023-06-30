package paymentreferencenumber

import (
	paymentReferenceNumberEnum "github.com/chargebee/chargebee-go/models/paymentreferencenumber/enum"
)

type PaymentReferenceNumber struct {
	Id        string                          `json:"id"`
	Type      paymentReferenceNumberEnum.Type `json:"type"`
	Number    string                          `json:"number"`
	InvoiceId string                          `json:"invoice_id"`
	Object    string                          `json:"object"`
}
