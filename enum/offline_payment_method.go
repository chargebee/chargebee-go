package enum

type OfflinePaymentMethod string

const (
	OfflinePaymentMethodNoPreference OfflinePaymentMethod = "no_preference"
	OfflinePaymentMethodCash         OfflinePaymentMethod = "cash"
	OfflinePaymentMethodCheck        OfflinePaymentMethod = "check"
	OfflinePaymentMethodBankTransfer OfflinePaymentMethod = "bank_transfer"
	OfflinePaymentMethodAchCredit    OfflinePaymentMethod = "ach_credit"
	OfflinePaymentMethodSepaCredit   OfflinePaymentMethod = "sepa_credit"
)
