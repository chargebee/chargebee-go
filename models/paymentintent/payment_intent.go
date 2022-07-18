package paymentintent

import (
	paymentIntentEnum "github.com/chargebee/chargebee-go/models/paymentintent/enum"
)

type PaymentIntent struct {
	Id                   string                              `json:"id"`
	Status               paymentIntentEnum.Status            `json:"status"`
	CurrencyCode         string                              `json:"currency_code"`
	Amount               int32                               `json:"amount"`
	GatewayAccountId     string                              `json:"gateway_account_id"`
	ExpiresAt            int64                               `json:"expires_at"`
	ReferenceId          string                              `json:"reference_id"`
	PaymentMethodType    paymentIntentEnum.PaymentMethodType `json:"payment_method_type"`
	SuccessUrl           string                              `json:"success_url"`
	FailureUrl           string                              `json:"failure_url"`
	CreatedAt            int64                               `json:"created_at"`
	ModifiedAt           int64                               `json:"modified_at"`
	ResourceVersion      int64                               `json:"resource_version"`
	UpdatedAt            int64                               `json:"updated_at"`
	CustomerId           string                              `json:"customer_id"`
	Gateway              string                              `json:"gateway"`
	ActivePaymentAttempt *PaymentAttempt                     `json:"active_payment_attempt"`
	BusinessEntityId     string                              `json:"business_entity_id"`
	Object               string                              `json:"object"`
}
type PaymentAttempt struct {
	Id                string                                 `json:"id"`
	Status            paymentIntentEnum.PaymentAttemptStatus `json:"status"`
	PaymentMethodType paymentIntentEnum.PaymentMethodType    `json:"payment_method_type"`
	IdAtGateway       string                                 `json:"id_at_gateway"`
	ErrorCode         string                                 `json:"error_code"`
	ErrorText         string                                 `json:"error_text"`
	CreatedAt         int64                                  `json:"created_at"`
	ModifiedAt        int64                                  `json:"modified_at"`
	Object            string                                 `json:"object"`
}
type CreateRequestParams struct {
	BusinessEntityId  string                              `json:"business_entity_id,omitempty"`
	CustomerId        string                              `json:"customer_id,omitempty"`
	Amount            *int32                              `json:"amount"`
	CurrencyCode      string                              `json:"currency_code"`
	GatewayAccountId  string                              `json:"gateway_account_id,omitempty"`
	ReferenceId       string                              `json:"reference_id,omitempty"`
	PaymentMethodType paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	SuccessUrl        string                              `json:"success_url,omitempty"`
	FailureUrl        string                              `json:"failure_url,omitempty"`
}
type UpdateRequestParams struct {
	Amount            *int32                              `json:"amount,omitempty"`
	CurrencyCode      string                              `json:"currency_code,omitempty"`
	GatewayAccountId  string                              `json:"gateway_account_id,omitempty"`
	PaymentMethodType paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	SuccessUrl        string                              `json:"success_url,omitempty"`
	FailureUrl        string                              `json:"failure_url,omitempty"`
}
