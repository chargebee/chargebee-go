package enum

type BankAccountAccountType string

const (
	BankAccountAccountTypeChecking         BankAccountAccountType = "checking"
	BankAccountAccountTypeSavings          BankAccountAccountType = "savings"
	BankAccountAccountTypeBusinessChecking BankAccountAccountType = "business_checking"
	BankAccountAccountTypeCurrent          BankAccountAccountType = "current"
)
