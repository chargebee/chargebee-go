package enum

type TaxExemptReason string

const (
	TaxExemptReasonTaxNotConfigured TaxExemptReason = "tax_not_configured"
	TaxExemptReasonRegionNonTaxable TaxExemptReason = "region_non_taxable"
	TaxExemptReasonExport           TaxExemptReason = "export"
	TaxExemptReasonCustomerExempt   TaxExemptReason = "customer_exempt"
	TaxExemptReasonProductExempt    TaxExemptReason = "product_exempt"
	TaxExemptReasonZeroRated        TaxExemptReason = "zero_rated"
	TaxExemptReasonReverseCharge    TaxExemptReason = "reverse_charge"
)
