package enum

type AllocationTaxApplication string

const (
	AllocationTaxApplicationPreTax  AllocationTaxApplication = "pre_tax"
	AllocationTaxApplicationPostTax AllocationTaxApplication = "post_tax"
)
