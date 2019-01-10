package enum

type LinkedOrderStatus string

const (
	LinkedOrderStatusNew                LinkedOrderStatus = "new"
	LinkedOrderStatusProcessing         LinkedOrderStatus = "processing"
	LinkedOrderStatusComplete           LinkedOrderStatus = "complete"
	LinkedOrderStatusCancelled          LinkedOrderStatus = "cancelled"
	LinkedOrderStatusVoided             LinkedOrderStatus = "voided"
	LinkedOrderStatusQueued             LinkedOrderStatus = "queued"
	LinkedOrderStatusAwaitingShipment   LinkedOrderStatus = "awaiting_shipment"
	LinkedOrderStatusOnHold             LinkedOrderStatus = "on_hold"
	LinkedOrderStatusDelivered          LinkedOrderStatus = "delivered"
	LinkedOrderStatusShipped            LinkedOrderStatus = "shipped"
	LinkedOrderStatusPartiallyDelivered LinkedOrderStatus = "partially_delivered"
	LinkedOrderStatusReturned           LinkedOrderStatus = "returned"
)
