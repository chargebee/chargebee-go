package enum

type ItemType string

const (
	ItemTypePlan         ItemType = "plan"
	ItemTypeAddon        ItemType = "addon"
	ItemTypeCharge       ItemType = "charge"
	ItemTypeSubscription ItemType = "subscription"
	ItemTypeItem         ItemType = "item"
)
