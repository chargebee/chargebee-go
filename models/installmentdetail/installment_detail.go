package installmentdetail

import (
	installmentDetailEnum "github.com/chargebee/chargebee-go/v3/models/installmentdetail/enum"
)

type InstallmentDetail struct {
	Id           string         `json:"id"`
	InvoiceId    string         `json:"invoice_id"`
	Amount       int64          `json:"amount"`
	Installments []*Installment `json:"installments"`
	Object       string         `json:"object"`
}
type Installment struct {
	Id              string                                  `json:"id"`
	InvoiceId       string                                  `json:"invoice_id"`
	Date            int64                                   `json:"date"`
	Amount          int64                                   `json:"amount"`
	Status          installmentDetailEnum.InstallmentStatus `json:"status"`
	CreatedAt       int64                                   `json:"created_at"`
	ResourceVersion int64                                   `json:"resource_version"`
	UpdatedAt       int64                                   `json:"updated_at"`
	Object          string                                  `json:"object"`
}
