package enum

type ContractTermActionAtTermEnd string

const (
	ContractTermActionAtTermEndRenew     ContractTermActionAtTermEnd = "renew"
	ContractTermActionAtTermEndEvergreen ContractTermActionAtTermEnd = "evergreen"
	ContractTermActionAtTermEndCancel    ContractTermActionAtTermEnd = "cancel"
	ContractTermActionAtTermEndRenewOnce ContractTermActionAtTermEnd = "renew_once"
)
