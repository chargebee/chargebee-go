package enum

type OfflinePaymentMethod string

const (
	OfflinePaymentMethodNoPreference            OfflinePaymentMethod = "no_preference"
	OfflinePaymentMethodCash                    OfflinePaymentMethod = "cash"
	OfflinePaymentMethodCheck                   OfflinePaymentMethod = "check"
	OfflinePaymentMethodBankTransfer            OfflinePaymentMethod = "bank_transfer"
	OfflinePaymentMethodAchCredit               OfflinePaymentMethod = "ach_credit"
	OfflinePaymentMethodSepaCredit              OfflinePaymentMethod = "sepa_credit"
	OfflinePaymentMethodBoleto                  OfflinePaymentMethod = "boleto"
	OfflinePaymentMethodUsAutomatedBankTransfer OfflinePaymentMethod = "us_automated_bank_transfer"
	OfflinePaymentMethodEuAutomatedBankTransfer OfflinePaymentMethod = "eu_automated_bank_transfer"
	OfflinePaymentMethodUkAutomatedBankTransfer OfflinePaymentMethod = "uk_automated_bank_transfer"
	OfflinePaymentMethodJpAutomatedBankTransfer OfflinePaymentMethod = "jp_automated_bank_transfer"
	OfflinePaymentMethodMxAutomatedBankTransfer OfflinePaymentMethod = "mx_automated_bank_transfer"
	OfflinePaymentMethodCustom                  OfflinePaymentMethod = "custom"
)
