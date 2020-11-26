package enum

type SubscriptionItemItemType string

const (
	SubscriptionItemItemTypePlan   SubscriptionItemItemType = "plan"
	SubscriptionItemItemTypeAddon  SubscriptionItemItemType = "addon"
	SubscriptionItemItemTypeCharge SubscriptionItemItemType = "charge"
)
