package enum

type UnbilledChargeEntityType string

const (
	UnbilledChargeEntityTypePlanSetup UnbilledChargeEntityType = "plan_setup"
	UnbilledChargeEntityTypePlan      UnbilledChargeEntityType = "plan"
	UnbilledChargeEntityTypeAddon     UnbilledChargeEntityType = "addon"
	UnbilledChargeEntityTypeAdhoc     UnbilledChargeEntityType = "adhoc"
)
