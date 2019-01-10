package enum

type OrderLineItemEntityType string

const (
	OrderLineItemEntityTypePlanSetup OrderLineItemEntityType = "plan_setup"
	OrderLineItemEntityTypePlan      OrderLineItemEntityType = "plan"
	OrderLineItemEntityTypeAddon     OrderLineItemEntityType = "addon"
	OrderLineItemEntityTypeAdhoc     OrderLineItemEntityType = "adhoc"
)
