package enum

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "new"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusComplete   OrderStatus = "complete"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusVoided     OrderStatus = "voided"
)
