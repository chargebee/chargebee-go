package enum

type TaxOverrideReason string

const (
	TaxOverrideReasonZeroRated                        TaxOverrideReason = "zero_rated"
	TaxOverrideReasonIdExempt                         TaxOverrideReason = "id_exempt"
	TaxOverrideReasonCustomerExempt                   TaxOverrideReason = "customer_exempt"
	TaxOverrideReasonRegionNonTaxable                 TaxOverrideReason = "region_non_taxable"
	TaxOverrideReasonProductExempt                    TaxOverrideReason = "product_exempt"
	TaxOverrideReasonExport                           TaxOverrideReason = "export"
	TaxOverrideReasonHighValuePhysicalGoods           TaxOverrideReason = "high_value_physical_goods"
	TaxOverrideReasonZeroValueItem                    TaxOverrideReason = "zero_value_item"
	TaxOverrideReasonTaxNotConfiguredExternalProvider TaxOverrideReason = "tax_not_configured_external_provider"
)
