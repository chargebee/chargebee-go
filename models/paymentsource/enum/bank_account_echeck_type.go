package enum

type BankAccountEcheckType string

const (
	BankAccountEcheckTypeWeb BankAccountEcheckType = "web"
	BankAccountEcheckTypePpd BankAccountEcheckType = "ppd"
	BankAccountEcheckTypeCcd BankAccountEcheckType = "ccd"
)
