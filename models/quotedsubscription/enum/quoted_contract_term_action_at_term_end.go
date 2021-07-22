package enum

type QuotedContractTermActionAtTermEnd string

const (
	QuotedContractTermActionAtTermEndRenew     QuotedContractTermActionAtTermEnd = "renew"
	QuotedContractTermActionAtTermEndEvergreen QuotedContractTermActionAtTermEnd = "evergreen"
	QuotedContractTermActionAtTermEndCancel    QuotedContractTermActionAtTermEnd = "cancel"
	QuotedContractTermActionAtTermEndRenewOnce QuotedContractTermActionAtTermEnd = "renew_once"
)
