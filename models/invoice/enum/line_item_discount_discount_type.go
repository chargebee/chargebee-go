package enum

type LineItemDiscountDiscountType string

const (
	LineItemDiscountDiscountTypeItemLevelCoupon       LineItemDiscountDiscountType = "item_level_coupon"
	LineItemDiscountDiscountTypeDocumentLevelCoupon   LineItemDiscountDiscountType = "document_level_coupon"
	LineItemDiscountDiscountTypePromotionalCredits    LineItemDiscountDiscountType = "promotional_credits"
	LineItemDiscountDiscountTypeProratedCredits       LineItemDiscountDiscountType = "prorated_credits"
	LineItemDiscountDiscountTypeItemLevelDiscount     LineItemDiscountDiscountType = "item_level_discount"
	LineItemDiscountDiscountTypeDocumentLevelDiscount LineItemDiscountDiscountType = "document_level_discount"
)
