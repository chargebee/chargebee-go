package enum

type Status string

const (
	StatusNew                Status = "new"
	StatusProcessing         Status = "processing"
	StatusComplete           Status = "complete"
	StatusCancelled          Status = "cancelled"
	StatusVoided             Status = "voided"
	StatusQueued             Status = "queued"
	StatusAwaitingShipment   Status = "awaiting_shipment"
	StatusOnHold             Status = "on_hold"
	StatusDelivered          Status = "delivered"
	StatusShipped            Status = "shipped"
	StatusPartiallyDelivered Status = "partially_delivered"
	StatusReturned           Status = "returned"
)
