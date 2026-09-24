package enum

type Taxability string

const (
	TaxabilityTaxable   Taxability = "taxable"
	TaxabilityExempt    Taxability = "exempt"
	TaxabilityZeroRated Taxability = "zero_rated"
)
