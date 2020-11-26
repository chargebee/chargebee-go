package enum

type SubscriptionItemChargeOnOption string

const (
	SubscriptionItemChargeOnOptionImmediately SubscriptionItemChargeOnOption = "immediately"
	SubscriptionItemChargeOnOptionOnEvent     SubscriptionItemChargeOnOption = "on_event"
)
