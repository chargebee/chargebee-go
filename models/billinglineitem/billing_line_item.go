package billinglineitem

type BillingLineItem struct {
	Id            string  `json:"id"`
	InvoiceNumber string  `json:"invoice_number"`
	Quantity      float64 `json:"quantity"`
	Object        string  `json:"object"`
}
