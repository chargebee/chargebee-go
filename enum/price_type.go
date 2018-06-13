package enum

type PriceType string

const (
	PriceTypeTaxExclusive PriceType = "tax_exclusive"
	PriceTypeTaxInclusive PriceType = "tax_inclusive"
)
