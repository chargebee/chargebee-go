package enum

type UnbilledChargeEstimateEntityType string

const (
	UnbilledChargeEstimateEntityTypeAdhoc           UnbilledChargeEstimateEntityType = "adhoc"
	UnbilledChargeEstimateEntityTypePlanItemPrice   UnbilledChargeEstimateEntityType = "plan_item_price"
	UnbilledChargeEstimateEntityTypeAddonItemPrice  UnbilledChargeEstimateEntityType = "addon_item_price"
	UnbilledChargeEstimateEntityTypeChargeItemPrice UnbilledChargeEstimateEntityType = "charge_item_price"
	UnbilledChargeEstimateEntityTypePlanSetup       UnbilledChargeEstimateEntityType = "plan_setup"
	UnbilledChargeEstimateEntityTypePlan            UnbilledChargeEstimateEntityType = "plan"
	UnbilledChargeEstimateEntityTypeAddon           UnbilledChargeEstimateEntityType = "addon"
)
