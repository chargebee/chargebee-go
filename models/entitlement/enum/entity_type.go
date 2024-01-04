package enum

type EntityType string

const (
	EntityTypePlan       EntityType = "plan"
	EntityTypeAddon      EntityType = "addon"
	EntityTypeCharge     EntityType = "charge"
	EntityTypePlanPrice  EntityType = "plan_price"
	EntityTypeAddonPrice EntityType = "addon_price"
)
