package chargebee

type Status string

const (
	StatusValid               Status = "valid"
	StatusExpiring            Status = "expiring"
	StatusExpired             Status = "expired"
	StatusInvalid             Status = "invalid"
	StatusPendingVerification Status = "pending_verification"
)

type CardBrand string

const (
	CardBrandVisa            CardBrand = "visa"
	CardBrandMastercard      CardBrand = "mastercard"
	CardBrandAmericanExpress CardBrand = "american_express"
	CardBrandDiscover        CardBrand = "discover"
	CardBrandJcb             CardBrand = "jcb"
	CardBrandDinersClub      CardBrand = "diners_club"
	CardBrandOther           CardBrand = "other"
	CardBrandBancontact      CardBrand = "bancontact"
	CardBrandCmrFalabella    CardBrand = "cmr_falabella"
	CardBrandTarjetaNaranja  CardBrand = "tarjeta_naranja"
	CardBrandNativa          CardBrand = "nativa"
	CardBrandCencosud        CardBrand = "cencosud"
	CardBrandCabal           CardBrand = "cabal"
	CardBrandArgencard       CardBrand = "argencard"
	CardBrandElo             CardBrand = "elo"
	CardBrandHipercard       CardBrand = "hipercard"
	CardBrandCarnet          CardBrand = "carnet"
	CardBrandRupay           CardBrand = "rupay"
	CardBrandMaestro         CardBrand = "maestro"
	CardBrandDankort         CardBrand = "dankort"
	CardBrandCartesBancaires CardBrand = "cartes_bancaires"
	CardBrandNotApplicable   CardBrand = "not_applicable"
)

type CardFundingType string

const (
	CardFundingTypeCredit        CardFundingType = "credit"
	CardFundingTypeDebit         CardFundingType = "debit"
	CardFundingTypePrepaid       CardFundingType = "prepaid"
	CardFundingTypeNotKnown      CardFundingType = "not_known"
	CardFundingTypeNotApplicable CardFundingType = "not_applicable"
)

type PaymentSource struct {
	Id               string             `json:"id"`
	ResourceVersion  int64              `json:"resource_version"`
	UpdatedAt        int64              `json:"updated_at"`
	CreatedAt        int64              `json:"created_at"`
	CustomerId       string             `json:"customer_id"`
	Type             enum.Type          `json:"type"`
	ReferenceId      string             `json:"reference_id"`
	Status           Status             `json:"status"`
	Gateway          enum.Gateway       `json:"gateway"`
	GatewayAccountId string             `json:"gateway_account_id"`
	IpAddress        string             `json:"ip_address"`
	IssuingCountry   string             `json:"issuing_country"`
	Card             *Card              `json:"card"`
	BankAccount      *BankAccount       `json:"bank_account"`
	Boleto           *CustVoucherSource `json:"boleto"`
	BillingAddress   *BillingAddress    `json:"billing_address"`
	AmazonPayment    *AmazonPayment     `json:"amazon_payment"`
	Upi              *Upi               `json:"upi"`
	Paypal           *Paypal            `json:"paypal"`
	Venmo            *Venmo             `json:"venmo"`
	KlarnaPayNow     *KlarnaPayNow      `json:"klarna_pay_now"`
	Mandates         []*Mandate         `json:"mandates"`
	Deleted          bool               `json:"deleted"`
	BusinessEntityId string             `json:"business_entity_id"`
	Object           string             `json:"object"`
}
type Card struct {
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
	Iin              string          `json:"iin"`
	Last4            string          `json:"last4"`
	Brand            CardBrand       `json:"brand"`
	FundingType      CardFundingType `json:"funding_type"`
	ExpiryMonth      int32           `json:"expiry_month"`
	ExpiryYear       int32           `json:"expiry_year"`
	BillingAddr1     string          `json:"billing_addr1"`
	BillingAddr2     string          `json:"billing_addr2"`
	BillingCity      string          `json:"billing_city"`
	BillingStateCode string          `json:"billing_state_code"`
	BillingState     string          `json:"billing_state"`
	BillingCountry   string          `json:"billing_country"`
	BillingZip       string          `json:"billing_zip"`
	MaskedNumber     string          `json:"masked_number"`
	Object           string          `json:"object"`
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
type KlarnaPayNow struct {
	Email  string `json:"email"`
	Object string `json:"object"`
}
type Mandate struct {
	Id             string `json:"id"`
	SubscriptionId string `json:"subscription_id"`
	CreatedAt      int64  `json:"created_at"`
	Object         string `json:"object"`
}
type CreateUsingTempTokenRequest struct {
	CustomerId                  string                 `json:"customer_id"`
	GatewayAccountId            string                 `json:"gateway_account_id,omitempty"`
	Type                        enum.Type              `json:"type"`
	TmpToken                    string                 `json:"tmp_token"`
	IssuingCountry              string                 `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                  `json:"replace_primary_payment_source,omitempty"`
	AdditionalInformation       map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateUsingPermanentTokenRequest struct {
	CustomerId                  string                                   `json:"customer_id"`
	Type                        enum.Type                                `json:"type"`
	GatewayAccountId            string                                   `json:"gateway_account_id,omitempty"`
	ReferenceId                 string                                   `json:"reference_id,omitempty"`
	IssuingCountry              string                                   `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                                    `json:"replace_primary_payment_source,omitempty"`
	PaymentMethodToken          string                                   `json:"payment_method_token,omitempty"`
	CustomerProfileToken        string                                   `json:"customer_profile_token,omitempty"`
	NetworkTransactionId        string                                   `json:"network_transaction_id,omitempty"`
	MandateId                   string                                   `json:"mandate_id,omitempty"`
	SkipRetrieval               *bool                                    `json:"skip_retrieval,omitempty"`
	Card                        *CreateUsingPermanentTokenCard           `json:"card,omitempty"`
	BillingAddress              *CreateUsingPermanentTokenBillingAddress `json:"billing_address,omitempty"`
	AdditionalInformation       map[string]interface{}                   `json:"additional_information,omitempty"`
}
type CreateUsingPermanentTokenCard struct {
	Last4       string                        `json:"last4,omitempty"`
	Iin         string                        `json:"iin,omitempty"`
	ExpiryMonth *int32                        `json:"expiry_month,omitempty"`
	ExpiryYear  *int32                        `json:"expiry_year,omitempty"`
	Brand       paymentSource.CardBrand       `json:"brand,omitempty"`
	FundingType paymentSource.CardFundingType `json:"funding_type,omitempty"`
}
type CreateUsingPermanentTokenBillingAddress struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Line1     string `json:"line1,omitempty"`
	Line2     string `json:"line2,omitempty"`
	Line3     string `json:"line3,omitempty"`
	City      string `json:"city,omitempty"`
	StateCode string `json:"state_code,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
}
type CreateUsingTokenRequest struct {
	CustomerId                  string `json:"customer_id"`
	ReplacePrimaryPaymentSource *bool  `json:"replace_primary_payment_source,omitempty"`
	TokenId                     string `json:"token_id"`
}
type CreateUsingPaymentIntentRequest struct {
	CustomerId                  string                                 `json:"customer_id"`
	PaymentIntent               *CreateUsingPaymentIntentPaymentIntent `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                                  `json:"replace_primary_payment_source,omitempty"`
}
type CreateUsingPaymentIntentPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInfo        map[string]interface{}          `json:"additional_info,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type CreateVoucherPaymentSourceRequest struct {
	CustomerId           string                                          `json:"customer_id"`
	VoucherPaymentSource *CreateVoucherPaymentSourceVoucherPaymentSource `json:"voucher_payment_source,omitempty"`
}
type CreateVoucherPaymentSourceVoucherPaymentSource struct {
	VoucherType      enum.VoucherType       `json:"voucher_type"`
	GatewayAccountId string                 `json:"gateway_account_id,omitempty"`
	TaxId            string                 `json:"tax_id,omitempty"`
	BillingAddress   map[string]interface{} `json:"billing_address,omitempty"`
}
type CreateCardRequest struct {
	CustomerId                  string          `json:"customer_id"`
	Card                        *CreateCardCard `json:"card,omitempty"`
	ReplacePrimaryPaymentSource *bool           `json:"replace_primary_payment_source,omitempty"`
}
type CreateCardCard struct {
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number"`
	ExpiryMonth           *int32                 `json:"expiry_month"`
	ExpiryYear            *int32                 `json:"expiry_year"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       card.PreferredScheme   `json:"preferred_scheme,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateBankAccountRequest struct {
	CustomerId                  string                        `json:"customer_id"`
	BankAccount                 *CreateBankAccountBankAccount `json:"bank_account,omitempty"`
	IssuingCountry              string                        `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool                         `json:"replace_primary_payment_source,omitempty"`
}
type CreateBankAccountBankAccount struct {
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
type UpdateCardRequest struct {
	Card                 *UpdateCardCard        `json:"card,omitempty"`
	GatewayMetaData      map[string]interface{} `json:"gateway_meta_data,omitempty"`
	ReferenceTransaction string                 `json:"reference_transaction,omitempty"`
}
type UpdateCardCard struct {
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
type UpdateBankAccountRequest struct {
	BankAccount *UpdateBankAccountBankAccount `json:"bank_account,omitempty"`
}
type UpdateBankAccountBankAccount struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}
type VerifyBankAccountRequest struct {
	Amount1 *int64 `json:"amount1"`
	Amount2 *int64 `json:"amount2"`
}
type ListRequest struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	SubscriptionId string                  `json:"subscription_id,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	CustomerId     *filter.StringFilter    `json:"customer_id,omitempty"`
	Type           *filter.EnumFilter      `json:"type,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type SwitchGatewayAccountRequest struct {
	GatewayAccountId string `json:"gateway_account_id"`
}
type ExportPaymentSourceRequest struct {
	GatewayAccountId string `json:"gateway_account_id"`
}

type CreateUsingTempTokenResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type CreateUsingPermanentTokenResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type CreateUsingTokenResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type CreateUsingPaymentIntentResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type CreateVoucherPaymentSourceResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type CreateCardResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type CreateBankAccountResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type UpdateCardResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type UpdateBankAccountResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type VerifyBankAccountResponse struct {
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
}

type RetrieveResponse struct {
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
}

type ListPaymentSourceResponse struct {
	PaymentSource *PaymentSource `json:"payment_source,omitempty"`
}

type ListResponse struct {
	List       []*ListPaymentSourceResponse `json:"list,omitempty"`
	NextOffset string                       `json:"next_offset,omitempty"`
}

type SwitchGatewayAccountResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type ExportPaymentSourceResponse struct {
	ThirdPartyPaymentMethod *thirdpartypaymentmethod.ThirdPartyPaymentMethod `json:"third_party_payment_method,omitempty"`
}

type DeleteResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}

type DeleteLocalResponse struct {
	Customer      *customer.Customer `json:"customer,omitempty"`
	PaymentSource *PaymentSource     `json:"payment_source,omitempty"`
}
