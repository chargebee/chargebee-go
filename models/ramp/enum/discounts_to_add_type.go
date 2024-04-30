package enum

type DiscountsToAddType string

const (
	DiscountsToAddTypeFixedAmount DiscountsToAddType = "fixed_amount"
	DiscountsToAddTypePercentage  DiscountsToAddType = "percentage"
)
