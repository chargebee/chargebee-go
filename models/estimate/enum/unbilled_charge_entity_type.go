package enum

type UnbilledChargeEntityType string

const (
	UnbilledChargeEntityTypePlanSetup       UnbilledChargeEntityType = "plan_setup"
	UnbilledChargeEntityTypePlan            UnbilledChargeEntityType = "plan"
	UnbilledChargeEntityTypeAddon           UnbilledChargeEntityType = "addon"
	UnbilledChargeEntityTypeAdhoc           UnbilledChargeEntityType = "adhoc"
	UnbilledChargeEntityTypePlanItemPrice   UnbilledChargeEntityType = "plan_item_price"
	UnbilledChargeEntityTypeAddonItemPrice  UnbilledChargeEntityType = "addon_item_price"
	UnbilledChargeEntityTypeChargeItemPrice UnbilledChargeEntityType = "charge_item_price"
)
