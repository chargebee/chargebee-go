package enum

type DiscountType string

const (
	DiscountTypeFixedAmount DiscountType = "fixed_amount"
	DiscountTypePercentage  DiscountType = "percentage"
)
