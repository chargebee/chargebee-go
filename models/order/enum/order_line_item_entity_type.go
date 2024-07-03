package enum

type OrderLineItemEntityType string

const (
	OrderLineItemEntityTypeAdhoc           OrderLineItemEntityType = "adhoc"
	OrderLineItemEntityTypePlanItemPrice   OrderLineItemEntityType = "plan_item_price"
	OrderLineItemEntityTypeAddonItemPrice  OrderLineItemEntityType = "addon_item_price"
	OrderLineItemEntityTypeChargeItemPrice OrderLineItemEntityType = "charge_item_price"
	OrderLineItemEntityTypePlanSetup       OrderLineItemEntityType = "plan_setup"
	OrderLineItemEntityTypePlan            OrderLineItemEntityType = "plan"
	OrderLineItemEntityTypeAddon           OrderLineItemEntityType = "addon"
)
