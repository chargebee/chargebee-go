package enum

type EntityType string

const (
	EntityTypePlanPrice   EntityType = "plan_price"
	EntityTypeAddonPrice  EntityType = "addon_price"
	EntityTypeChargePrice EntityType = "charge_price"
	EntityTypeCharge      EntityType = "charge"
)
