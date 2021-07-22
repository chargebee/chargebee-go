package enum

type LineItemDiscountDiscountType string

const (
	LineItemDiscountDiscountTypeItemLevelCoupon     LineItemDiscountDiscountType = "item_level_coupon"
	LineItemDiscountDiscountTypeDocumentLevelCoupon LineItemDiscountDiscountType = "document_level_coupon"
	LineItemDiscountDiscountTypePromotionalCredits  LineItemDiscountDiscountType = "promotional_credits"
	LineItemDiscountDiscountTypeProratedCredits     LineItemDiscountDiscountType = "prorated_credits"
	LineItemDiscountDiscountTypeCustomDiscount      LineItemDiscountDiscountType = "custom_discount"
)
