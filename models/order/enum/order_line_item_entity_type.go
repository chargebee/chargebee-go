package enum

type OrderLineItemEntityType string

const (
	OrderLineItemEntityTypePlanSetup       OrderLineItemEntityType = "plan_setup"
	OrderLineItemEntityTypePlan            OrderLineItemEntityType = "plan"
	OrderLineItemEntityTypeAddon           OrderLineItemEntityType = "addon"
	OrderLineItemEntityTypeAdhoc           OrderLineItemEntityType = "adhoc"
	OrderLineItemEntityTypePlanItemPrice   OrderLineItemEntityType = "plan_item_price"
	OrderLineItemEntityTypeAddonItemPrice  OrderLineItemEntityType = "addon_item_price"
	OrderLineItemEntityTypeChargeItemPrice OrderLineItemEntityType = "charge_item_price"
)
