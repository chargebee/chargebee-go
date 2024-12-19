package enum

type InitialPurchaseTransactionType string

const (
	InitialPurchaseTransactionTypePurchase InitialPurchaseTransactionType = "purchase"
	InitialPurchaseTransactionTypeRenewal  InitialPurchaseTransactionType = "renewal"
)
