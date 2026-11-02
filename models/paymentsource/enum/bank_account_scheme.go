package enum

type BankAccountScheme string

const (
	BankAccountSchemeAch      BankAccountScheme = "ach"
	BankAccountSchemeBacs     BankAccountScheme = "bacs"
	BankAccountSchemeSepaCore BankAccountScheme = "sepa_core"
	BankAccountSchemeAutogiro BankAccountScheme = "autogiro"
	BankAccountSchemeBecs     BankAccountScheme = "becs"
	BankAccountSchemeBecsNz   BankAccountScheme = "becs_nz"
	BankAccountSchemePad      BankAccountScheme = "pad"
)
