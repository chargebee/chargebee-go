package enum

type EntityType string

const (
	EntityTypeAdhoc           EntityType = "adhoc"
	EntityTypePlanItemPrice   EntityType = "plan_item_price"
	EntityTypeAddonItemPrice  EntityType = "addon_item_price"
	EntityTypeChargeItemPrice EntityType = "charge_item_price"
	EntityTypePlanSetup       EntityType = "plan_setup"
	EntityTypePlan            EntityType = "plan"
	EntityTypeAddon           EntityType = "addon"
)
