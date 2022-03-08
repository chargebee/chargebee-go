package enum

type AccountType string

const (
	AccountTypeChecking         AccountType = "checking"
	AccountTypeSavings          AccountType = "savings"
	AccountTypeBusinessChecking AccountType = "business_checking"
	AccountTypeCurrent          AccountType = "current"
)
