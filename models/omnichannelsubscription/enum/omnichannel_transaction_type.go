package enum

type OmnichannelTransactionType string

const (
	OmnichannelTransactionTypePurchase OmnichannelTransactionType = "purchase"
	OmnichannelTransactionTypeRenewal  OmnichannelTransactionType = "renewal"
)
