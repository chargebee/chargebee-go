package enum

type EntityType string

const (
	EntityTypeCustomer       EntityType = "customer"
	EntityTypeSubscription   EntityType = "subscription"
	EntityTypeInvoice        EntityType = "invoice"
	EntityTypeQuote          EntityType = "quote"
	EntityTypeCreditNote     EntityType = "credit_note"
	EntityTypeTransaction    EntityType = "transaction"
	EntityTypePlan           EntityType = "plan"
	EntityTypeAddon          EntityType = "addon"
	EntityTypeCoupon         EntityType = "coupon"
	EntityTypeOrder          EntityType = "order"
	EntityTypeBusinessEntity EntityType = "business_entity"
	EntityTypeItemFamily     EntityType = "item_family"
	EntityTypeItem           EntityType = "item"
	EntityTypeItemPrice      EntityType = "item_price"
	EntityTypeProduct        EntityType = "product"
	EntityTypeVariant        EntityType = "variant"
)
