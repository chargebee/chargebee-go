package order

import (
	"github.com/chargebee/chargebee-go/filter"
	orderEnum "github.com/chargebee/chargebee-go/models/order/enum"
)

type Order struct {
	Id                string           `json:"id"`
	InvoiceId         string           `json:"invoice_id"`
	Status            orderEnum.Status `json:"status"`
	ReferenceId       string           `json:"reference_id"`
	FulfillmentStatus string           `json:"fulfillment_status"`
	Note              string           `json:"note"`
	TrackingId        string           `json:"tracking_id"`
	BatchId           string           `json:"batch_id"`
	CreatedBy         string           `json:"created_by"`
	CreatedAt         int64            `json:"created_at"`
	StatusUpdateAt    int64            `json:"status_update_at"`
	Object            string           `json:"object"`
}
type CreateRequestParams struct {
	Id                string           `json:"id,omitempty"`
	InvoiceId         string           `json:"invoice_id"`
	Status            orderEnum.Status `json:"status,omitempty"`
	ReferenceId       string           `json:"reference_id,omitempty"`
	FulfillmentStatus string           `json:"fulfillment_status,omitempty"`
	Note              string           `json:"note,omitempty"`
	TrackingId        string           `json:"tracking_id,omitempty"`
	BatchId           string           `json:"batch_id,omitempty"`
}

type UpdateRequestParams struct {
	Status            orderEnum.Status `json:"status,omitempty"`
	ReferenceId       string           `json:"reference_id,omitempty"`
	FulfillmentStatus string           `json:"fulfillment_status,omitempty"`
	Note              string           `json:"note,omitempty"`
	TrackingId        string           `json:"tracking_id,omitempty"`
	BatchId           string           `json:"batch_id,omitempty"`
}

type ListRequestParams struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Offset    string                  `json:"offset,omitempty"`
	Id        *filter.StringFilter    `json:"id,omitempty"`
	InvoiceId *filter.StringFilter    `json:"invoice_id,omitempty"`
	Status    *filter.EnumFilter      `json:"status,omitempty"`
	CreatedAt *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy    *filter.SortFilter      `json:"sort_by,omitempty"`
}

type OrdersForInvoiceRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
