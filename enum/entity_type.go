package enum

type EntityType string

const (
	EntityTypeCustomer          EntityType = "customer"
	EntityTypeSubscription      EntityType = "subscription"
	EntityTypeInvoice           EntityType = "invoice"
	EntityTypeQuote             EntityType = "quote"
	EntityTypeCreditNote        EntityType = "credit_note"
	EntityTypeTransaction       EntityType = "transaction"
	EntityTypePlan              EntityType = "plan"
	EntityTypeAddon             EntityType = "addon"
	EntityTypeCoupon            EntityType = "coupon"
	EntityTypeItemFamily        EntityType = "item_family"
	EntityTypeItem              EntityType = "item"
	EntityTypeItemPrice         EntityType = "item_price"
	EntityTypePlanItem          EntityType = "plan_item"
	EntityTypeAddonItem         EntityType = "addon_item"
	EntityTypeChargeItem        EntityType = "charge_item"
	EntityTypePlanPrice         EntityType = "plan_price"
	EntityTypeAddonPrice        EntityType = "addon_price"
	EntityTypeChargePrice       EntityType = "charge_price"
	EntityTypeDifferentialPrice EntityType = "differential_price"
	EntityTypeAttachedItem      EntityType = "attached_item"
)
