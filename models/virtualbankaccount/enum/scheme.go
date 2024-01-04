package enum

type Scheme string

const (
	SchemeAchCredit               Scheme = "ach_credit"
	SchemeSepaCredit              Scheme = "sepa_credit"
	SchemeUsAutomatedBankTransfer Scheme = "us_automated_bank_transfer"
	SchemeGbAutomatedBankTransfer Scheme = "gb_automated_bank_transfer"
	SchemeEuAutomatedBankTransfer Scheme = "eu_automated_bank_transfer"
	SchemeJpAutomatedBankTransfer Scheme = "jp_automated_bank_transfer"
	SchemeMxAutomatedBankTransfer Scheme = "mx_automated_bank_transfer"
)
