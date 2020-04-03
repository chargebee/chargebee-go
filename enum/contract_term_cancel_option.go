package enum

type ContractTermCancelOption string

const (
	ContractTermCancelOptionTerminateImmediately ContractTermCancelOption = "terminate_immediately"
	ContractTermCancelOptionEndOfContractTerm    ContractTermCancelOption = "end_of_contract_term"
)
