package invoice

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	creditNoteEnum "github.com/chargebee/chargebee-go/models/creditnote/enum"
	invoiceEnum "github.com/chargebee/chargebee-go/models/invoice/enum"
	paymentIntentEnum "github.com/chargebee/chargebee-go/models/paymentintent/enum"
	transactionEnum "github.com/chargebee/chargebee-go/models/transaction/enum"
)

type Invoice struct {
	Id                      string                    `json:"id"`
	PoNumber                string                    `json:"po_number"`
	CustomerId              string                    `json:"customer_id"`
	SubscriptionId          string                    `json:"subscription_id"`
	Recurring               bool                      `json:"recurring"`
	Status                  invoiceEnum.Status        `json:"status"`
	VatNumber               string                    `json:"vat_number"`
	PriceType               enum.PriceType            `json:"price_type"`
	Date                    int64                     `json:"date"`
	DueDate                 int64                     `json:"due_date"`
	NetTermDays             int32                     `json:"net_term_days"`
	ExchangeRate            float64                   `json:"exchange_rate"`
	CurrencyCode            string                    `json:"currency_code"`
	Total                   int32                     `json:"total"`
	AmountPaid              int32                     `json:"amount_paid"`
	AmountAdjusted          int32                     `json:"amount_adjusted"`
	WriteOffAmount          int32                     `json:"write_off_amount"`
	CreditsApplied          int32                     `json:"credits_applied"`
	AmountDue               int32                     `json:"amount_due"`
	PaidAt                  int64                     `json:"paid_at"`
	DunningStatus           invoiceEnum.DunningStatus `json:"dunning_status"`
	NextRetryAt             int64                     `json:"next_retry_at"`
	VoidedAt                int64                     `json:"voided_at"`
	ResourceVersion         int64                     `json:"resource_version"`
	UpdatedAt               int64                     `json:"updated_at"`
	SubTotal                int32                     `json:"sub_total"`
	SubTotalInLocalCurrency int32                     `json:"sub_total_in_local_currency"`
	TotalInLocalCurrency    int32                     `json:"total_in_local_currency"`
	LocalCurrencyCode       string                    `json:"local_currency_code"`
	Tax                     int32                     `json:"tax"`
	FirstInvoice            bool                      `json:"first_invoice"`
	NewSalesAmount          int32                     `json:"new_sales_amount"`
	HasAdvanceCharges       bool                      `json:"has_advance_charges"`
	TermFinalized           bool                      `json:"term_finalized"`
	IsGifted                bool                      `json:"is_gifted"`
	GeneratedAt             int64                     `json:"generated_at"`
	ExpectedPaymentDate     int64                     `json:"expected_payment_date"`
	AmountToCollect         int32                     `json:"amount_to_collect"`
	RoundOffAmount          int32                     `json:"round_off_amount"`
	LineItems               []*LineItem               `json:"line_items"`
	Discounts               []*Discount               `json:"discounts"`
	LineItemDiscounts       []*LineItemDiscount       `json:"line_item_discounts"`
	Taxes                   []*Tax                    `json:"taxes"`
	LineItemTaxes           []*LineItemTax            `json:"line_item_taxes"`
	LineItemTiers           []*LineItemTier           `json:"line_item_tiers"`
	LinkedPayments          []*LinkedPayment          `json:"linked_payments"`
	DunningAttempts         []*DunningAttempt         `json:"dunning_attempts"`
	AppliedCredits          []*AppliedCredit          `json:"applied_credits"`
	AdjustmentCreditNotes   []*AdjustmentCreditNote   `json:"adjustment_credit_notes"`
	IssuedCreditNotes       []*IssuedCreditNote       `json:"issued_credit_notes"`
	LinkedOrders            []*LinkedOrder            `json:"linked_orders"`
	Notes                   []*Note                   `json:"notes"`
	ShippingAddress         *ShippingAddress          `json:"shipping_address"`
	BillingAddress          *BillingAddress           `json:"billing_address"`
	Einvoice                *Einvoice                 `json:"einvoice"`
	PaymentOwner            string                    `json:"payment_owner"`
	VoidReasonCode          string                    `json:"void_reason_code"`
	Deleted                 bool                      `json:"deleted"`
	VatNumberPrefix         string                    `json:"vat_number_prefix"`
	Object                  string                    `json:"object"`
}
type LineItem struct {
	Id                      string                         `json:"id"`
	SubscriptionId          string                         `json:"subscription_id"`
	DateFrom                int64                          `json:"date_from"`
	DateTo                  int64                          `json:"date_to"`
	UnitAmount              int32                          `json:"unit_amount"`
	Quantity                int32                          `json:"quantity"`
	Amount                  int32                          `json:"amount"`
	PricingModel            enum.PricingModel              `json:"pricing_model"`
	IsTaxed                 bool                           `json:"is_taxed"`
	TaxAmount               int32                          `json:"tax_amount"`
	TaxRate                 float64                        `json:"tax_rate"`
	UnitAmountInDecimal     string                         `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string                         `json:"quantity_in_decimal"`
	AmountInDecimal         string                         `json:"amount_in_decimal"`
	DiscountAmount          int32                          `json:"discount_amount"`
	ItemLevelDiscountAmount int32                          `json:"item_level_discount_amount"`
	Description             string                         `json:"description"`
	EntityDescription       string                         `json:"entity_description"`
	EntityType              invoiceEnum.LineItemEntityType `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason           `json:"tax_exempt_reason"`
	EntityId                string                         `json:"entity_id"`
	CustomerId              string                         `json:"customer_id"`
	Object                  string                         `json:"object"`
}
type Discount struct {
	Amount      int32                          `json:"amount"`
	Description string                         `json:"description"`
	EntityType  invoiceEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                         `json:"entity_id"`
	Object      string                         `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                                   `json:"line_item_id"`
	DiscountType   invoiceEnum.LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                                   `json:"coupon_id"`
	EntityId       string                                   `json:"entity_id"`
	DiscountAmount int32                                    `json:"discount_amount"`
	Object         string                                   `json:"object"`
}
type Tax struct {
	Name        string `json:"name"`
	Amount      int32  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type LineItemTax struct {
	LineItemId               string            `json:"line_item_id"`
	TaxName                  string            `json:"tax_name"`
	TaxRate                  float64           `json:"tax_rate"`
	IsPartialTaxApplied      bool              `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool              `json:"is_non_compliance_tax"`
	TaxableAmount            int32             `json:"taxable_amount"`
	TaxAmount                int32             `json:"tax_amount"`
	TaxJurisType             enum.TaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string            `json:"tax_juris_name"`
	TaxJurisCode             string            `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int32             `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string            `json:"local_currency_code"`
	Object                   string            `json:"object"`
}
type LineItemTier struct {
	LineItemId            string `json:"line_item_id"`
	StartingUnit          int32  `json:"starting_unit"`
	EndingUnit            int32  `json:"ending_unit"`
	QuantityUsed          int32  `json:"quantity_used"`
	UnitAmount            int32  `json:"unit_amount"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal"`
	Object                string `json:"object"`
}
type LinkedPayment struct {
	TxnId         string                 `json:"txn_id"`
	AppliedAmount int32                  `json:"applied_amount"`
	AppliedAt     int64                  `json:"applied_at"`
	TxnStatus     transactionEnum.Status `json:"txn_status"`
	TxnDate       int64                  `json:"txn_date"`
	TxnAmount     int32                  `json:"txn_amount"`
	Object        string                 `json:"object"`
}
type DunningAttempt struct {
	Attempt       int32                  `json:"attempt"`
	TransactionId string                 `json:"transaction_id"`
	DunningType   enum.DunningType       `json:"dunning_type"`
	CreatedAt     int64                  `json:"created_at"`
	TxnStatus     transactionEnum.Status `json:"txn_status"`
	TxnAmount     int32                  `json:"txn_amount"`
	Object        string                 `json:"object"`
}
type AppliedCredit struct {
	CnId               string                    `json:"cn_id"`
	AppliedAmount      int32                     `json:"applied_amount"`
	AppliedAt          int64                     `json:"applied_at"`
	CnReasonCode       creditNoteEnum.ReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode string                    `json:"cn_create_reason_code"`
	CnDate             int64                     `json:"cn_date"`
	CnStatus           creditNoteEnum.Status     `json:"cn_status"`
	Object             string                    `json:"object"`
}
type AdjustmentCreditNote struct {
	CnId               string                    `json:"cn_id"`
	CnReasonCode       creditNoteEnum.ReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode string                    `json:"cn_create_reason_code"`
	CnDate             int64                     `json:"cn_date"`
	CnTotal            int32                     `json:"cn_total"`
	CnStatus           creditNoteEnum.Status     `json:"cn_status"`
	Object             string                    `json:"object"`
}
type IssuedCreditNote struct {
	CnId               string                    `json:"cn_id"`
	CnReasonCode       creditNoteEnum.ReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode string                    `json:"cn_create_reason_code"`
	CnDate             int64                     `json:"cn_date"`
	CnTotal            int32                     `json:"cn_total"`
	CnStatus           creditNoteEnum.Status     `json:"cn_status"`
	Object             string                    `json:"object"`
}
type LinkedOrder struct {
	Id                string                           `json:"id"`
	DocumentNumber    string                           `json:"document_number"`
	Status            invoiceEnum.LinkedOrderStatus    `json:"status"`
	OrderType         invoiceEnum.LinkedOrderOrderType `json:"order_type"`
	ReferenceId       string                           `json:"reference_id"`
	FulfillmentStatus string                           `json:"fulfillment_status"`
	BatchId           string                           `json:"batch_id"`
	CreatedAt         int64                            `json:"created_at"`
	Object            string                           `json:"object"`
}
type Note struct {
	EntityType invoiceEnum.NoteEntityType `json:"entity_type"`
	Note       string                     `json:"note"`
	EntityId   string                     `json:"entity_id"`
	Object     string                     `json:"object"`
}
type ShippingAddress struct {
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
type Einvoice struct {
	Id      string                     `json:"id"`
	Status  invoiceEnum.EinvoiceStatus `json:"status"`
	Message string                     `json:"message"`
	Object  string                     `json:"object"`
}
type CreateRequestParams struct {
	CustomerId                  string                       `json:"customer_id,omitempty"`
	SubscriptionId              string                       `json:"subscription_id,omitempty"`
	CurrencyCode                string                       `json:"currency_code,omitempty"`
	Addons                      []*CreateAddonParams         `json:"addons,omitempty"`
	InvoiceDate                 *int64                       `json:"invoice_date,omitempty"`
	Charges                     []*CreateChargeParams        `json:"charges,omitempty"`
	InvoiceNote                 string                       `json:"invoice_note,omitempty"`
	RemoveGeneralNote           *bool                        `json:"remove_general_note,omitempty"`
	NotesToRemove               []*CreateNotesToRemoveParams `json:"notes_to_remove,omitempty"`
	PoNumber                    string                       `json:"po_number,omitempty"`
	Coupon                      string                       `json:"coupon,omitempty"`
	CouponIds                   []string                     `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId  string                       `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId             string                       `json:"payment_source_id,omitempty"`
	AutoCollection              enum.AutoCollection          `json:"auto_collection,omitempty"`
	ShippingAddress             *CreateShippingAddressParams `json:"shipping_address,omitempty"`
	Card                        *CreateCardParams            `json:"card,omitempty"`
	BankAccount                 *CreateBankAccountParams     `json:"bank_account,omitempty"`
	TokenId                     string                       `json:"token_id,omitempty"`
	PaymentMethod               *CreatePaymentMethodParams   `json:"payment_method,omitempty"`
	PaymentIntent               *CreatePaymentIntentParams   `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                        `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                        `json:"retain_payment_source,omitempty"`
}
type CreateAddonParams struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	Taxable                *bool                `json:"taxable,omitempty"`
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type CreateNotesToRemoveParams struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateCardParams struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	IpAddress             string                 `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateBankAccountParams struct {
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
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	SwedishIdentityNumber string                 `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{} `json:"billing_address,omitempty"`
}
type CreatePaymentMethodParams struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreatePaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type CreateForChargeItemsAndChargesRequestParams struct {
	CustomerId                  string                                               `json:"customer_id,omitempty"`
	SubscriptionId              string                                               `json:"subscription_id,omitempty"`
	CurrencyCode                string                                               `json:"currency_code,omitempty"`
	ItemPrices                  []*CreateForChargeItemsAndChargesItemPriceParams     `json:"item_prices,omitempty"`
	ItemTiers                   []*CreateForChargeItemsAndChargesItemTierParams      `json:"item_tiers,omitempty"`
	Charges                     []*CreateForChargeItemsAndChargesChargeParams        `json:"charges,omitempty"`
	InvoiceNote                 string                                               `json:"invoice_note,omitempty"`
	RemoveGeneralNote           *bool                                                `json:"remove_general_note,omitempty"`
	NotesToRemove               []*CreateForChargeItemsAndChargesNotesToRemoveParams `json:"notes_to_remove,omitempty"`
	PoNumber                    string                                               `json:"po_number,omitempty"`
	Coupon                      string                                               `json:"coupon,omitempty"`
	CouponIds                   []string                                             `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId  string                                               `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId             string                                               `json:"payment_source_id,omitempty"`
	AutoCollection              enum.AutoCollection                                  `json:"auto_collection,omitempty"`
	InvoiceDate                 *int64                                               `json:"invoice_date,omitempty"`
	ShippingAddress             *CreateForChargeItemsAndChargesShippingAddressParams `json:"shipping_address,omitempty"`
	Card                        *CreateForChargeItemsAndChargesCardParams            `json:"card,omitempty"`
	BankAccount                 *CreateForChargeItemsAndChargesBankAccountParams     `json:"bank_account,omitempty"`
	TokenId                     string                                               `json:"token_id,omitempty"`
	PaymentMethod               *CreateForChargeItemsAndChargesPaymentMethodParams   `json:"payment_method,omitempty"`
	PaymentIntent               *CreateForChargeItemsAndChargesPaymentIntentParams   `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                                                `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                                                `json:"retain_payment_source,omitempty"`
}
type CreateForChargeItemsAndChargesItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateForChargeItemsAndChargesItemTierParams struct {
	ItemPriceId           string `json:"item_price_id,omitempty"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CreateForChargeItemsAndChargesChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	Taxable                *bool                `json:"taxable,omitempty"`
	TaxProfileId           string               `json:"tax_profile_id,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type CreateForChargeItemsAndChargesNotesToRemoveParams struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateForChargeItemsAndChargesShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateForChargeItemsAndChargesCardParams struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	BillingAddr1          string                 `json:"billing_addr1,omitempty"`
	BillingAddr2          string                 `json:"billing_addr2,omitempty"`
	BillingCity           string                 `json:"billing_city,omitempty"`
	BillingStateCode      string                 `json:"billing_state_code,omitempty"`
	BillingState          string                 `json:"billing_state,omitempty"`
	BillingZip            string                 `json:"billing_zip,omitempty"`
	BillingCountry        string                 `json:"billing_country,omitempty"`
	IpAddress             string                 `json:"ip_address,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateForChargeItemsAndChargesBankAccountParams struct {
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
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	SwedishIdentityNumber string                 `json:"swedish_identity_number,omitempty"`
	BillingAddress        map[string]interface{} `json:"billing_address,omitempty"`
}
type CreateForChargeItemsAndChargesPaymentMethodParams struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateForChargeItemsAndChargesPaymentIntentParams struct {
	Id                    string                              `json:"id,omitempty"`
	GatewayAccountId      string                              `json:"gateway_account_id,omitempty"`
	GwToken               string                              `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntentEnum.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                              `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                              `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}              `json:"additional_information,omitempty"`
}
type ChargeRequestParams struct {
	CustomerId             string               `json:"customer_id,omitempty"`
	SubscriptionId         string               `json:"subscription_id,omitempty"`
	CurrencyCode           string               `json:"currency_code,omitempty"`
	Amount                 *int32               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
	CouponIds              []string             `json:"coupon_ids,omitempty"`
	Coupon                 string               `json:"coupon,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	PoNumber               string               `json:"po_number,omitempty"`
	InvoiceDate            *int64               `json:"invoice_date,omitempty"`
	PaymentSourceId        string               `json:"payment_source_id,omitempty"`
}
type ChargeAddonRequestParams struct {
	CustomerId              string   `json:"customer_id,omitempty"`
	SubscriptionId          string   `json:"subscription_id,omitempty"`
	AddonId                 string   `json:"addon_id"`
	AddonQuantity           *int32   `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int32   `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string   `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string   `json:"addon_unit_price_in_decimal,omitempty"`
	DateFrom                *int64   `json:"date_from,omitempty"`
	DateTo                  *int64   `json:"date_to,omitempty"`
	CouponIds               []string `json:"coupon_ids,omitempty"`
	Coupon                  string   `json:"coupon,omitempty"`
	PoNumber                string   `json:"po_number,omitempty"`
	InvoiceDate             *int64   `json:"invoice_date,omitempty"`
	PaymentSourceId         string   `json:"payment_source_id,omitempty"`
}
type CreateForChargeItemRequestParams struct {
	CustomerId      string                               `json:"customer_id,omitempty"`
	SubscriptionId  string                               `json:"subscription_id,omitempty"`
	ItemPrice       *CreateForChargeItemItemPriceParams  `json:"item_price,omitempty"`
	ItemTiers       []*CreateForChargeItemItemTierParams `json:"item_tiers,omitempty"`
	PoNumber        string                               `json:"po_number,omitempty"`
	Coupon          string                               `json:"coupon,omitempty"`
	PaymentSourceId string                               `json:"payment_source_id,omitempty"`
	InvoiceDate     *int64                               `json:"invoice_date,omitempty"`
}
type CreateForChargeItemItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateForChargeItemItemTierParams struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type StopDunningRequestParams struct {
	Comment string `json:"comment,omitempty"`
}
type ImportInvoiceRequestParams struct {
	Id                string                              `json:"id"`
	CurrencyCode      string                              `json:"currency_code,omitempty"`
	CustomerId        string                              `json:"customer_id,omitempty"`
	SubscriptionId    string                              `json:"subscription_id,omitempty"`
	PoNumber          string                              `json:"po_number,omitempty"`
	PriceType         enum.PriceType                      `json:"price_type,omitempty"`
	TaxOverrideReason enum.TaxOverrideReason              `json:"tax_override_reason,omitempty"`
	VatNumber         string                              `json:"vat_number,omitempty"`
	VatNumberPrefix   string                              `json:"vat_number_prefix,omitempty"`
	Date              *int64                              `json:"date"`
	Total             *int32                              `json:"total"`
	RoundOff          *int32                              `json:"round_off,omitempty"`
	Status            invoiceEnum.Status                  `json:"status,omitempty"`
	DueDate           *int64                              `json:"due_date,omitempty"`
	NetTermDays       *int32                              `json:"net_term_days,omitempty"`
	UseForProration   *bool                               `json:"use_for_proration,omitempty"`
	LineItems         []*ImportInvoiceLineItemParams      `json:"line_items,omitempty"`
	LineItemTiers     []*ImportInvoiceLineItemTierParams  `json:"line_item_tiers,omitempty"`
	Discounts         []*ImportInvoiceDiscountParams      `json:"discounts,omitempty"`
	Taxes             []*ImportInvoiceTaxParams           `json:"taxes,omitempty"`
	Payments          []*ImportInvoicePaymentParams       `json:"payments,omitempty"`
	Notes             []*ImportInvoiceNoteParams          `json:"notes,omitempty"`
	BillingAddress    *ImportInvoiceBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress   *ImportInvoiceShippingAddressParams `json:"shipping_address,omitempty"`
}
type ImportInvoiceLineItemParams struct {
	Id                         string                         `json:"id,omitempty"`
	DateFrom                   *int64                         `json:"date_from,omitempty"`
	DateTo                     *int64                         `json:"date_to,omitempty"`
	Description                string                         `json:"description"`
	UnitAmount                 *int32                         `json:"unit_amount,omitempty"`
	Quantity                   *int32                         `json:"quantity,omitempty"`
	Amount                     *int32                         `json:"amount,omitempty"`
	UnitAmountInDecimal        string                         `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal          string                         `json:"quantity_in_decimal,omitempty"`
	AmountInDecimal            string                         `json:"amount_in_decimal,omitempty"`
	EntityType                 invoiceEnum.LineItemEntityType `json:"entity_type,omitempty"`
	EntityId                   string                         `json:"entity_id,omitempty"`
	ItemLevelDiscount1EntityId string                         `json:"item_level_discount1_entity_id,omitempty"`
	ItemLevelDiscount1Amount   *int32                         `json:"item_level_discount1_amount,omitempty"`
	ItemLevelDiscount2EntityId string                         `json:"item_level_discount2_entity_id,omitempty"`
	ItemLevelDiscount2Amount   *int32                         `json:"item_level_discount2_amount,omitempty"`
	Tax1Name                   string                         `json:"tax1_name,omitempty"`
	Tax1Amount                 *int32                         `json:"tax1_amount,omitempty"`
	Tax2Name                   string                         `json:"tax2_name,omitempty"`
	Tax2Amount                 *int32                         `json:"tax2_amount,omitempty"`
	Tax3Name                   string                         `json:"tax3_name,omitempty"`
	Tax3Amount                 *int32                         `json:"tax3_amount,omitempty"`
	Tax4Name                   string                         `json:"tax4_name,omitempty"`
	Tax4Amount                 *int32                         `json:"tax4_amount,omitempty"`
	Tax5Name                   string                         `json:"tax5_name,omitempty"`
	Tax5Amount                 *int32                         `json:"tax5_amount,omitempty"`
	Tax6Name                   string                         `json:"tax6_name,omitempty"`
	Tax6Amount                 *int32                         `json:"tax6_amount,omitempty"`
	Tax7Name                   string                         `json:"tax7_name,omitempty"`
	Tax7Amount                 *int32                         `json:"tax7_amount,omitempty"`
	Tax8Name                   string                         `json:"tax8_name,omitempty"`
	Tax8Amount                 *int32                         `json:"tax8_amount,omitempty"`
	Tax9Name                   string                         `json:"tax9_name,omitempty"`
	Tax9Amount                 *int32                         `json:"tax9_amount,omitempty"`
	Tax10Name                  string                         `json:"tax10_name,omitempty"`
	Tax10Amount                *int32                         `json:"tax10_amount,omitempty"`
}
type ImportInvoiceLineItemTierParams struct {
	LineItemId            string `json:"line_item_id"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	QuantityUsed          *int32 `json:"quantity_used,omitempty"`
	UnitAmount            *int32 `json:"unit_amount,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal,omitempty"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal,omitempty"`
}
type ImportInvoiceDiscountParams struct {
	LineItemId  string                         `json:"line_item_id,omitempty"`
	EntityType  invoiceEnum.DiscountEntityType `json:"entity_type"`
	EntityId    string                         `json:"entity_id,omitempty"`
	Description string                         `json:"description,omitempty"`
	Amount      *int32                         `json:"amount"`
}
type ImportInvoiceTaxParams struct {
	Name        string            `json:"name"`
	Rate        *float64          `json:"rate"`
	Amount      *int32            `json:"amount,omitempty"`
	Description string            `json:"description,omitempty"`
	JurisType   enum.TaxJurisType `json:"juris_type,omitempty"`
	JurisName   string            `json:"juris_name,omitempty"`
	JurisCode   string            `json:"juris_code,omitempty"`
}
type ImportInvoicePaymentParams struct {
	Amount          *int32             `json:"amount"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method"`
	Date            *int64             `json:"date,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
}
type ImportInvoiceNoteParams struct {
	EntityType invoiceEnum.NoteEntityType `json:"entity_type,omitempty"`
	EntityId   string                     `json:"entity_id,omitempty"`
	Note       string                     `json:"note,omitempty"`
}
type ImportInvoiceBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type ImportInvoiceShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type ApplyPaymentsRequestParams struct {
	Transactions []*ApplyPaymentsTransactionParams `json:"transactions,omitempty"`
	Comment      string                            `json:"comment,omitempty"`
}
type ApplyPaymentsTransactionParams struct {
	Id string `json:"id,omitempty"`
}
type ApplyCreditsRequestParams struct {
	CreditNotes []*ApplyCreditsCreditNoteParams `json:"credit_notes,omitempty"`
	Comment     string                          `json:"comment,omitempty"`
}
type ApplyCreditsCreditNoteParams struct {
	Id string `json:"id,omitempty"`
}
type ListRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	PaidOnAfter    *int64                  `json:"paid_on_after,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	Id             *filter.StringFilter    `json:"id,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter    `json:"customer_id,omitempty"`
	Recurring      *filter.BooleanFilter   `json:"recurring,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	PriceType      *filter.EnumFilter      `json:"price_type,omitempty"`
	Date           *filter.TimestampFilter `json:"date,omitempty"`
	PaidAt         *filter.TimestampFilter `json:"paid_at,omitempty"`
	Total          *filter.NumberFilter    `json:"total,omitempty"`
	AmountPaid     *filter.NumberFilter    `json:"amount_paid,omitempty"`
	AmountAdjusted *filter.NumberFilter    `json:"amount_adjusted,omitempty"`
	CreditsApplied *filter.NumberFilter    `json:"credits_applied,omitempty"`
	AmountDue      *filter.NumberFilter    `json:"amount_due,omitempty"`
	DunningStatus  *filter.EnumFilter      `json:"dunning_status,omitempty"`
	PaymentOwner   *filter.StringFilter    `json:"payment_owner,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	VoidedAt       *filter.TimestampFilter `json:"voided_at,omitempty"`
	VoidReasonCode *filter.StringFilter    `json:"void_reason_code,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type InvoicesForCustomerRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type InvoicesForSubscriptionRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type PdfRequestParams struct {
	DispositionType enum.DispositionType `json:"disposition_type,omitempty"`
}
type AddChargeRequestParams struct {
	Amount                 *int32                   `json:"amount"`
	Description            string                   `json:"description"`
	AvalaraSaleType        enum.AvalaraSaleType     `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                   `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                   `json:"avalara_service_type,omitempty"`
	LineItem               *AddChargeLineItemParams `json:"line_item,omitempty"`
	Comment                string                   `json:"comment,omitempty"`
	SubscriptionId         string                   `json:"subscription_id,omitempty"`
}
type AddChargeLineItemParams struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}
type AddAddonChargeRequestParams struct {
	AddonId                 string                        `json:"addon_id"`
	AddonQuantity           *int32                        `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int32                        `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string                        `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string                        `json:"addon_unit_price_in_decimal,omitempty"`
	LineItem                *AddAddonChargeLineItemParams `json:"line_item,omitempty"`
	Comment                 string                        `json:"comment,omitempty"`
	SubscriptionId          string                        `json:"subscription_id,omitempty"`
}
type AddAddonChargeLineItemParams struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}
type AddChargeItemRequestParams struct {
	ItemPrice      *AddChargeItemItemPriceParams  `json:"item_price,omitempty"`
	ItemTiers      []*AddChargeItemItemTierParams `json:"item_tiers,omitempty"`
	Comment        string                         `json:"comment,omitempty"`
	SubscriptionId string                         `json:"subscription_id,omitempty"`
}
type AddChargeItemItemPriceParams struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int32 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type AddChargeItemItemTierParams struct {
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	Price                 *int32 `json:"price,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string `json:"price_in_decimal,omitempty"`
}
type CloseRequestParams struct {
	Comment           string                      `json:"comment,omitempty"`
	InvoiceNote       string                      `json:"invoice_note,omitempty"`
	RemoveGeneralNote *bool                       `json:"remove_general_note,omitempty"`
	NotesToRemove     []*CloseNotesToRemoveParams `json:"notes_to_remove,omitempty"`
	InvoiceDate       *int64                      `json:"invoice_date,omitempty"`
}
type CloseNotesToRemoveParams struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CollectPaymentRequestParams struct {
	Amount                     *int32 `json:"amount,omitempty"`
	AuthorizationTransactionId string `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string `json:"payment_source_id,omitempty"`
	Comment                    string `json:"comment,omitempty"`
}
type RecordPaymentRequestParams struct {
	Transaction *RecordPaymentTransactionParams `json:"transaction,omitempty"`
	Comment     string                          `json:"comment,omitempty"`
}
type RecordPaymentTransactionParams struct {
	Amount          *int32                 `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod     `json:"payment_method"`
	ReferenceNumber string                 `json:"reference_number,omitempty"`
	IdAtGateway     string                 `json:"id_at_gateway,omitempty"`
	Status          transactionEnum.Status `json:"status,omitempty"`
	Date            *int64                 `json:"date,omitempty"`
	ErrorCode       string                 `json:"error_code,omitempty"`
	ErrorText       string                 `json:"error_text,omitempty"`
}
type RefundRequestParams struct {
	RefundAmount  *int32                  `json:"refund_amount,omitempty"`
	CreditNote    *RefundCreditNoteParams `json:"credit_note,omitempty"`
	Comment       string                  `json:"comment,omitempty"`
	CustomerNotes string                  `json:"customer_notes,omitempty"`
}
type RefundCreditNoteParams struct {
	ReasonCode       creditNoteEnum.ReasonCode `json:"reason_code,omitempty"`
	CreateReasonCode string                    `json:"create_reason_code,omitempty"`
}
type RecordRefundRequestParams struct {
	Transaction   *RecordRefundTransactionParams `json:"transaction,omitempty"`
	CreditNote    *RecordRefundCreditNoteParams  `json:"credit_note,omitempty"`
	Comment       string                         `json:"comment,omitempty"`
	CustomerNotes string                         `json:"customer_notes,omitempty"`
}
type RecordRefundTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date"`
}
type RecordRefundCreditNoteParams struct {
	ReasonCode       creditNoteEnum.ReasonCode `json:"reason_code,omitempty"`
	CreateReasonCode string                    `json:"create_reason_code,omitempty"`
}
type RemovePaymentRequestParams struct {
	Transaction *RemovePaymentTransactionParams `json:"transaction,omitempty"`
}
type RemovePaymentTransactionParams struct {
	Id string `json:"id"`
}
type RemoveCreditNoteRequestParams struct {
	CreditNote *RemoveCreditNoteCreditNoteParams `json:"credit_note,omitempty"`
}
type RemoveCreditNoteCreditNoteParams struct {
	Id string `json:"id"`
}
type VoidInvoiceRequestParams struct {
	Comment        string `json:"comment,omitempty"`
	VoidReasonCode string `json:"void_reason_code,omitempty"`
}
type WriteOffRequestParams struct {
	Comment string `json:"comment,omitempty"`
}
type DeleteRequestParams struct {
	Comment      string `json:"comment,omitempty"`
	ClaimCredits *bool  `json:"claim_credits,omitempty"`
}
type UpdateDetailsRequestParams struct {
	BillingAddress  *UpdateDetailsBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress *UpdateDetailsShippingAddressParams `json:"shipping_address,omitempty"`
	VatNumber       string                              `json:"vat_number,omitempty"`
	VatNumberPrefix string                              `json:"vat_number_prefix,omitempty"`
	PoNumber        string                              `json:"po_number,omitempty"`
	Comment         string                              `json:"comment,omitempty"`
}
type UpdateDetailsBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateDetailsShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
