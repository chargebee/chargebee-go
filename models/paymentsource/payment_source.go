package paymentsource

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	paymentIntentEnum "github.com/chargebee/chargebee-go/models/paymentintent/enum"
	paymentSourceEnum "github.com/chargebee/chargebee-go/models/paymentsource/enum"
)

type PaymentSource struct {
	Id               string                   `json:"id"`
	ResourceVersion  int64                    `json:"resource_version"`
	UpdatedAt        int64                    `json:"updated_at"`
	CreatedAt        int64                    `json:"created_at"`
	CustomerId       string                   `json:"customer_id"`
	Type             enum.Type                `json:"type"`
	ReferenceId      string                   `json:"reference_id"`
	Status           paymentSourceEnum.Status `json:"status"`
	Gateway          enum.Gateway             `json:"gateway"`
	GatewayAccountId string                   `json:"gateway_account_id"`
	IpAddress        string                   `json:"ip_address"`
	IssuingCountry   string                   `json:"issuing_country"`
	Card             *Card                    `json:"card"`
	BankAccount      *BankAccount             `json:"bank_account"`
	Boleto           *CustVoucherSource       `json:"boleto"`
	BillingAddress   *BillingAddress          `json:"billing_address"`
	AmazonPayment    *AmazonPayment           `json:"amazon_payment"`
	Upi              *Upi                     `json:"upi"`
	Paypal           *Paypal                  `json:"paypal"`
	Venmo            *Venmo                   `json:"venmo"`
	Mandates         []*Mandate               `json:"mandates"`
	Deleted          bool                     `json:"deleted"`
	BusinessEntityId string                   `json:"business_entity_id"`
	Object           string                   `json:"object"`
}
type Card struct {
	FirstName        string                            `json:"first_name"`
	LastName         string                            `json:"last_name"`
	Iin              string                            `json:"iin"`
	Last4            string                            `json:"last4"`
	Brand            paymentSourceEnum.CardBrand       `json:"brand"`
	FundingType      paymentSourceEnum.CardFundingType `json:"funding_type"`
	ExpiryMonth      int32                             `json:"expiry_month"`
	ExpiryYear       int32                             `json:"expiry_year"`
	BillingAddr1     string                            `json:"billing_addr1"`
	BillingAddr2     string                            `json:"billing_addr2"`
	BillingCity      string                            `json:"billing_city"`
	BillingStateCode string                            `json:"billing_state_code"`
	BillingState     string                            `json:"billing_state"`
	BillingCountry   string                            `json:"billing_country"`
	BillingZip       string                            `json:"billing_zip"`
	MaskedNumber     string                            `json:"masked_number"`
	Object           string                            `json:"object"`
}
type BankAccount struct {
	Last4             string                 `json:"last4"`
	NameOnAccount     string                 `json:"name_on_account"`
	FirstName         string                 `json:"first_name"`
	LastName          string                 `json:"last_name"`
	DirectDebitScheme enum.DirectDebitScheme `json:"direct_debit_scheme"`
	BankName          string                 `json:"bank_name"`
	MandateId         string                 `json:"mandate_id"`
	AccountType       enum.AccountType       `json:"account_type"`
	EcheckType        enum.EcheckType        `json:"echeck_type"`
	AccountHolderType enum.AccountHolderType `json:"account_holder_type"`
	Email             string                 `json:"email"`
	Object            string                 `json:"object"`
}
type CustVoucherSource struct {
	Last4     string `json:"last4"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Object    string `json:"object"`
}
type BillingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type AmazonPayment struct {
	Email       string `json:"email"`
	AgreementId string `json:"agreement_id"`
	Object      string `json:"object"`
}
type Upi struct {
	Vpa    string `json:"vpa"`
	Object string `json:"object"`
}
type Paypal struct {
	Email       string `json:"email"`
	AgreementId string `json:"agreement_id"`
	Object      string `json:"object"`
}
type Venmo struct {
	UserName string `json:"user_name"`
	Object   string `json:"object"`
}
type Mandate struct {
	Id             string `json:"id"`
	SubscriptionId string `json:"subscription_id"`
	CreatedAt      int64  `json:"created_at"`
	Object         string `json:"object"`
}
type CreateUsingTempTokenRequestParams struct {
	CustomerId                  string                 `json:"customer_id"`
	GatewayAccountId            string                 `json:"gateway_account_id,omitempty"`
	Type                        enum.Type              `json:"type"`
	TmpToken                    string                 `json:"tmp_token"`
	IssuingCountry              string                 `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                  `json:"replace_primary_payment_source,omitempty"`
	AdditionalInformation       map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateUsingPermanentTokenRequestParams struct {
	CustomerId                  string                 `json:"customer_id"`
	Type                        enum.Type              `json:"type"`
	GatewayAccountId            string                 `json:"gateway_account_id,omitempty"`
	ReferenceId                 string                 `json:"reference_id"`
	IssuingCountry              string                 `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                  `json:"replace_primary_payment_source,omitempty"`
	AdditionalInformation       map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateUsingTokenRequestParams struct {
	CustomerId                  string `json:"customer_id"`
	ReplacePrimaryPaymentSource *bool  `json:"replace_primary_payment_source,omitempty"`
	TokenId                     string `json:"token_id"`
}
type CreateUsingPaymentIntentRequestParams struct {
	CustomerId                  string                                       `json:"customer_id"`
	PaymentIntent               *CreateUsingPaymentIntentPaymentIntentParams `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                                        `json:"replace_primary_payment_source,omitempty"`
}
type CreateUsingPaymentIntentPaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInfo        map[string]interface{}              `json:"additional_info,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type CreateVoucherPaymentSourceRequestParams struct {
	CustomerId           string                                                `json:"customer_id"`
	VoucherPaymentSource *CreateVoucherPaymentSourceVoucherPaymentSourceParams `json:"voucher_payment_source,omitempty"`
}
type CreateVoucherPaymentSourceVoucherPaymentSourceParams struct {
	VoucherType      enum.VoucherType       `json:"voucher_type"`
	GatewayAccountId string                 `json:"gateway_account_id,omitempty"`
	TaxId            string                 `json:"tax_id,omitempty"`
	BillingAddress   map[string]interface{} `json:"billing_address,omitempty"`
}
type CreateCardRequestParams struct {
	CustomerId                  string                `json:"customer_id"`
	Card                        *CreateCardCardParams `json:"card,omitempty"`
	ReplacePrimaryPaymentSource *bool                 `json:"replace_primary_payment_source,omitempty"`
}
type CreateCardCardParams struct {
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number"`
	ExpiryMonth           *int32                 `json:"expiry_month"`
	ExpiryYear            *int32                 `json:"expiry_year"`
	Cvv                   string                 `json:"cvv,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateBankAccountRequestParams struct {
	CustomerId                  string                              `json:"customer_id"`
	BankAccount                 *CreateBankAccountBankAccountParams `json:"bank_account,omitempty"`
	IssuingCountry              string                              `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                               `json:"replace_primary_payment_source,omitempty"`
}
type CreateBankAccountBankAccountParams struct {
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	Iban                  string                 `json:"iban,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Company               string                 `json:"company,omitempty"`
	Email                 string                 `json:"email,omitempty"`
	Phone                 string                 `json:"phone,omitempty"`
	BankName              string                 `json:"bank_name,omitempty"`
	AccountNumber         string                 `json:"account_number,omitempty"`
	RoutingNumber         string                 `json:"routing_number,omitempty"`
	BankCode              string                 `json:"bank_code,omitempty"`
	AccountType           enum.AccountType       `json:"account_type,omitempty"`
	AccountHolderType     enum.AccountHolderType `json:"account_holder_type,omitempty"`
	EcheckType            enum.EcheckType        `json:"echeck_type,omitempty"`
	SwedishIdentityNumber string                 `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{} `json:"billing_address,omitempty"`
}
type UpdateCardRequestParams struct {
	Card                 *UpdateCardCardParams  `json:"card,omitempty"`
	GatewayMetaData      map[string]interface{} `json:"gateway_meta_data,omitempty"`
	ReferenceTransaction string                 `json:"reference_transaction,omitempty"`
}
type UpdateCardCardParams struct {
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdateBankAccountRequestParams struct {
	BankAccount *UpdateBankAccountBankAccountParams `json:"bank_account,omitempty"`
}
type UpdateBankAccountBankAccountParams struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}
type VerifyBankAccountRequestParams struct {
	Amount1 *int32 `json:"amount1"`
	Amount2 *int32 `json:"amount2"`
}
type ListRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	SubscriptionId string                  `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter    `json:"customer_id,omitempty"`
	Type           *filter.EnumFilter      `json:"type,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type SwitchGatewayAccountRequestParams struct {
	GatewayAccountId string `json:"gateway_account_id"`
}
type ExportPaymentSourceRequestParams struct {
	GatewayAccountId string `json:"gateway_account_id"`
}
