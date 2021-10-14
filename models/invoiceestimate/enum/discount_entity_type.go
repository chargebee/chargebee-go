package enum

type DiscountEntityType string

const (
	DiscountEntityTypeItemLevelCoupon       DiscountEntityType = "item_level_coupon"
	DiscountEntityTypeDocumentLevelCoupon   DiscountEntityType = "document_level_coupon"
	DiscountEntityTypePromotionalCredits    DiscountEntityType = "promotional_credits"
	DiscountEntityTypeProratedCredits       DiscountEntityType = "prorated_credits"
	DiscountEntityTypeItemLevelDiscount     DiscountEntityType = "item_level_discount"
	DiscountEntityTypeDocumentLevelDiscount DiscountEntityType = "document_level_discount"
)
