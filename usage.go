package chargebee

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
type CreateRequest struct {
	Id           string            `json:"id,omitempty"`
	ItemPriceId  string            `json:"item_price_id"`
	Quantity     string            `json:"quantity"`
	UsageDate    *int64            `json:"usage_date"`
	DedupeOption enum.DedupeOption `json:"dedupe_option,omitempty"`
	Note         string            `json:"note,omitempty"`
}
type RetrieveRequest struct {
	Id string `json:"id"`
}
type DeleteRequest struct {
	Id string `json:"id"`
}
type ListRequest struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	Id             *filter.StringFilter    `json:"id,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id,omitempty"`
	UsageDate      *filter.TimestampFilter `json:"usage_date,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	ItemPriceId    *filter.StringFilter    `json:"item_price_id,omitempty"`
	InvoiceId      *filter.StringFilter    `json:"invoice_id,omitempty"`
	Source         *filter.EnumFilter      `json:"source,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type PdfRequest struct {
	Invoice         *PdfInvoice          `json:"invoice,omitempty"`
	DispositionType enum.DispositionType `json:"disposition_type,omitempty"`
}
type PdfInvoice struct {
	Id string `json:"id"`
}

type CreateResponse struct {
	Usage *Usage `json:"usage,omitempty"`
}

type RetrieveResponse struct {
	Usage *Usage `json:"usage,omitempty"`
}

type DeleteResponse struct {
	Usage *Usage `json:"usage,omitempty"`
}

type ListUsageResponse struct {
	Usage *Usage `json:"usage,omitempty"`
}

type ListResponse struct {
	List       []*ListUsageResponse `json:"list,omitempty"`
	NextOffset string               `json:"next_offset,omitempty"`
}

type PdfResponse struct {
	Download *download.Download `json:"download,omitempty"`
}
