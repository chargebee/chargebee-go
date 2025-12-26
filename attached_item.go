package chargebee

type AttachedItemType string

const (
	AttachedItemTypeRecommended AttachedItemType = "recommended"
	AttachedItemTypeMandatory   AttachedItemType = "mandatory"
	AttachedItemTypeOptional    AttachedItemType = "optional"
)

type AttachedItemStatus string

const (
	AttachedItemStatusActive   AttachedItemStatus = "active"
	AttachedItemStatusArchived AttachedItemStatus = "archived"
	AttachedItemStatusDeleted  AttachedItemStatus = "deleted"
)

// just struct
type AttachedItem struct {
	Id                string             `json:"id"`
	ParentItemId      string             `json:"parent_item_id"`
	ItemId            string             `json:"item_id"`
	Type              AttachedItemType   `json:"type"`
	Status            AttachedItemStatus `json:"status"`
	Quantity          int32              `json:"quantity"`
	QuantityInDecimal string             `json:"quantity_in_decimal"`
	BillingCycles     int32              `json:"billing_cycles"`
	ChargeOnEvent     ChargeOnEvent      `json:"charge_on_event"`
	ChargeOnce        bool               `json:"charge_once"`
	CreatedAt         int64              `json:"created_at"`
	ResourceVersion   int64              `json:"resource_version"`
	UpdatedAt         int64              `json:"updated_at"`
	Channel           Channel            `json:"channel"`
	BusinessEntityId  string             `json:"business_entity_id"`
	Deleted           bool               `json:"deleted"`
	Object            string             `json:"object"`
}

// sub resources
// operations
// input params
type AttachedItemCreateRequest struct {
	ItemId            string           `json:"item_id"`
	Type              AttachedItemType `json:"type,omitempty"`
	BillingCycles     *int32           `json:"billing_cycles,omitempty"`
	Quantity          *int32           `json:"quantity,omitempty"`
	QuantityInDecimal string           `json:"quantity_in_decimal,omitempty"`
	ChargeOnEvent     ChargeOnEvent    `json:"charge_on_event,omitempty"`
	ChargeOnce        *bool            `json:"charge_once,omitempty"`
	BusinessEntityId  string           `json:"business_entity_id,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *AttachedItemCreateRequest) payload() any { return r }

type AttachedItemUpdateRequest struct {
	ParentItemId      string           `json:"parent_item_id"`
	Type              AttachedItemType `json:"type,omitempty"`
	BillingCycles     *int32           `json:"billing_cycles,omitempty"`
	Quantity          *int32           `json:"quantity,omitempty"`
	QuantityInDecimal string           `json:"quantity_in_decimal,omitempty"`
	ChargeOnEvent     ChargeOnEvent    `json:"charge_on_event,omitempty"`
	ChargeOnce        *bool            `json:"charge_once,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *AttachedItemUpdateRequest) payload() any { return r }

type AttachedItemRetrieveRequest struct {
	ParentItemId string `json:"parent_item_id"`
	apiRequest   `json:"-" form:"-"`
}

func (r *AttachedItemRetrieveRequest) payload() any { return r }

type AttachedItemDeleteRequest struct {
	ParentItemId string `json:"parent_item_id"`
	apiRequest   `json:"-" form:"-"`
}

func (r *AttachedItemDeleteRequest) payload() any { return r }

type AttachedItemListRequest struct {
	Limit         *int32           `json:"limit,omitempty"`
	Offset        string           `json:"offset,omitempty"`
	Id            *StringFilter    `json:"id,omitempty"`
	ItemId        *StringFilter    `json:"item_id,omitempty"`
	Type          *EnumFilter      `json:"type,omitempty"`
	ItemType      *EnumFilter      `json:"item_type,omitempty"`
	ChargeOnEvent *EnumFilter      `json:"charge_on_event,omitempty"`
	UpdatedAt     *TimestampFilter `json:"updated_at,omitempty"`
	apiRequest    `json:"-" form:"-"`
}

func (r *AttachedItemListRequest) payload() any { return r }

// operation response
type AttachedItemCreateResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
	apiResponse
}

// operation response
type AttachedItemUpdateResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
	apiResponse
}

// operation response
type AttachedItemRetrieveResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
	apiResponse
}

// operation response
type AttachedItemDeleteResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
	apiResponse
}

// operation sub response
type AttachedItemListAttachedItemResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
}

// operation response
type AttachedItemListResponse struct {
	List       []*AttachedItemListAttachedItemResponse `json:"list,omitempty"`
	NextOffset string                                  `json:"next_offset,omitempty"`
	apiResponse
}
