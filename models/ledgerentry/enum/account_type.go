package enum

type AccountType string

const (
	AccountTypeProvisioned AccountType = "provisioned"
	AccountTypeOverdraft   AccountType = "overdraft"
)
