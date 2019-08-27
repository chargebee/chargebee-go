package paymentintent

import (
	paymentIntentEnum "github.com/chargebee/chargebee-go/models/paymentintent/enum"
)

type PaymentIntent struct {
	Id                   string                   `json:"id"`
	Status               paymentIntentEnum.Status `json:"status"`
	CurrencyCode         string                   `json:"currency_code"`
	Amount               int32                    `json:"amount"`
	GatewayAccountId     string                   `json:"gateway_account_id"`
	ExpiresAt            int64                    `json:"expires_at"`
	ReferenceId          string                   `json:"reference_id"`
	CreatedAt            int64                    `json:"created_at"`
	ModifiedAt           int64                    `json:"modified_at"`
	CustomerId           string                   `json:"customer_id"`
	Gateway              string                   `json:"gateway"`
	ActivePaymentAttempt *PaymentAttempt          `json:"active_payment_attempt"`
	Object               string                   `json:"object"`
}
type PaymentAttempt struct {
	Id          string                                 `json:"id"`
	Status      paymentIntentEnum.PaymentAttemptStatus `json:"status"`
	IdAtGateway string                                 `json:"id_at_gateway"`
	ErrorCode   string                                 `json:"error_code"`
	ErrorText   string                                 `json:"error_text"`
	CreatedAt   int64                                  `json:"created_at"`
	ModifiedAt  int64                                  `json:"modified_at"`
	Object      string                                 `json:"object"`
}
type CreateRequestParams struct {
	CustomerId       string `json:"customer_id,omitempty"`
	Amount           *int32 `json:"amount"`
	CurrencyCode     string `json:"currency_code"`
	GatewayAccountId string `json:"gateway_account_id,omitempty"`
	ReferenceId      string `json:"reference_id,omitempty"`
}
type UpdateRequestParams struct {
	Amount           *int32 `json:"amount,omitempty"`
	CurrencyCode     string `json:"currency_code,omitempty"`
	GatewayAccountId string `json:"gateway_account_id,omitempty"`
}
