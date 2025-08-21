package enum

type UnbilledChargeEntityType string

const (
	UnbilledChargeEntityTypeAdhoc           UnbilledChargeEntityType = "adhoc"
	UnbilledChargeEntityTypePlanItemPrice   UnbilledChargeEntityType = "plan_item_price"
	UnbilledChargeEntityTypeAddonItemPrice  UnbilledChargeEntityType = "addon_item_price"
	UnbilledChargeEntityTypeChargeItemPrice UnbilledChargeEntityType = "charge_item_price"
	UnbilledChargeEntityTypePlanSetup       UnbilledChargeEntityType = "plan_setup"
	UnbilledChargeEntityTypePlan            UnbilledChargeEntityType = "plan"
	UnbilledChargeEntityTypeAddon           UnbilledChargeEntityType = "addon"
)
