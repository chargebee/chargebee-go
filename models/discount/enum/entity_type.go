package enum

type EntityType string

const (
	EntityTypeItemLevelCoupon       EntityType = "item_level_coupon"
	EntityTypeDocumentLevelCoupon   EntityType = "document_level_coupon"
	EntityTypeItemLevelDiscount     EntityType = "item_level_discount"
	EntityTypeDocumentLevelDiscount EntityType = "document_level_discount"
)
