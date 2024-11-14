package enum

type OmnichannelSubscriptionItemStatus string

const (
	OmnichannelSubscriptionItemStatusActive    OmnichannelSubscriptionItemStatus = "active"
	OmnichannelSubscriptionItemStatusExpired   OmnichannelSubscriptionItemStatus = "expired"
	OmnichannelSubscriptionItemStatusCancelled OmnichannelSubscriptionItemStatus = "cancelled"
)
