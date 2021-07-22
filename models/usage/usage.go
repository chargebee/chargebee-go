package usage

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
)

type Usage struct {
	Id              string      `json:"id"`
	UsageDate       int64       `json:"usage_date"`
	SubscriptionId  string      `json:"subscription_id"`
	ItemPriceId     string      `json:"item_price_id"`
	InvoiceId       string      `json:"invoice_id"`
	LineItemId      string      `json:"line_item_id"`
	Quantity        string      `json:"quantity"`
	Source          enum.Source `json:"source"`
	Note            string      `json:"note"`
	ResourceVersion int64       `json:"resource_version"`
	UpdatedAt       int64       `json:"updated_at"`
	CreatedAt       int64       `json:"created_at"`
	Object          string      `json:"object"`
}
type CreateRequestParams struct {
	Id           string            `json:"id,omitempty"`
	ItemPriceId  string            `json:"item_price_id"`
	Quantity     string            `json:"quantity"`
	UsageDate    *int64            `json:"usage_date"`
	DedupeOption enum.DedupeOption `json:"dedupe_option,omitempty"`
	Note         string            `json:"note,omitempty"`
}
type RetrieveRequestParams struct {
	Id string `json:"id"`
}
type DeleteRequestParams struct {
	Id string `json:"id"`
}
type ListRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	Id             *filter.StringFilter    `json:"id,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id"`
	UsageDate      *filter.TimestampFilter `json:"usage_date,omitempty"`
	ItemPriceId    *filter.StringFilter    `json:"item_price_id,omitempty"`
	InvoiceId      *filter.StringFilter    `json:"invoice_id,omitempty"`
	Source         *filter.EnumFilter      `json:"source,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type PdfRequestParams struct {
	Invoice         *PdfInvoiceParams    `json:"invoice,omitempty"`
	DispositionType enum.DispositionType `json:"disposition_type,omitempty"`
}
type PdfInvoiceParams struct {
	Id string `json:"id"`
}
