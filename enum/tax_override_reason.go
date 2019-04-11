package enum

type TaxOverrideReason string

const (
	TaxOverrideReasonIdExempt       TaxOverrideReason = "id_exempt"
	TaxOverrideReasonCustomerExempt TaxOverrideReason = "customer_exempt"
	TaxOverrideReasonExport         TaxOverrideReason = "export"
)
