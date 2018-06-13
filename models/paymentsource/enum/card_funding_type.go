package enum

type CardFundingType string

const (
	CardFundingTypeCredit        CardFundingType = "credit"
	CardFundingTypeDebit         CardFundingType = "debit"
	CardFundingTypePrepaid       CardFundingType = "prepaid"
	CardFundingTypeNotKnown      CardFundingType = "not_known"
	CardFundingTypeNotApplicable CardFundingType = "not_applicable"
)
