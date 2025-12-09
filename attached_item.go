package chargebee

type Type string

const (
	TypeRecommended Type = "recommended"
	TypeMandatory   Type = "mandatory"
	TypeOptional    Type = "optional"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
	StatusDeleted  Status = "deleted"
)

type AttachedItem struct {
	Id                string             `json:"id"`
	ParentItemId      string             `json:"parent_item_id"`
	ItemId            string             `json:"item_id"`
	Type              Type               `json:"type"`
	Status            Status             `json:"status"`
	Quantity          int32              `json:"quantity"`
	QuantityInDecimal string             `json:"quantity_in_decimal"`
	BillingCycles     int32              `json:"billing_cycles"`
	ChargeOnEvent     enum.ChargeOnEvent `json:"charge_on_event"`
	ChargeOnce        bool               `json:"charge_once"`
	CreatedAt         int64              `json:"created_at"`
	ResourceVersion   int64              `json:"resource_version"`
	UpdatedAt         int64              `json:"updated_at"`
	Channel           enum.Channel       `json:"channel"`
	BusinessEntityId  string             `json:"business_entity_id"`
	Deleted           bool               `json:"deleted"`
	Object            string             `json:"object"`
}
type CreateRequest struct {
	ItemId            string             `json:"item_id"`
	Type              Type               `json:"type,omitempty"`
	BillingCycles     *int32             `json:"billing_cycles,omitempty"`
	Quantity          *int32             `json:"quantity,omitempty"`
	QuantityInDecimal string             `json:"quantity_in_decimal,omitempty"`
	ChargeOnEvent     enum.ChargeOnEvent `json:"charge_on_event,omitempty"`
	ChargeOnce        *bool              `json:"charge_once,omitempty"`
	BusinessEntityId  string             `json:"business_entity_id,omitempty"`
}
type UpdateRequest struct {
	ParentItemId      string             `json:"parent_item_id"`
	Type              Type               `json:"type,omitempty"`
	BillingCycles     *int32             `json:"billing_cycles,omitempty"`
	Quantity          *int32             `json:"quantity,omitempty"`
	QuantityInDecimal string             `json:"quantity_in_decimal,omitempty"`
	ChargeOnEvent     enum.ChargeOnEvent `json:"charge_on_event,omitempty"`
	ChargeOnce        *bool              `json:"charge_once,omitempty"`
}
type RetrieveRequest struct {
	ParentItemId string `json:"parent_item_id"`
}
type DeleteRequest struct {
	ParentItemId string `json:"parent_item_id"`
}
type ListRequest struct {
	Limit         *int32                  `json:"limit,omitempty"`
	Offset        string                  `json:"offset,omitempty"`
	Id            *filter.StringFilter    `json:"id,omitempty"`
	ItemId        *filter.StringFilter    `json:"item_id,omitempty"`
	Type          *filter.EnumFilter      `json:"type,omitempty"`
	ItemType      *filter.EnumFilter      `json:"item_type,omitempty"`
	ChargeOnEvent *filter.EnumFilter      `json:"charge_on_event,omitempty"`
	UpdatedAt     *filter.TimestampFilter `json:"updated_at,omitempty"`
}

type CreateResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
}

type UpdateResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
}

type RetrieveResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
}

type DeleteResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
}

type ListAttachedItemResponse struct {
	AttachedItem *AttachedItem `json:"attached_item,omitempty"`
}

type ListResponse struct {
	List       []*ListAttachedItemResponse `json:"list,omitempty"`
	NextOffset string                      `json:"next_offset,omitempty"`
}
