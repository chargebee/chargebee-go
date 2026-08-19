package vaultedpaymentmethod

type VaultedPaymentMethod struct {
	Id           string `json:"id"`
	CustomerId   string `json:"customer_id"`
	CreditCardId string `json:"credit_card_id"`
	CreatedAt    int64  `json:"created_at"`
	ModifiedAt   int64  `json:"modified_at"`
	Object       string `json:"object"`
}
