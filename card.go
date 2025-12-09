package chargebee

type Status string

const (
	StatusValid    Status = "valid"
	StatusExpiring Status = "expiring"
	StatusExpired  Status = "expired"
)

type CardType string

const (
	CardTypeVisa            CardType = "visa"
	CardTypeMastercard      CardType = "mastercard"
	CardTypeAmericanExpress CardType = "american_express"
	CardTypeDiscover        CardType = "discover"
	CardTypeJcb             CardType = "jcb"
	CardTypeDinersClub      CardType = "diners_club"
	CardTypeBancontact      CardType = "bancontact"
	CardTypeCmrFalabella    CardType = "cmr_falabella"
	CardTypeTarjetaNaranja  CardType = "tarjeta_naranja"
	CardTypeNativa          CardType = "nativa"
	CardTypeCencosud        CardType = "cencosud"
	CardTypeCabal           CardType = "cabal"
	CardTypeArgencard       CardType = "argencard"
	CardTypeElo             CardType = "elo"
	CardTypeHipercard       CardType = "hipercard"
	CardTypeCarnet          CardType = "carnet"
	CardTypeRupay           CardType = "rupay"
	CardTypeMaestro         CardType = "maestro"
	CardTypeDankort         CardType = "dankort"
	CardTypeCartesBancaires CardType = "cartes_bancaires"
	CardTypeOther           CardType = "other"
	CardTypeNotApplicable   CardType = "not_applicable"
)

type FundingType string

const (
	FundingTypeCredit        FundingType = "credit"
	FundingTypeDebit         FundingType = "debit"
	FundingTypePrepaid       FundingType = "prepaid"
	FundingTypeNotKnown      FundingType = "not_known"
	FundingTypeNotApplicable FundingType = "not_applicable"
)

type PoweredBy string

const (
	PoweredByIdeal          PoweredBy = "ideal"
	PoweredBySofort         PoweredBy = "sofort"
	PoweredByBancontact     PoweredBy = "bancontact"
	PoweredByGiropay        PoweredBy = "giropay"
	PoweredByCard           PoweredBy = "card"
	PoweredByLatamLocalCard PoweredBy = "latam_local_card"
	PoweredByPayconiq       PoweredBy = "payconiq"
	PoweredByNotApplicable  PoweredBy = "not_applicable"
)

type PreferredScheme string

const (
	PreferredSchemeCartesBancaires PreferredScheme = "cartes_bancaires"
	PreferredSchemeMastercard      PreferredScheme = "mastercard"
	PreferredSchemeVisa            PreferredScheme = "visa"
)

type Card struct {
	PaymentSourceId  string       `json:"payment_source_id"`
	Status           Status       `json:"status"`
	Gateway          enum.Gateway `json:"gateway"`
	GatewayAccountId string       `json:"gateway_account_id"`
	RefTxId          string       `json:"ref_tx_id"`
	FirstName        string       `json:"first_name"`
	LastName         string       `json:"last_name"`
	Iin              string       `json:"iin"`
	Last4            string       `json:"last4"`
	CardType         CardType     `json:"card_type"`
	FundingType      FundingType  `json:"funding_type"`
	ExpiryMonth      int32        `json:"expiry_month"`
	ExpiryYear       int32        `json:"expiry_year"`
	IssuingCountry   string       `json:"issuing_country"`
	BillingAddr1     string       `json:"billing_addr1"`
	BillingAddr2     string       `json:"billing_addr2"`
	BillingCity      string       `json:"billing_city"`
	BillingStateCode string       `json:"billing_state_code"`
	BillingState     string       `json:"billing_state"`
	BillingCountry   string       `json:"billing_country"`
	BillingZip       string       `json:"billing_zip"`
	CreatedAt        int64        `json:"created_at"`
	ResourceVersion  int64        `json:"resource_version"`
	UpdatedAt        int64        `json:"updated_at"`
	IpAddress        string       `json:"ip_address"`
	PoweredBy        PoweredBy    `json:"powered_by"`
	CustomerId       string       `json:"customer_id"`
	MaskedNumber     string       `json:"masked_number"`
	Object           string       `json:"object"`
}
type UpdateCardForCustomerRequest struct {
	Gateway          enum.Gateway                   `json:"gateway,omitempty"`
	GatewayAccountId string                         `json:"gateway_account_id,omitempty"`
	TmpToken         string                         `json:"tmp_token,omitempty"`
	FirstName        string                         `json:"first_name,omitempty"`
	LastName         string                         `json:"last_name,omitempty"`
	Number           string                         `json:"number"`
	ExpiryMonth      *int32                         `json:"expiry_month"`
	ExpiryYear       *int32                         `json:"expiry_year"`
	Cvv              string                         `json:"cvv,omitempty"`
	PreferredScheme  PreferredScheme                `json:"preferred_scheme,omitempty"`
	BillingAddr1     string                         `json:"billing_addr1,omitempty"`
	BillingAddr2     string                         `json:"billing_addr2,omitempty"`
	BillingCity      string                         `json:"billing_city,omitempty"`
	BillingStateCode string                         `json:"billing_state_code,omitempty"`
	BillingState     string                         `json:"billing_state,omitempty"`
	BillingZip       string                         `json:"billing_zip,omitempty"`
	BillingCountry   string                         `json:"billing_country,omitempty"`
	IpAddress        string                         `json:"ip_address,omitempty"`
	Customer         *UpdateCardForCustomerCustomer `json:"customer,omitempty"`
}
type UpdateCardForCustomerCustomer struct {
	VatNumber string `json:"vat_number,omitempty"`
}
type SwitchGatewayForCustomerRequest struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id"`
}
type CopyCardForCustomerRequest struct {
	GatewayAccountId string `json:"gateway_account_id"`
}

type RetrieveResponse struct {
	Card *Card `json:"card,omitempty"`
}

type UpdateCardForCustomerResponse struct {
	Customer *customer.Customer `json:"customer,omitempty"`
	Card     *Card              `json:"card,omitempty"`
}

type SwitchGatewayForCustomerResponse struct {
	Customer *customer.Customer `json:"customer,omitempty"`
	Card     *Card              `json:"card,omitempty"`
}

type CopyCardForCustomerResponse struct {
	ThirdPartyPaymentMethod *thirdpartypaymentmethod.ThirdPartyPaymentMethod `json:"third_party_payment_method,omitempty"`
}

type DeleteCardForCustomerResponse struct {
	Customer *customer.Customer `json:"customer,omitempty"`
}
