package enum

type LineItemEntityType string

const (
	LineItemEntityTypePlanSetup       LineItemEntityType = "plan_setup"
	LineItemEntityTypePlan            LineItemEntityType = "plan"
	LineItemEntityTypeAddon           LineItemEntityType = "addon"
	LineItemEntityTypeAdhoc           LineItemEntityType = "adhoc"
	LineItemEntityTypePlanItemPrice   LineItemEntityType = "plan_item_price"
	LineItemEntityTypeAddonItemPrice  LineItemEntityType = "addon_item_price"
	LineItemEntityTypeChargeItemPrice LineItemEntityType = "charge_item_price"
)
