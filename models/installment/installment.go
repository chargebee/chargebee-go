package installment

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	installmentEnum "github.com/chargebee/chargebee-go/v3/models/installment/enum"
)

type Installment struct {
	Id              string                 `json:"id"`
	InvoiceId       string                 `json:"invoice_id"`
	Date            int64                  `json:"date"`
	Amount          int64                  `json:"amount"`
	Status          installmentEnum.Status `json:"status"`
	CreatedAt       int64                  `json:"created_at"`
	ResourceVersion int64                  `json:"resource_version"`
	UpdatedAt       int64                  `json:"updated_at"`
	Object          string                 `json:"object"`
}
type ListRequestParams struct {
	Limit     *int32               `json:"limit,omitempty"`
	Offset    string               `json:"offset,omitempty"`
	SortBy    string               `json:"sort_by,omitempty"`
	InvoiceId *filter.StringFilter `json:"invoice_id"`
}
