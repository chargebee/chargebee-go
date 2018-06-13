package enum

type EntityType string

const (
	EntityTypePlanSetup EntityType = "plan_setup"
	EntityTypePlan      EntityType = "plan"
	EntityTypeAddon     EntityType = "addon"
	EntityTypeAdhoc     EntityType = "adhoc"
)
