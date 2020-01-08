package enum

type TaxjarExemptionCategory string

const (
	TaxjarExemptionCategoryWholesale  TaxjarExemptionCategory = "wholesale"
	TaxjarExemptionCategoryGovernment TaxjarExemptionCategory = "government"
	TaxjarExemptionCategoryOther      TaxjarExemptionCategory = "other"
)
