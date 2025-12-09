package chargebee

type Status string

const (
	StatusInited     Status = "inited"
	StatusInProgress Status = "in_progress"
	StatusAuthorized Status = "authorized"
	StatusConsumed   Status = "consumed"
	StatusExpired    Status = "expired"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCard                  PaymentMethodType = "card"
	PaymentMethodTypeIdeal                 PaymentMethodType = "ideal"
	PaymentMethodTypeSofort                PaymentMethodType = "sofort"
	PaymentMethodTypeBancontact            PaymentMethodType = "bancontact"
	PaymentMethodTypeGooglePay             PaymentMethodType = "google_pay"
	PaymentMethodTypeDotpay                PaymentMethodType = "dotpay"
	PaymentMethodTypeGiropay               PaymentMethodType = "giropay"
	PaymentMethodTypeApplePay              PaymentMethodType = "apple_pay"
	PaymentMethodTypeUpi                   PaymentMethodType = "upi"
	PaymentMethodTypeNetbankingEmandates   PaymentMethodType = "netbanking_emandates"
	PaymentMethodTypePaypalExpressCheckout PaymentMethodType = "paypal_express_checkout"
	PaymentMethodTypeDirectDebit           PaymentMethodType = "direct_debit"
	PaymentMethodTypeBoleto                PaymentMethodType = "boleto"
	PaymentMethodTypeVenmo                 PaymentMethodType = "venmo"
	PaymentMethodTypeAmazonPayments        PaymentMethodType = "amazon_payments"
	PaymentMethodTypePayTo                 PaymentMethodType = "pay_to"
	PaymentMethodTypeFasterPayments        PaymentMethodType = "faster_payments"
	PaymentMethodTypeSepaInstantTransfer   PaymentMethodType = "sepa_instant_transfer"
	PaymentMethodTypeKlarnaPayNow          PaymentMethodType = "klarna_pay_now"
	PaymentMethodTypeOnlineBankingPoland   PaymentMethodType = "online_banking_poland"
	PaymentMethodTypePayconiqByBancontact  PaymentMethodType = "payconiq_by_bancontact"
)

type PaymentAttemptStatus string

const (
	PaymentAttemptStatusInited                 PaymentAttemptStatus = "inited"
	PaymentAttemptStatusRequiresIdentification PaymentAttemptStatus = "requires_identification"
	PaymentAttemptStatusRequiresChallenge      PaymentAttemptStatus = "requires_challenge"
	PaymentAttemptStatusRequiresRedirection    PaymentAttemptStatus = "requires_redirection"
	PaymentAttemptStatusAuthorized             PaymentAttemptStatus = "authorized"
	PaymentAttemptStatusRefused                PaymentAttemptStatus = "refused"
	PaymentAttemptStatusPendingAuthorization   PaymentAttemptStatus = "pending_authorization"
)

type PaymentAttemptStatus string

const (
	PaymentAttemptStatusInited                 PaymentAttemptStatus = "inited"
	PaymentAttemptStatusRequiresIdentification PaymentAttemptStatus = "requires_identification"
	PaymentAttemptStatusRequiresChallenge      PaymentAttemptStatus = "requires_challenge"
	PaymentAttemptStatusRequiresRedirection    PaymentAttemptStatus = "requires_redirection"
	PaymentAttemptStatusAuthorized             PaymentAttemptStatus = "authorized"
	PaymentAttemptStatusRefused                PaymentAttemptStatus = "refused"
	PaymentAttemptStatusPendingAuthorization   PaymentAttemptStatus = "pending_authorization"
)

type PaymentIntent struct {
	Id                   string            `json:"id"`
	Status               Status            `json:"status"`
	CurrencyCode         string            `json:"currency_code"`
	Amount               int64             `json:"amount"`
	GatewayAccountId     string            `json:"gateway_account_id"`
	ExpiresAt            int64             `json:"expires_at"`
	ReferenceId          string            `json:"reference_id"`
	PaymentMethodType    PaymentMethodType `json:"payment_method_type"`
	SuccessUrl           string            `json:"success_url"`
	FailureUrl           string            `json:"failure_url"`
	CreatedAt            int64             `json:"created_at"`
	ModifiedAt           int64             `json:"modified_at"`
	ResourceVersion      int64             `json:"resource_version"`
	UpdatedAt            int64             `json:"updated_at"`
	CustomerId           string            `json:"customer_id"`
	Gateway              string            `json:"gateway"`
	ActivePaymentAttempt *PaymentAttempt   `json:"active_payment_attempt"`
	PaymentAttempts      []*PaymentAttempt `json:"payment_attempts"`
	BusinessEntityId     string            `json:"business_entity_id"`
	Object               string            `json:"object"`
}
type PaymentAttempt struct {
	Id                string                                 `json:"id"`
	Status            PaymentAttemptStatus                   `json:"status"`
	PaymentMethodType paymentIntentEnum.PaymentMethodType    `json:"payment_method_type"`
	IdAtGateway       string                                 `json:"id_at_gateway"`
	ErrorCode         string                                 `json:"error_code"`
	ErrorText         string                                 `json:"error_text"`
	CheckoutDetails   string                                 `json:"checkout_details"`
	CreatedAt         int64                                  `json:"created_at"`
	ModifiedAt        int64                                  `json:"modified_at"`
	ErrorDetail       *gatewayerrordetail.GatewayErrorDetail `json:"error_detail"`
	Object            string                                 `json:"object"`
}
type CreateRequest struct {
	BusinessEntityId  string            `json:"business_entity_id,omitempty"`
	CustomerId        string            `json:"customer_id,omitempty"`
	Amount            *int64            `json:"amount"`
	CurrencyCode      string            `json:"currency_code"`
	GatewayAccountId  string            `json:"gateway_account_id,omitempty"`
	ReferenceId       string            `json:"reference_id,omitempty"`
	PaymentMethodType PaymentMethodType `json:"payment_method_type,omitempty"`
	SuccessUrl        string            `json:"success_url,omitempty"`
	FailureUrl        string            `json:"failure_url,omitempty"`
}
type UpdateRequest struct {
	Amount            *int64            `json:"amount,omitempty"`
	CurrencyCode      string            `json:"currency_code,omitempty"`
	GatewayAccountId  string            `json:"gateway_account_id,omitempty"`
	PaymentMethodType PaymentMethodType `json:"payment_method_type,omitempty"`
	SuccessUrl        string            `json:"success_url,omitempty"`
	FailureUrl        string            `json:"failure_url,omitempty"`
}

type CreateResponse struct {
	PaymentIntent *PaymentIntent `json:"payment_intent,omitempty"`
}

type UpdateResponse struct {
	PaymentIntent *PaymentIntent `json:"payment_intent,omitempty"`
}

type RetrieveResponse struct {
	PaymentIntent *PaymentIntent `json:"payment_intent,omitempty"`
}
