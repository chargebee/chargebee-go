package enum

type OmnichannelSubscriptionItemStatus string

const (
	OmnichannelSubscriptionItemStatusActive        OmnichannelSubscriptionItemStatus = "active"
	OmnichannelSubscriptionItemStatusExpired       OmnichannelSubscriptionItemStatus = "expired"
	OmnichannelSubscriptionItemStatusCancelled     OmnichannelSubscriptionItemStatus = "cancelled"
	OmnichannelSubscriptionItemStatusInDunning     OmnichannelSubscriptionItemStatus = "in_dunning"
	OmnichannelSubscriptionItemStatusInGracePeriod OmnichannelSubscriptionItemStatus = "in_grace_period"
	OmnichannelSubscriptionItemStatusPaused        OmnichannelSubscriptionItemStatus = "paused"
)
