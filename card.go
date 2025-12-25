package chargebee

type CardStatus string

const (
	CardStatusValid    CardStatus = "valid"
	CardStatusExpiring CardStatus = "expiring"
	CardStatusExpired  CardStatus = "expired"
)

type CardGateway string

const (
	CardGatewayChargebee             CardGateway = "chargebee"
	CardGatewayChargebeePayments     CardGateway = "chargebee_payments"
	CardGatewayAdyen                 CardGateway = "adyen"
	CardGatewayStripe                CardGateway = "stripe"
	CardGatewayWepay                 CardGateway = "wepay"
	CardGatewayBraintree             CardGateway = "braintree"
	CardGatewayAuthorizeNet          CardGateway = "authorize_net"
	CardGatewayPaypalPro             CardGateway = "paypal_pro"
	CardGatewayPin                   CardGateway = "pin"
	CardGatewayEway                  CardGateway = "eway"
	CardGatewayEwayRapid             CardGateway = "eway_rapid"
	CardGatewayWorldpay              CardGateway = "worldpay"
	CardGatewayBalancedPayments      CardGateway = "balanced_payments"
	CardGatewayBeanstream            CardGateway = "beanstream"
	CardGatewayBluepay               CardGateway = "bluepay"
	CardGatewayElavon                CardGateway = "elavon"
	CardGatewayFirstDataGlobal       CardGateway = "first_data_global"
	CardGatewayHdfc                  CardGateway = "hdfc"
	CardGatewayMigs                  CardGateway = "migs"
	CardGatewayNmi                   CardGateway = "nmi"
	CardGatewayOgone                 CardGateway = "ogone"
	CardGatewayPaymill               CardGateway = "paymill"
	CardGatewayPaypalPayflowPro      CardGateway = "paypal_payflow_pro"
	CardGatewaySagePay               CardGateway = "sage_pay"
	CardGatewayTco                   CardGateway = "tco"
	CardGatewayWirecard              CardGateway = "wirecard"
	CardGatewayAmazonPayments        CardGateway = "amazon_payments"
	CardGatewayPaypalExpressCheckout CardGateway = "paypal_express_checkout"
	CardGatewayGocardless            CardGateway = "gocardless"
	CardGatewayOrbital               CardGateway = "orbital"
	CardGatewayMonerisUs             CardGateway = "moneris_us"
	CardGatewayMoneris               CardGateway = "moneris"
	CardGatewayBluesnap              CardGateway = "bluesnap"
	CardGatewayCybersource           CardGateway = "cybersource"
	CardGatewayVantiv                CardGateway = "vantiv"
	CardGatewayCheckoutCom           CardGateway = "checkout_com"
	CardGatewayPaypal                CardGateway = "paypal"
	CardGatewayIngenicoDirect        CardGateway = "ingenico_direct"
	CardGatewayExact                 CardGateway = "exact"
	CardGatewayMollie                CardGateway = "mollie"
	CardGatewayQuickbooks            CardGateway = "quickbooks"
	CardGatewayRazorpay              CardGateway = "razorpay"
	CardGatewayGlobalPayments        CardGateway = "global_payments"
	CardGatewayBankOfAmerica         CardGateway = "bank_of_america"
	CardGatewayEcentric              CardGateway = "ecentric"
	CardGatewayMetricsGlobal         CardGateway = "metrics_global"
	CardGatewayWindcave              CardGateway = "windcave"
	CardGatewayPayCom                CardGateway = "pay_com"
	CardGatewayEbanx                 CardGateway = "ebanx"
	CardGatewayDlocal                CardGateway = "dlocal"
	CardGatewayNuvei                 CardGateway = "nuvei"
	CardGatewaySolidgate             CardGateway = "solidgate"
	CardGatewayPaystack              CardGateway = "paystack"
	CardGatewayJpMorgan              CardGateway = "jp_morgan"
	CardGatewayDeutscheBank          CardGateway = "deutsche_bank"
	CardGatewayEzidebit              CardGateway = "ezidebit"
	CardGatewayNotApplicable         CardGateway = "not_applicable"
)

type CardCardType string

const (
	CardCardTypeVisa            CardCardType = "visa"
	CardCardTypeMastercard      CardCardType = "mastercard"
	CardCardTypeAmericanExpress CardCardType = "american_express"
	CardCardTypeDiscover        CardCardType = "discover"
	CardCardTypeJcb             CardCardType = "jcb"
	CardCardTypeDinersClub      CardCardType = "diners_club"
	CardCardTypeBancontact      CardCardType = "bancontact"
	CardCardTypeCmrFalabella    CardCardType = "cmr_falabella"
	CardCardTypeTarjetaNaranja  CardCardType = "tarjeta_naranja"
	CardCardTypeNativa          CardCardType = "nativa"
	CardCardTypeCencosud        CardCardType = "cencosud"
	CardCardTypeCabal           CardCardType = "cabal"
	CardCardTypeArgencard       CardCardType = "argencard"
	CardCardTypeElo             CardCardType = "elo"
	CardCardTypeHipercard       CardCardType = "hipercard"
	CardCardTypeCarnet          CardCardType = "carnet"
	CardCardTypeRupay           CardCardType = "rupay"
	CardCardTypeMaestro         CardCardType = "maestro"
	CardCardTypeDankort         CardCardType = "dankort"
	CardCardTypeCartesBancaires CardCardType = "cartes_bancaires"
	CardCardTypeOther           CardCardType = "other"
	CardCardTypeNotApplicable   CardCardType = "not_applicable"
)

type CardFundingType string

const (
	CardFundingTypeCredit        CardFundingType = "credit"
	CardFundingTypeDebit         CardFundingType = "debit"
	CardFundingTypePrepaid       CardFundingType = "prepaid"
	CardFundingTypeNotKnown      CardFundingType = "not_known"
	CardFundingTypeNotApplicable CardFundingType = "not_applicable"
)

type CardPoweredBy string

const (
	CardPoweredByIdeal          CardPoweredBy = "ideal"
	CardPoweredBySofort         CardPoweredBy = "sofort"
	CardPoweredByBancontact     CardPoweredBy = "bancontact"
	CardPoweredByGiropay        CardPoweredBy = "giropay"
	CardPoweredByCard           CardPoweredBy = "card"
	CardPoweredByLatamLocalCard CardPoweredBy = "latam_local_card"
	CardPoweredByPayconiq       CardPoweredBy = "payconiq"
	CardPoweredByNotApplicable  CardPoweredBy = "not_applicable"
)

type CardPreferredScheme string

const (
	CardPreferredSchemeCartesBancaires CardPreferredScheme = "cartes_bancaires"
	CardPreferredSchemeMastercard      CardPreferredScheme = "mastercard"
	CardPreferredSchemeVisa            CardPreferredScheme = "visa"
)

// just struct
type Card struct {
	PaymentSourceId  string          `json:"payment_source_id"`
	Status           CardStatus      `json:"status"`
	Gateway          CardGateway     `json:"gateway"`
	GatewayAccountId string          `json:"gateway_account_id"`
	RefTxId          string          `json:"ref_tx_id"`
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
	Iin              string          `json:"iin"`
	Last4            string          `json:"last4"`
	CardType         CardCardType    `json:"card_type"`
	FundingType      CardFundingType `json:"funding_type"`
	ExpiryMonth      int32           `json:"expiry_month"`
	ExpiryYear       int32           `json:"expiry_year"`
	IssuingCountry   string          `json:"issuing_country"`
	BillingAddr1     string          `json:"billing_addr1"`
	BillingAddr2     string          `json:"billing_addr2"`
	BillingCity      string          `json:"billing_city"`
	BillingStateCode string          `json:"billing_state_code"`
	BillingState     string          `json:"billing_state"`
	BillingCountry   string          `json:"billing_country"`
	BillingZip       string          `json:"billing_zip"`
	CreatedAt        int64           `json:"created_at"`
	ResourceVersion  int64           `json:"resource_version"`
	UpdatedAt        int64           `json:"updated_at"`
	IpAddress        string          `json:"ip_address"`
	PoweredBy        CardPoweredBy   `json:"powered_by"`
	CustomerId       string          `json:"customer_id"`
	MaskedNumber     string          `json:"masked_number"`
	Object           string          `json:"object"`
}

// sub resources
// operations
// input params
type CardUpdateCardForCustomerRequest struct {
	GatewayAccountId string                             `json:"gateway_account_id,omitempty"`
	TmpToken         string                             `json:"tmp_token,omitempty"`
	FirstName        string                             `json:"first_name,omitempty"`
	LastName         string                             `json:"last_name,omitempty"`
	Number           string                             `json:"number"`
	ExpiryMonth      *int32                             `json:"expiry_month"`
	ExpiryYear       *int32                             `json:"expiry_year"`
	Cvv              string                             `json:"cvv,omitempty"`
	PreferredScheme  CardPreferredScheme                `json:"preferred_scheme,omitempty"`
	BillingAddr1     string                             `json:"billing_addr1,omitempty"`
	BillingAddr2     string                             `json:"billing_addr2,omitempty"`
	BillingCity      string                             `json:"billing_city,omitempty"`
	BillingStateCode string                             `json:"billing_state_code,omitempty"`
	BillingState     string                             `json:"billing_state,omitempty"`
	BillingZip       string                             `json:"billing_zip,omitempty"`
	BillingCountry   string                             `json:"billing_country,omitempty"`
	Customer         *CardUpdateCardForCustomerCustomer `json:"customer,omitempty"`
	apiRequest       `json:"-" form:"-"`
}

func (r *CardUpdateCardForCustomerRequest) payload() any { return r }

// input sub resource params single
type CardUpdateCardForCustomerCustomer struct {
	VatNumber string `json:"vat_number,omitempty"`
}

type CardSwitchGatewayForCustomerRequest struct {
	GatewayAccountId string `json:"gateway_account_id"`
	apiRequest       `json:"-" form:"-"`
}

func (r *CardSwitchGatewayForCustomerRequest) payload() any { return r }

type CardCopyCardForCustomerRequest struct {
	GatewayAccountId string `json:"gateway_account_id"`
	apiRequest       `json:"-" form:"-"`
}

func (r *CardCopyCardForCustomerRequest) payload() any { return r }

// operation response
type CardRetrieveResponse struct {
	Card *Card `json:"card,omitempty"`
	apiResponse
}

// operation response
type CardUpdateCardForCustomerResponse struct {
	Customer Customer `json:"customer,omitempty"`
	Card     *Card    `json:"card,omitempty"`
	apiResponse
}

// operation response
type CardSwitchGatewayForCustomerResponse struct {
	Customer Customer `json:"customer,omitempty"`
	Card     *Card    `json:"card,omitempty"`
	apiResponse
}

// operation response
type CardCopyCardForCustomerResponse struct {
	ThirdPartyPaymentMethod ThirdPartyPaymentMethod `json:"third_party_payment_method,omitempty"`
	apiResponse
}

// operation response
type CardDeleteCardForCustomerResponse struct {
	Customer Customer `json:"customer,omitempty"`
	apiResponse
}
