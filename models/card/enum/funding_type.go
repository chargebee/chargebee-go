package enum

type FundingType string

const (
	FundingTypeCredit        FundingType = "credit"
	FundingTypeDebit         FundingType = "debit"
	FundingTypePrepaid       FundingType = "prepaid"
	FundingTypeNotKnown      FundingType = "not_known"
	FundingTypeNotApplicable FundingType = "not_applicable"
)
