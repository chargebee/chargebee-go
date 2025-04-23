package enum

type PricingType string

const (
	PricingTypePerUnit PricingType = "per_unit"
	PricingTypeFlatFee PricingType = "flat_fee"
	PricingTypePackage PricingType = "package"
)
