package enum

type DiscountEntityType string

const (
	DiscountEntityTypeItemLevelCoupon       DiscountEntityType = "item_level_coupon"
	DiscountEntityTypeDocumentLevelCoupon   DiscountEntityType = "document_level_coupon"
	DiscountEntityTypeItemLevelDiscount     DiscountEntityType = "item_level_discount"
	DiscountEntityTypeDocumentLevelDiscount DiscountEntityType = "document_level_discount"
)
