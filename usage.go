package chargebee

// just struct
type Usage struct {
	Id              string `json:"id"`
	UsageDate       int64  `json:"usage_date"`
	SubscriptionId  string `json:"subscription_id"`
	ItemPriceId     string `json:"item_price_id"`
	InvoiceId       string `json:"invoice_id"`
	LineItemId      string `json:"line_item_id"`
	Quantity        string `json:"quantity"`
	Source          Source `json:"source"`
	Note            string `json:"note"`
	ResourceVersion int64  `json:"resource_version"`
	UpdatedAt       int64  `json:"updated_at"`
	CreatedAt       int64  `json:"created_at"`
	Object          string `json:"object"`
}

// sub resources
// operations
// input params
type UsageCreateRequest struct {
	Id          string `json:"id,omitempty"`
	ItemPriceId string `json:"item_price_id"`
	Quantity    string `json:"quantity"`
	UsageDate   *int64 `json:"usage_date"`
	Note        string `json:"note,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *UsageCreateRequest) payload() any { return r }

type UsageRetrieveRequest struct {
	Id         string `json:"id"`
	apiRequest `json:"-" form:"-"`
}

func (r *UsageRetrieveRequest) payload() any { return r }

type UsageDeleteRequest struct {
	Id         string `json:"id"`
	apiRequest `json:"-" form:"-"`
}

func (r *UsageDeleteRequest) payload() any { return r }

type UsageListRequest struct {
	Limit          *int32           `json:"limit,omitempty"`
	Offset         string           `json:"offset,omitempty"`
	Id             *StringFilter    `json:"id,omitempty"`
	SubscriptionId *StringFilter    `json:"subscription_id,omitempty"`
	UsageDate      *TimestampFilter `json:"usage_date,omitempty"`
	UpdatedAt      *TimestampFilter `json:"updated_at,omitempty"`
	ItemPriceId    *StringFilter    `json:"item_price_id,omitempty"`
	InvoiceId      *StringFilter    `json:"invoice_id,omitempty"`
	Source         *EnumFilter      `json:"source,omitempty"`
	SortBy         *SortFilter      `json:"sort_by,omitempty"`
	apiRequest     `json:"-" form:"-"`
}

func (r *UsageListRequest) payload() any { return r }

type UsagePdfRequest struct {
	Invoice         *UsagePdfInvoice `json:"invoice,omitempty"`
	DispositionType DispositionType  `json:"disposition_type,omitempty"`
	apiRequest      `json:"-" form:"-"`
}

func (r *UsagePdfRequest) payload() any { return r }

// input sub resource params single
type UsagePdfInvoice struct {
	Id string `json:"id"`
}

// operation response
type UsageCreateResponse struct {
	Usage *Usage `json:"usage,omitempty"`
	apiResponse
}

// operation response
type UsageRetrieveResponse struct {
	Usage *Usage `json:"usage,omitempty"`
	apiResponse
}

// operation response
type UsageDeleteResponse struct {
	Usage *Usage `json:"usage,omitempty"`
	apiResponse
}

// operation sub response
type UsageListUsageResponse struct {
	Usage *Usage `json:"usage,omitempty"`
}

// operation response
type UsageListResponse struct {
	List       []*UsageListUsageResponse `json:"list,omitempty"`
	NextOffset string                    `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type UsagePdfResponse struct {
	Download Download `json:"download,omitempty"`
	apiResponse
}
