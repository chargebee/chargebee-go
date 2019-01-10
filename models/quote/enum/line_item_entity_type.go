package enum

type LineItemEntityType string

const (
	LineItemEntityTypePlanSetup LineItemEntityType = "plan_setup"
	LineItemEntityTypePlan      LineItemEntityType = "plan"
	LineItemEntityTypeAddon     LineItemEntityType = "addon"
	LineItemEntityTypeAdhoc     LineItemEntityType = "adhoc"
)
