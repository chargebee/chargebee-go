package attacheditem

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	attachedItemEnum "github.com/chargebee/chargebee-go/models/attacheditem/enum"
)

type AttachedItem struct {
	Id                string                  `json:"id"`
	ParentItemId      string                  `json:"parent_item_id"`
	ItemId            string                  `json:"item_id"`
	Type              attachedItemEnum.Type   `json:"type"`
	Status            attachedItemEnum.Status `json:"status"`
	Quantity          int32                   `json:"quantity"`
	QuantityInDecimal string                  `json:"quantity_in_decimal"`
	BillingCycles     int32                   `json:"billing_cycles"`
	ChargeOnEvent     enum.ChargeOnEvent      `json:"charge_on_event"`
	ChargeOnce        bool                    `json:"charge_once"`
	CreatedAt         int64                   `json:"created_at"`
	ResourceVersion   int64                   `json:"resource_version"`
	UpdatedAt         int64                   `json:"updated_at"`
	Channel           enum.Channel            `json:"channel"`
	Object            string                  `json:"object"`
}
type CreateRequestParams struct {
	ItemId            string                `json:"item_id"`
	Type              attachedItemEnum.Type `json:"type,omitempty"`
	BillingCycles     *int32                `json:"billing_cycles,omitempty"`
	Quantity          *int32                `json:"quantity,omitempty"`
	QuantityInDecimal string                `json:"quantity_in_decimal,omitempty"`
	ChargeOnEvent     enum.ChargeOnEvent    `json:"charge_on_event,omitempty"`
	ChargeOnce        *bool                 `json:"charge_once,omitempty"`
}
type UpdateRequestParams struct {
	ParentItemId      string                `json:"parent_item_id"`
	Type              attachedItemEnum.Type `json:"type,omitempty"`
	BillingCycles     *int32                `json:"billing_cycles,omitempty"`
	Quantity          *int32                `json:"quantity,omitempty"`
	QuantityInDecimal string                `json:"quantity_in_decimal,omitempty"`
	ChargeOnEvent     enum.ChargeOnEvent    `json:"charge_on_event,omitempty"`
	ChargeOnce        *bool                 `json:"charge_once,omitempty"`
}
type RetrieveRequestParams struct {
	ParentItemId string `json:"parent_item_id"`
}
type DeleteRequestParams struct {
	ParentItemId string `json:"parent_item_id"`
}
type ListRequestParams struct {
	Limit         *int32                  `json:"limit,omitempty"`
	Offset        string                  `json:"offset,omitempty"`
	Id            *filter.StringFilter    `json:"id,omitempty"`
	ItemId        *filter.StringFilter    `json:"item_id,omitempty"`
	Type          *filter.EnumFilter      `json:"type,omitempty"`
	ItemType      *filter.EnumFilter      `json:"item_type,omitempty"`
	ChargeOnEvent *filter.EnumFilter      `json:"charge_on_event,omitempty"`
	UpdatedAt     *filter.TimestampFilter `json:"updated_at,omitempty"`
}
