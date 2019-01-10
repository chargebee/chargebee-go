package enum

type OrderLineItemStatus string

const (
	OrderLineItemStatusQueued             OrderLineItemStatus = "queued"
	OrderLineItemStatusAwaitingShipment   OrderLineItemStatus = "awaiting_shipment"
	OrderLineItemStatusOnHold             OrderLineItemStatus = "on_hold"
	OrderLineItemStatusDelivered          OrderLineItemStatus = "delivered"
	OrderLineItemStatusShipped            OrderLineItemStatus = "shipped"
	OrderLineItemStatusPartiallyDelivered OrderLineItemStatus = "partially_delivered"
	OrderLineItemStatusReturned           OrderLineItemStatus = "returned"
	OrderLineItemStatusCancelled          OrderLineItemStatus = "cancelled"
)
