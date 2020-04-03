package enum

type ContractTermStatus string

const (
	ContractTermStatusActive     ContractTermStatus = "active"
	ContractTermStatusCompleted  ContractTermStatus = "completed"
	ContractTermStatusCancelled  ContractTermStatus = "cancelled"
	ContractTermStatusTerminated ContractTermStatus = "terminated"
)
