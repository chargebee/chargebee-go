package enum

type ReasonCode string

const (
	ReasonCodeWriteOff                 ReasonCode = "write_off"
	ReasonCodeSubscriptionChange       ReasonCode = "subscription_change"
	ReasonCodeSubscriptionCancellation ReasonCode = "subscription_cancellation"
	ReasonCodeSubscriptionPause        ReasonCode = "subscription_pause"
	ReasonCodeChargeback               ReasonCode = "chargeback"
	ReasonCodeProductUnsatisfactory    ReasonCode = "product_unsatisfactory"
	ReasonCodeServiceUnsatisfactory    ReasonCode = "service_unsatisfactory"
	ReasonCodeOrderChange              ReasonCode = "order_change"
	ReasonCodeOrderCancellation        ReasonCode = "order_cancellation"
	ReasonCodeWaiver                   ReasonCode = "waiver"
	ReasonCodeOther                    ReasonCode = "other"
	ReasonCodeFraudulent               ReasonCode = "fraudulent"
)
