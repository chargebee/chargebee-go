package chargebee

type Status string

const (
	StatusPaid       Status = "paid"
	StatusPosted     Status = "posted"
	StatusPaymentDue Status = "payment_due"
	StatusNotPaid    Status = "not_paid"
	StatusVoided     Status = "voided"
	StatusPending    Status = "pending"
)

type DunningStatus string

const (
	DunningStatusInProgress DunningStatus = "in_progress"
	DunningStatusExhausted  DunningStatus = "exhausted"
	DunningStatusStopped    DunningStatus = "stopped"
	DunningStatusSuccess    DunningStatus = "success"
)

type LineItemEntityType string

const (
	LineItemEntityTypeAdhoc           LineItemEntityType = "adhoc"
	LineItemEntityTypePlanItemPrice   LineItemEntityType = "plan_item_price"
	LineItemEntityTypeAddonItemPrice  LineItemEntityType = "addon_item_price"
	LineItemEntityTypeChargeItemPrice LineItemEntityType = "charge_item_price"
	LineItemEntityTypePlanSetup       LineItemEntityType = "plan_setup"
	LineItemEntityTypePlan            LineItemEntityType = "plan"
	LineItemEntityTypeAddon           LineItemEntityType = "addon"
)

type LineItemDiscountDiscountType string

const (
	LineItemDiscountDiscountTypeItemLevelCoupon       LineItemDiscountDiscountType = "item_level_coupon"
	LineItemDiscountDiscountTypeDocumentLevelCoupon   LineItemDiscountDiscountType = "document_level_coupon"
	LineItemDiscountDiscountTypePromotionalCredits    LineItemDiscountDiscountType = "promotional_credits"
	LineItemDiscountDiscountTypeProratedCredits       LineItemDiscountDiscountType = "prorated_credits"
	LineItemDiscountDiscountTypeItemLevelDiscount     LineItemDiscountDiscountType = "item_level_discount"
	LineItemDiscountDiscountTypeDocumentLevelDiscount LineItemDiscountDiscountType = "document_level_discount"
)

type DiscountEntityType string

const (
	DiscountEntityTypeItemLevelCoupon       DiscountEntityType = "item_level_coupon"
	DiscountEntityTypeDocumentLevelCoupon   DiscountEntityType = "document_level_coupon"
	DiscountEntityTypePromotionalCredits    DiscountEntityType = "promotional_credits"
	DiscountEntityTypeProratedCredits       DiscountEntityType = "prorated_credits"
	DiscountEntityTypeItemLevelDiscount     DiscountEntityType = "item_level_discount"
	DiscountEntityTypeDocumentLevelDiscount DiscountEntityType = "document_level_discount"
)

type DiscountDiscountType string

const (
	DiscountDiscountTypeFixedAmount DiscountDiscountType = "fixed_amount"
	DiscountDiscountTypePercentage  DiscountDiscountType = "percentage"
)

type AppliedCreditTaxApplication string

const (
	AppliedCreditTaxApplicationPreTax  AppliedCreditTaxApplication = "pre_tax"
	AppliedCreditTaxApplicationPostTax AppliedCreditTaxApplication = "post_tax"
)

type LinkedOrderStatus string

const (
	LinkedOrderStatusNew                LinkedOrderStatus = "new"
	LinkedOrderStatusProcessing         LinkedOrderStatus = "processing"
	LinkedOrderStatusComplete           LinkedOrderStatus = "complete"
	LinkedOrderStatusCancelled          LinkedOrderStatus = "cancelled"
	LinkedOrderStatusVoided             LinkedOrderStatus = "voided"
	LinkedOrderStatusQueued             LinkedOrderStatus = "queued"
	LinkedOrderStatusAwaitingShipment   LinkedOrderStatus = "awaiting_shipment"
	LinkedOrderStatusOnHold             LinkedOrderStatus = "on_hold"
	LinkedOrderStatusDelivered          LinkedOrderStatus = "delivered"
	LinkedOrderStatusShipped            LinkedOrderStatus = "shipped"
	LinkedOrderStatusPartiallyDelivered LinkedOrderStatus = "partially_delivered"
	LinkedOrderStatusReturned           LinkedOrderStatus = "returned"
)

type LinkedOrderOrderType string

const (
	LinkedOrderOrderTypeManual          LinkedOrderOrderType = "manual"
	LinkedOrderOrderTypeSystemGenerated LinkedOrderOrderType = "system_generated"
)

type NoteEntityType string

const (
	NoteEntityTypeCoupon          NoteEntityType = "coupon"
	NoteEntityTypeSubscription    NoteEntityType = "subscription"
	NoteEntityTypeCustomer        NoteEntityType = "customer"
	NoteEntityTypePlanItemPrice   NoteEntityType = "plan_item_price"
	NoteEntityTypeAddonItemPrice  NoteEntityType = "addon_item_price"
	NoteEntityTypeChargeItemPrice NoteEntityType = "charge_item_price"
	NoteEntityTypeTax             NoteEntityType = "tax"
	NoteEntityTypePlan            NoteEntityType = "plan"
	NoteEntityTypeAddon           NoteEntityType = "addon"
)

type EinvoiceStatus string

const (
	EinvoiceStatusScheduled  EinvoiceStatus = "scheduled"
	EinvoiceStatusSkipped    EinvoiceStatus = "skipped"
	EinvoiceStatusInProgress EinvoiceStatus = "in_progress"
	EinvoiceStatusSuccess    EinvoiceStatus = "success"
	EinvoiceStatusFailed     EinvoiceStatus = "failed"
	EinvoiceStatusRegistered EinvoiceStatus = "registered"
)

type Invoice struct {
	Id                        string                  `json:"id"`
	CustomerId                string                  `json:"customer_id"`
	PaymentOwner              string                  `json:"payment_owner"`
	SubscriptionId            string                  `json:"subscription_id"`
	Recurring                 bool                    `json:"recurring"`
	Status                    Status                  `json:"status"`
	Date                      int64                   `json:"date"`
	DueDate                   int64                   `json:"due_date"`
	NetTermDays               int32                   `json:"net_term_days"`
	PoNumber                  string                  `json:"po_number"`
	VatNumber                 string                  `json:"vat_number"`
	PriceType                 enum.PriceType          `json:"price_type"`
	ExchangeRate              float64                 `json:"exchange_rate"`
	LocalCurrencyExchangeRate float64                 `json:"local_currency_exchange_rate"`
	CurrencyCode              string                  `json:"currency_code"`
	LocalCurrencyCode         string                  `json:"local_currency_code"`
	Tax                       int64                   `json:"tax"`
	SubTotal                  int64                   `json:"sub_total"`
	SubTotalInLocalCurrency   int64                   `json:"sub_total_in_local_currency"`
	Total                     int64                   `json:"total"`
	TotalInLocalCurrency      int64                   `json:"total_in_local_currency"`
	AmountDue                 int64                   `json:"amount_due"`
	AmountAdjusted            int64                   `json:"amount_adjusted"`
	AmountPaid                int64                   `json:"amount_paid"`
	PaidAt                    int64                   `json:"paid_at"`
	WriteOffAmount            int64                   `json:"write_off_amount"`
	CreditsApplied            int64                   `json:"credits_applied"`
	DunningStatus             DunningStatus           `json:"dunning_status"`
	NextRetryAt               int64                   `json:"next_retry_at"`
	VoidedAt                  int64                   `json:"voided_at"`
	ResourceVersion           int64                   `json:"resource_version"`
	UpdatedAt                 int64                   `json:"updated_at"`
	LineItemsNextOffset       string                  `json:"line_items_next_offset"`
	FirstInvoice              bool                    `json:"first_invoice"`
	NewSalesAmount            int64                   `json:"new_sales_amount"`
	HasAdvanceCharges         bool                    `json:"has_advance_charges"`
	TermFinalized             bool                    `json:"term_finalized"`
	IsGifted                  bool                    `json:"is_gifted"`
	GeneratedAt               int64                   `json:"generated_at"`
	ExpectedPaymentDate       int64                   `json:"expected_payment_date"`
	AmountToCollect           int64                   `json:"amount_to_collect"`
	RoundOffAmount            int64                   `json:"round_off_amount"`
	LineItems                 []*LineItem             `json:"line_items"`
	LineItemTiers             []*LineItemTier         `json:"line_item_tiers"`
	LineItemDiscounts         []*LineItemDiscount     `json:"line_item_discounts"`
	LineItemTaxes             []*LineItemTax          `json:"line_item_taxes"`
	LineItemCredits           []*LineItemCredit       `json:"line_item_credits"`
	LineItemAddresses         []*LineItemAddress      `json:"line_item_addresses"`
	Discounts                 []*Discount             `json:"discounts"`
	Taxes                     []*Tax                  `json:"taxes"`
	TaxOrigin                 *TaxOrigin              `json:"tax_origin"`
	LinkedPayments            []*LinkedPayment        `json:"linked_payments"`
	ReferenceTransactions     []*ReferenceTransaction `json:"reference_transactions"`
	DunningAttempts           []*DunningAttempt       `json:"dunning_attempts"`
	AppliedCredits            []*AppliedCredit        `json:"applied_credits"`
	AdjustmentCreditNotes     []*AdjustmentCreditNote `json:"adjustment_credit_notes"`
	IssuedCreditNotes         []*IssuedCreditNote     `json:"issued_credit_notes"`
	LinkedOrders              []*LinkedOrder          `json:"linked_orders"`
	Notes                     []*Note                 `json:"notes"`
	ShippingAddress           *ShippingAddress        `json:"shipping_address"`
	BillingAddress            *BillingAddress         `json:"billing_address"`
	StatementDescriptor       *StatementDescriptor    `json:"statement_descriptor"`
	Einvoice                  *Einvoice               `json:"einvoice"`
	VoidReasonCode            string                  `json:"void_reason_code"`
	Deleted                   bool                    `json:"deleted"`
	TaxCategory               string                  `json:"tax_category"`
	VatNumberPrefix           string                  `json:"vat_number_prefix"`
	Channel                   enum.Channel            `json:"channel"`
	BusinessEntityId          string                  `json:"business_entity_id"`
	SiteDetailsAtCreation     *SiteDetailsAtCreation  `json:"site_details_at_creation"`
	CustomField               map[string]interface{}  `json:"custom_field"`
	Object                    string                  `json:"object"`
}
type LineItem struct {
	Id                      string               `json:"id"`
	SubscriptionId          string               `json:"subscription_id"`
	DateFrom                int64                `json:"date_from"`
	DateTo                  int64                `json:"date_to"`
	UnitAmount              int64                `json:"unit_amount"`
	Quantity                int32                `json:"quantity"`
	Amount                  int64                `json:"amount"`
	PricingModel            enum.PricingModel    `json:"pricing_model"`
	IsTaxed                 bool                 `json:"is_taxed"`
	TaxAmount               int64                `json:"tax_amount"`
	TaxRate                 float64              `json:"tax_rate"`
	UnitAmountInDecimal     string               `json:"unit_amount_in_decimal"`
	QuantityInDecimal       string               `json:"quantity_in_decimal"`
	AmountInDecimal         string               `json:"amount_in_decimal"`
	DiscountAmount          int64                `json:"discount_amount"`
	ItemLevelDiscountAmount int64                `json:"item_level_discount_amount"`
	Metered                 bool                 `json:"metered"`
	IsPercentagePricing     bool                 `json:"is_percentage_pricing"`
	ReferenceLineItemId     string               `json:"reference_line_item_id"`
	Description             string               `json:"description"`
	EntityDescription       string               `json:"entity_description"`
	EntityType              LineItemEntityType   `json:"entity_type"`
	TaxExemptReason         enum.TaxExemptReason `json:"tax_exempt_reason"`
	EntityId                string               `json:"entity_id"`
	CustomerId              string               `json:"customer_id"`
	Object                  string               `json:"object"`
}
type LineItemTier struct {
	LineItemId            string           `json:"line_item_id"`
	StartingUnit          int32            `json:"starting_unit"`
	EndingUnit            int32            `json:"ending_unit"`
	QuantityUsed          int32            `json:"quantity_used"`
	UnitAmount            int64            `json:"unit_amount"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal"`
	QuantityUsedInDecimal string           `json:"quantity_used_in_decimal"`
	UnitAmountInDecimal   string           `json:"unit_amount_in_decimal"`
	PricingType           enum.PricingType `json:"pricing_type"`
	PackageSize           int32            `json:"package_size"`
	Object                string           `json:"object"`
}
type LineItemDiscount struct {
	LineItemId     string                       `json:"line_item_id"`
	DiscountType   LineItemDiscountDiscountType `json:"discount_type"`
	CouponId       string                       `json:"coupon_id"`
	EntityId       string                       `json:"entity_id"`
	DiscountAmount int64                        `json:"discount_amount"`
	Object         string                       `json:"object"`
}
type LineItemTax struct {
	LineItemId               string            `json:"line_item_id"`
	TaxName                  string            `json:"tax_name"`
	TaxRate                  float64           `json:"tax_rate"`
	DateTo                   int64             `json:"date_to"`
	DateFrom                 int64             `json:"date_from"`
	ProratedTaxableAmount    float64           `json:"prorated_taxable_amount"`
	IsPartialTaxApplied      bool              `json:"is_partial_tax_applied"`
	IsNonComplianceTax       bool              `json:"is_non_compliance_tax"`
	TaxableAmount            int64             `json:"taxable_amount"`
	TaxAmount                int64             `json:"tax_amount"`
	TaxJurisType             enum.TaxJurisType `json:"tax_juris_type"`
	TaxJurisName             string            `json:"tax_juris_name"`
	TaxJurisCode             string            `json:"tax_juris_code"`
	TaxAmountInLocalCurrency int64             `json:"tax_amount_in_local_currency"`
	LocalCurrencyCode        string            `json:"local_currency_code"`
	Object                   string            `json:"object"`
}
type LineItemCredit struct {
	CnId          string  `json:"cn_id"`
	AppliedAmount float64 `json:"applied_amount"`
	LineItemId    string  `json:"line_item_id"`
	Object        string  `json:"object"`
}
type LineItemAddress struct {
	LineItemId       string                `json:"line_item_id"`
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
type Discount struct {
	Amount        int64                `json:"amount"`
	Description   string               `json:"description"`
	LineItemId    string               `json:"line_item_id"`
	EntityType    DiscountEntityType   `json:"entity_type"`
	DiscountType  DiscountDiscountType `json:"discount_type"`
	EntityId      string               `json:"entity_id"`
	CouponSetCode string               `json:"coupon_set_code"`
	Object        string               `json:"object"`
}
type Tax struct {
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Object      string `json:"object"`
}
type TaxOrigin struct {
	Country            string `json:"country"`
	RegistrationNumber string `json:"registration_number"`
	Object             string `json:"object"`
}
type LinkedPayment struct {
	TxnId         string             `json:"txn_id"`
	AppliedAmount int64              `json:"applied_amount"`
	AppliedAt     int64              `json:"applied_at"`
	TxnStatus     transaction.Status `json:"txn_status"`
	TxnDate       int64              `json:"txn_date"`
	TxnAmount     int64              `json:"txn_amount"`
	Object        string             `json:"object"`
}
type ReferenceTransaction struct {
	AppliedAmount       int64                           `json:"applied_amount"`
	AppliedAt           int64                           `json:"applied_at"`
	TxnId               string                          `json:"txn_id"`
	TxnStatus           transaction.Status              `json:"txn_status"`
	TxnDate             int64                           `json:"txn_date"`
	TxnAmount           int64                           `json:"txn_amount"`
	TxnType             transaction.Type                `json:"txn_type"`
	AmountCapturable    int64                           `json:"amount_capturable"`
	AuthorizationReason transaction.AuthorizationReason `json:"authorization_reason"`
	Object              string                          `json:"object"`
}
type DunningAttempt struct {
	Attempt       int32              `json:"attempt"`
	TransactionId string             `json:"transaction_id"`
	DunningType   enum.DunningType   `json:"dunning_type"`
	CreatedAt     int64              `json:"created_at"`
	TxnStatus     transaction.Status `json:"txn_status"`
	TxnAmount     int64              `json:"txn_amount"`
	Object        string             `json:"object"`
}
type AppliedCredit struct {
	CnId               string                      `json:"cn_id"`
	AppliedAmount      int64                       `json:"applied_amount"`
	AppliedAt          int64                       `json:"applied_at"`
	CnReasonCode       creditNote.ReasonCode       `json:"cn_reason_code"`
	CnCreateReasonCode string                      `json:"cn_create_reason_code"`
	CnDate             int64                       `json:"cn_date"`
	CnStatus           creditNote.Status           `json:"cn_status"`
	TaxApplication     AppliedCreditTaxApplication `json:"tax_application"`
	Object             string                      `json:"object"`
}
type AdjustmentCreditNote struct {
	CnId               string                `json:"cn_id"`
	CnReasonCode       creditNote.ReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode string                `json:"cn_create_reason_code"`
	CnDate             int64                 `json:"cn_date"`
	CnTotal            int64                 `json:"cn_total"`
	CnStatus           creditNote.Status     `json:"cn_status"`
	Object             string                `json:"object"`
}
type IssuedCreditNote struct {
	CnId               string                `json:"cn_id"`
	CnReasonCode       creditNote.ReasonCode `json:"cn_reason_code"`
	CnCreateReasonCode string                `json:"cn_create_reason_code"`
	CnDate             int64                 `json:"cn_date"`
	CnTotal            int64                 `json:"cn_total"`
	CnStatus           creditNote.Status     `json:"cn_status"`
	Object             string                `json:"object"`
}
type LinkedOrder struct {
	Id                string               `json:"id"`
	DocumentNumber    string               `json:"document_number"`
	Status            LinkedOrderStatus    `json:"status"`
	OrderType         LinkedOrderOrderType `json:"order_type"`
	ReferenceId       string               `json:"reference_id"`
	FulfillmentStatus string               `json:"fulfillment_status"`
	BatchId           string               `json:"batch_id"`
	CreatedAt         int64                `json:"created_at"`
	Object            string               `json:"object"`
}
type Note struct {
	Note       string         `json:"note"`
	EntityId   string         `json:"entity_id"`
	EntityType NoteEntityType `json:"entity_type"`
	Object     string         `json:"object"`
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
type StatementDescriptor struct {
	Id         string `json:"id"`
	Descriptor string `json:"descriptor"`
	Object     string `json:"object"`
}
type Einvoice struct {
	Id              string         `json:"id"`
	ReferenceNumber string         `json:"reference_number"`
	Status          EinvoiceStatus `json:"status"`
	Message         string         `json:"message"`
	Object          string         `json:"object"`
}
type SiteDetailsAtCreation struct {
	Timezone            string          `json:"timezone"`
	OrganizationAddress json.RawMessage `json:"organization_address"`
	Object              string          `json:"object"`
}
type CreateRequest struct {
	CustomerId                  string                     `json:"customer_id,omitempty"`
	SubscriptionId              string                     `json:"subscription_id,omitempty"`
	CurrencyCode                string                     `json:"currency_code,omitempty"`
	Addons                      []*CreateAddon             `json:"addons,omitempty"`
	InvoiceDate                 *int64                     `json:"invoice_date,omitempty"`
	Charges                     []*CreateCharge            `json:"charges,omitempty"`
	TaxProvidersFields          []*CreateTaxProvidersField `json:"tax_providers_fields,omitempty"`
	InvoiceNote                 string                     `json:"invoice_note,omitempty"`
	RemoveGeneralNote           *bool                      `json:"remove_general_note,omitempty"`
	NotesToRemove               []*CreateNotesToRemove     `json:"notes_to_remove,omitempty"`
	PoNumber                    string                     `json:"po_number,omitempty"`
	Coupon                      string                     `json:"coupon,omitempty"`
	CouponIds                   []string                   `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId  string                     `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId             string                     `json:"payment_source_id,omitempty"`
	AutoCollection              enum.AutoCollection        `json:"auto_collection,omitempty"`
	ShippingAddress             *CreateShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor         *CreateStatementDescriptor `json:"statement_descriptor,omitempty"`
	Card                        *CreateCard                `json:"card,omitempty"`
	BankAccount                 *CreateBankAccount         `json:"bank_account,omitempty"`
	TokenId                     string                     `json:"token_id,omitempty"`
	PaymentMethod               *CreatePaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent               *CreatePaymentIntent       `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                      `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                      `json:"retain_payment_source,omitempty"`
	PaymentInitiator            enum.PaymentInitiator      `json:"payment_initiator,omitempty"`
}
type CreateAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
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
type CreateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateNotesToRemove struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateShippingAddress struct {
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
type CreateStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type CreateCard struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       card.PreferredScheme   `json:"preferred_scheme,omitempty"`
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
type CreateBankAccount struct {
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
type CreatePaymentMethod struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreatePaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type CreateForChargeItemsAndChargesRequest struct {
	CustomerId                  string                                             `json:"customer_id,omitempty"`
	SubscriptionId              string                                             `json:"subscription_id,omitempty"`
	CurrencyCode                string                                             `json:"currency_code,omitempty"`
	ItemPrices                  []*CreateForChargeItemsAndChargesItemPrice         `json:"item_prices,omitempty"`
	ItemTiers                   []*CreateForChargeItemsAndChargesItemTier          `json:"item_tiers,omitempty"`
	Charges                     []*CreateForChargeItemsAndChargesCharge            `json:"charges,omitempty"`
	InvoiceNote                 string                                             `json:"invoice_note,omitempty"`
	RemoveGeneralNote           *bool                                              `json:"remove_general_note,omitempty"`
	NotesToRemove               []*CreateForChargeItemsAndChargesNotesToRemove     `json:"notes_to_remove,omitempty"`
	PoNumber                    string                                             `json:"po_number,omitempty"`
	Coupon                      string                                             `json:"coupon,omitempty"`
	CouponIds                   []string                                           `json:"coupon_ids,omitempty"`
	AuthorizationTransactionId  string                                             `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId             string                                             `json:"payment_source_id,omitempty"`
	AutoCollection              enum.AutoCollection                                `json:"auto_collection,omitempty"`
	TaxProvidersFields          []*CreateForChargeItemsAndChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
	Discounts                   []*CreateForChargeItemsAndChargesDiscount          `json:"discounts,omitempty"`
	InvoiceDate                 *int64                                             `json:"invoice_date,omitempty"`
	ShippingAddress             *CreateForChargeItemsAndChargesShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor         *CreateForChargeItemsAndChargesStatementDescriptor `json:"statement_descriptor,omitempty"`
	Card                        *CreateForChargeItemsAndChargesCard                `json:"card,omitempty"`
	BankAccount                 *CreateForChargeItemsAndChargesBankAccount         `json:"bank_account,omitempty"`
	TokenId                     string                                             `json:"token_id,omitempty"`
	PaymentMethod               *CreateForChargeItemsAndChargesPaymentMethod       `json:"payment_method,omitempty"`
	PaymentIntent               *CreateForChargeItemsAndChargesPaymentIntent       `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                                              `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                                              `json:"retain_payment_source,omitempty"`
	PaymentInitiator            enum.PaymentInitiator                              `json:"payment_initiator,omitempty"`
}
type CreateForChargeItemsAndChargesItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateForChargeItemsAndChargesItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CreateForChargeItemsAndChargesCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
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
type CreateForChargeItemsAndChargesNotesToRemove struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CreateForChargeItemsAndChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateForChargeItemsAndChargesDiscount struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Amount      *int64       `json:"amount,omitempty"`
	Quantity    *int32       `json:"quantity,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
}
type CreateForChargeItemsAndChargesShippingAddress struct {
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
type CreateForChargeItemsAndChargesStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type CreateForChargeItemsAndChargesCard struct {
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	FirstName             string                 `json:"first_name,omitempty"`
	LastName              string                 `json:"last_name,omitempty"`
	Number                string                 `json:"number,omitempty"`
	ExpiryMonth           *int32                 `json:"expiry_month,omitempty"`
	ExpiryYear            *int32                 `json:"expiry_year,omitempty"`
	Cvv                   string                 `json:"cvv,omitempty"`
	PreferredScheme       card.PreferredScheme   `json:"preferred_scheme,omitempty"`
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
type CreateForChargeItemsAndChargesBankAccount struct {
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
type CreateForChargeItemsAndChargesPaymentMethod struct {
	Type                  enum.Type              `json:"type,omitempty"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CreateForChargeItemsAndChargesPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type ChargeRequest struct {
	CustomerId             string                     `json:"customer_id,omitempty"`
	SubscriptionId         string                     `json:"subscription_id,omitempty"`
	CurrencyCode           string                     `json:"currency_code,omitempty"`
	Amount                 *int64                     `json:"amount,omitempty"`
	AmountInDecimal        string                     `json:"amount_in_decimal,omitempty"`
	Description            string                     `json:"description"`
	DateFrom               *int64                     `json:"date_from,omitempty"`
	DateTo                 *int64                     `json:"date_to,omitempty"`
	CouponIds              []string                   `json:"coupon_ids,omitempty"`
	Coupon                 string                     `json:"coupon,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType       `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32                     `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32                     `json:"avalara_service_type,omitempty"`
	PoNumber               string                     `json:"po_number,omitempty"`
	InvoiceDate            *int64                     `json:"invoice_date,omitempty"`
	TaxProvidersFields     []*ChargeTaxProvidersField `json:"tax_providers_fields,omitempty"`
	PaymentSourceId        string                     `json:"payment_source_id,omitempty"`
	PaymentInitiator       enum.PaymentInitiator      `json:"payment_initiator,omitempty"`
}
type ChargeTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type ChargeAddonRequest struct {
	CustomerId              string                `json:"customer_id,omitempty"`
	SubscriptionId          string                `json:"subscription_id,omitempty"`
	AddonId                 string                `json:"addon_id"`
	AddonQuantity           *int32                `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int64                `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string                `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string                `json:"addon_unit_price_in_decimal,omitempty"`
	DateFrom                *int64                `json:"date_from,omitempty"`
	DateTo                  *int64                `json:"date_to,omitempty"`
	CouponIds               []string              `json:"coupon_ids,omitempty"`
	Coupon                  string                `json:"coupon,omitempty"`
	PoNumber                string                `json:"po_number,omitempty"`
	InvoiceDate             *int64                `json:"invoice_date,omitempty"`
	PaymentSourceId         string                `json:"payment_source_id,omitempty"`
	PaymentInitiator        enum.PaymentInitiator `json:"payment_initiator,omitempty"`
}
type CreateForChargeItemRequest struct {
	CustomerId       string                         `json:"customer_id,omitempty"`
	SubscriptionId   string                         `json:"subscription_id,omitempty"`
	ItemPrice        *CreateForChargeItemItemPrice  `json:"item_price,omitempty"`
	ItemTiers        []*CreateForChargeItemItemTier `json:"item_tiers,omitempty"`
	PoNumber         string                         `json:"po_number,omitempty"`
	Coupon           string                         `json:"coupon,omitempty"`
	PaymentSourceId  string                         `json:"payment_source_id,omitempty"`
	PaymentInitiator enum.PaymentInitiator          `json:"payment_initiator,omitempty"`
	InvoiceDate      *int64                         `json:"invoice_date,omitempty"`
}
type CreateForChargeItemItemPrice struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type CreateForChargeItemItemTier struct {
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type StopDunningRequest struct {
	Comment string `json:"comment,omitempty"`
}
type PauseDunningRequest struct {
	ExpectedPaymentDate *int64 `json:"expected_payment_date"`
	Comment             string `json:"comment,omitempty"`
}
type ResumeDunningRequest struct {
	Comment string `json:"comment,omitempty"`
}
type ImportInvoiceRequest struct {
	Id                      string                                 `json:"id"`
	CurrencyCode            string                                 `json:"currency_code,omitempty"`
	CustomerId              string                                 `json:"customer_id,omitempty"`
	SubscriptionId          string                                 `json:"subscription_id,omitempty"`
	PoNumber                string                                 `json:"po_number,omitempty"`
	PriceType               enum.PriceType                         `json:"price_type,omitempty"`
	TaxOverrideReason       enum.TaxOverrideReason                 `json:"tax_override_reason,omitempty"`
	VatNumber               string                                 `json:"vat_number,omitempty"`
	VatNumberPrefix         string                                 `json:"vat_number_prefix,omitempty"`
	Date                    *int64                                 `json:"date"`
	Total                   *int64                                 `json:"total"`
	RoundOff                *int64                                 `json:"round_off,omitempty"`
	Status                  Status                                 `json:"status,omitempty"`
	VoidedAt                *int64                                 `json:"voided_at,omitempty"`
	VoidReasonCode          string                                 `json:"void_reason_code,omitempty"`
	IsWrittenOff            *bool                                  `json:"is_written_off,omitempty"`
	WriteOffAmount          *int64                                 `json:"write_off_amount,omitempty"`
	WriteOffDate            *int64                                 `json:"write_off_date,omitempty"`
	DueDate                 *int64                                 `json:"due_date,omitempty"`
	NetTermDays             *int32                                 `json:"net_term_days,omitempty"`
	HasAdvanceCharges       *bool                                  `json:"has_advance_charges,omitempty"`
	UseForProration         *bool                                  `json:"use_for_proration,omitempty"`
	LineItems               []*ImportInvoiceLineItem               `json:"line_items,omitempty"`
	PaymentReferenceNumbers []*ImportInvoicePaymentReferenceNumber `json:"payment_reference_numbers,omitempty"`
	LineItemTiers           []*ImportInvoiceLineItemTier           `json:"line_item_tiers,omitempty"`
	Discounts               []*ImportInvoiceDiscount               `json:"discounts,omitempty"`
	Taxes                   []*ImportInvoiceTax                    `json:"taxes,omitempty"`
	CreditNote              *ImportInvoiceCreditNote               `json:"credit_note,omitempty"`
	Payments                []*ImportInvoicePayment                `json:"payments,omitempty"`
	Notes                   []*ImportInvoiceNote                   `json:"notes,omitempty"`
	BillingAddress          *ImportInvoiceBillingAddress           `json:"billing_address,omitempty"`
	ShippingAddress         *ImportInvoiceShippingAddress          `json:"shipping_address,omitempty"`
	LineItemAddresses       []*ImportInvoiceLineItemAddress        `json:"line_item_addresses,omitempty"`
}
type ImportInvoiceLineItem struct {
	Id                         string                     `json:"id,omitempty"`
	DateFrom                   *int64                     `json:"date_from,omitempty"`
	DateTo                     *int64                     `json:"date_to,omitempty"`
	SubscriptionId             string                     `json:"subscription_id,omitempty"`
	Description                string                     `json:"description"`
	UnitAmount                 *int64                     `json:"unit_amount,omitempty"`
	Quantity                   *int32                     `json:"quantity,omitempty"`
	Amount                     *int64                     `json:"amount,omitempty"`
	UnitAmountInDecimal        string                     `json:"unit_amount_in_decimal,omitempty"`
	QuantityInDecimal          string                     `json:"quantity_in_decimal,omitempty"`
	AmountInDecimal            string                     `json:"amount_in_decimal,omitempty"`
	EntityType                 invoice.LineItemEntityType `json:"entity_type,omitempty"`
	EntityId                   string                     `json:"entity_id,omitempty"`
	ItemLevelDiscount1EntityId string                     `json:"item_level_discount1_entity_id,omitempty"`
	ItemLevelDiscount1Amount   *int64                     `json:"item_level_discount1_amount,omitempty"`
	ItemLevelDiscount2EntityId string                     `json:"item_level_discount2_entity_id,omitempty"`
	ItemLevelDiscount2Amount   *int64                     `json:"item_level_discount2_amount,omitempty"`
	Tax1Name                   string                     `json:"tax1_name,omitempty"`
	Tax1Amount                 *int64                     `json:"tax1_amount,omitempty"`
	Tax2Name                   string                     `json:"tax2_name,omitempty"`
	Tax2Amount                 *int64                     `json:"tax2_amount,omitempty"`
	Tax3Name                   string                     `json:"tax3_name,omitempty"`
	Tax3Amount                 *int64                     `json:"tax3_amount,omitempty"`
	Tax4Name                   string                     `json:"tax4_name,omitempty"`
	Tax4Amount                 *int64                     `json:"tax4_amount,omitempty"`
	Tax5Name                   string                     `json:"tax5_name,omitempty"`
	Tax5Amount                 *int64                     `json:"tax5_amount,omitempty"`
	Tax6Name                   string                     `json:"tax6_name,omitempty"`
	Tax6Amount                 *int64                     `json:"tax6_amount,omitempty"`
	Tax7Name                   string                     `json:"tax7_name,omitempty"`
	Tax7Amount                 *int64                     `json:"tax7_amount,omitempty"`
	Tax8Name                   string                     `json:"tax8_name,omitempty"`
	Tax8Amount                 *int64                     `json:"tax8_amount,omitempty"`
	Tax9Name                   string                     `json:"tax9_name,omitempty"`
	Tax9Amount                 *int64                     `json:"tax9_amount,omitempty"`
	Tax10Name                  string                     `json:"tax10_name,omitempty"`
	Tax10Amount                *int64                     `json:"tax10_amount,omitempty"`
	CreatedAt                  *int64                     `json:"created_at,omitempty"`
}
type ImportInvoicePaymentReferenceNumber struct {
	Id     string                      `json:"id,omitempty"`
	Type   paymentReferenceNumber.Type `json:"type"`
	Number string                      `json:"number"`
}
type ImportInvoiceLineItemTier struct {
	LineItemId            string `json:"line_item_id"`
	StartingUnit          *int32 `json:"starting_unit,omitempty"`
	EndingUnit            *int32 `json:"ending_unit,omitempty"`
	QuantityUsed          *int32 `json:"quantity_used,omitempty"`
	UnitAmount            *int64 `json:"unit_amount,omitempty"`
	StartingUnitInDecimal string `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string `json:"ending_unit_in_decimal,omitempty"`
	QuantityUsedInDecimal string `json:"quantity_used_in_decimal,omitempty"`
	UnitAmountInDecimal   string `json:"unit_amount_in_decimal,omitempty"`
}
type ImportInvoiceDiscount struct {
	LineItemId  string                     `json:"line_item_id,omitempty"`
	EntityType  invoice.DiscountEntityType `json:"entity_type"`
	EntityId    string                     `json:"entity_id,omitempty"`
	Description string                     `json:"description,omitempty"`
	Amount      *int64                     `json:"amount"`
}
type ImportInvoiceTax struct {
	Name        string            `json:"name"`
	Rate        *float64          `json:"rate"`
	Amount      *int64            `json:"amount,omitempty"`
	Description string            `json:"description,omitempty"`
	JurisType   enum.TaxJurisType `json:"juris_type,omitempty"`
	JurisName   string            `json:"juris_name,omitempty"`
	JurisCode   string            `json:"juris_code,omitempty"`
}
type ImportInvoiceCreditNote struct {
	Id string `json:"id,omitempty"`
}
type ImportInvoicePayment struct {
	Id              string             `json:"id,omitempty"`
	Amount          *int64             `json:"amount"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method"`
	Date            *int64             `json:"date,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
}
type ImportInvoiceNote struct {
	EntityType invoice.NoteEntityType `json:"entity_type,omitempty"`
	EntityId   string                 `json:"entity_id,omitempty"`
	Note       string                 `json:"note,omitempty"`
}
type ImportInvoiceBillingAddress struct {
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
type ImportInvoiceShippingAddress struct {
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
type ImportInvoiceLineItemAddress struct {
	LineItemId       string                `json:"line_item_id,omitempty"`
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
type ApplyPaymentsRequest struct {
	Transactions []*ApplyPaymentsTransaction `json:"transactions,omitempty"`
	Comment      string                      `json:"comment,omitempty"`
}
type ApplyPaymentsTransaction struct {
	Id     string `json:"id,omitempty"`
	Amount *int64 `json:"amount,omitempty"`
}
type DeleteLineItemsRequest struct {
	LineItems []*DeleteLineItemsLineItem `json:"line_items,omitempty"`
}
type DeleteLineItemsLineItem struct {
	Id string `json:"id,omitempty"`
}
type ApplyCreditsRequest struct {
	CreditNotes []*ApplyCreditsCreditNote `json:"credit_notes,omitempty"`
	Comment     string                    `json:"comment,omitempty"`
}
type ApplyCreditsCreditNote struct {
	Id string `json:"id,omitempty"`
}
type ListRequest struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	Einvoice       *ListEinvoice           `json:"einvoice,omitempty"`
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
	Channel        *filter.EnumFilter      `json:"channel,omitempty"`
	VoidedAt       *filter.TimestampFilter `json:"voided_at,omitempty"`
	VoidReasonCode *filter.StringFilter    `json:"void_reason_code,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type ListEinvoice struct {
	Status *filter.EnumFilter `json:"status,omitempty"`
}
type InvoicesForCustomerRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type InvoicesForSubscriptionRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type RetrieveRequest struct {
	LineItem        *RetrieveLineItem `json:"line_item,omitempty"`
	LineItemsLimit  *int32            `json:"line_items_limit,omitempty"`
	LineItemsOffset string            `json:"line_items_offset,omitempty"`
}
type RetrieveLineItem struct {
	SubscriptionId *filter.StringFilter `json:"subscription_id,omitempty"`
	CustomerId     *filter.StringFilter `json:"customer_id,omitempty"`
}
type PdfRequest struct {
	DispositionType enum.DispositionType `json:"disposition_type,omitempty"`
}
type ListPaymentReferenceNumbersRequest struct {
	Limit                  *int32                                             `json:"limit,omitempty"`
	Offset                 string                                             `json:"offset,omitempty"`
	PaymentReferenceNumber *ListPaymentReferenceNumbersPaymentReferenceNumber `json:"payment_reference_number,omitempty"`
	Id                     *filter.StringFilter                               `json:"id,omitempty"`
}
type ListPaymentReferenceNumbersPaymentReferenceNumber struct {
	Number *filter.StringFilter `json:"number,omitempty"`
}
type AddChargeRequest struct {
	Amount                 *int64               `json:"amount"`
	Description            string               `json:"description"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	AvalaraTaxCode         string               `json:"avalara_tax_code,omitempty"`
	HsnCode                string               `json:"hsn_code,omitempty"`
	TaxjarProductCode      string               `json:"taxjar_product_code,omitempty"`
	LineItem               *AddChargeLineItem   `json:"line_item,omitempty"`
	Comment                string               `json:"comment,omitempty"`
	SubscriptionId         string               `json:"subscription_id,omitempty"`
}
type AddChargeLineItem struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}
type AddAddonChargeRequest struct {
	AddonId                 string                  `json:"addon_id"`
	AddonQuantity           *int32                  `json:"addon_quantity,omitempty"`
	AddonUnitPrice          *int64                  `json:"addon_unit_price,omitempty"`
	AddonQuantityInDecimal  string                  `json:"addon_quantity_in_decimal,omitempty"`
	AddonUnitPriceInDecimal string                  `json:"addon_unit_price_in_decimal,omitempty"`
	LineItem                *AddAddonChargeLineItem `json:"line_item,omitempty"`
	Comment                 string                  `json:"comment,omitempty"`
	SubscriptionId          string                  `json:"subscription_id,omitempty"`
}
type AddAddonChargeLineItem struct {
	DateFrom *int64 `json:"date_from,omitempty"`
	DateTo   *int64 `json:"date_to,omitempty"`
}
type AddChargeItemRequest struct {
	ItemPrice      *AddChargeItemItemPrice  `json:"item_price,omitempty"`
	ItemTiers      []*AddChargeItemItemTier `json:"item_tiers,omitempty"`
	Comment        string                   `json:"comment,omitempty"`
	SubscriptionId string                   `json:"subscription_id,omitempty"`
}
type AddChargeItemItemPrice struct {
	ItemPriceId        string `json:"item_price_id"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	DateFrom           *int64 `json:"date_from,omitempty"`
	DateTo             *int64 `json:"date_to,omitempty"`
}
type AddChargeItemItemTier struct {
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
}
type CloseRequest struct {
	Comment           string                `json:"comment,omitempty"`
	InvoiceNote       string                `json:"invoice_note,omitempty"`
	RemoveGeneralNote *bool                 `json:"remove_general_note,omitempty"`
	NotesToRemove     []*CloseNotesToRemove `json:"notes_to_remove,omitempty"`
	InvoiceDate       *int64                `json:"invoice_date,omitempty"`
}
type CloseNotesToRemove struct {
	EntityType enum.EntityType `json:"entity_type,omitempty"`
	EntityId   string          `json:"entity_id,omitempty"`
}
type CollectPaymentRequest struct {
	Amount                     *int64                `json:"amount,omitempty"`
	AuthorizationTransactionId string                `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                `json:"payment_source_id,omitempty"`
	Comment                    string                `json:"comment,omitempty"`
	PaymentInitiator           enum.PaymentInitiator `json:"payment_initiator,omitempty"`
}
type RecordPaymentRequest struct {
	Transaction *RecordPaymentTransaction `json:"transaction,omitempty"`
	Comment     string                    `json:"comment,omitempty"`
}
type RecordPaymentTransaction struct {
	Amount                *int64             `json:"amount,omitempty"`
	PaymentMethod         enum.PaymentMethod `json:"payment_method"`
	ReferenceNumber       string             `json:"reference_number,omitempty"`
	CustomPaymentMethodId string             `json:"custom_payment_method_id,omitempty"`
	IdAtGateway           string             `json:"id_at_gateway,omitempty"`
	Status                transaction.Status `json:"status,omitempty"`
	Date                  *int64             `json:"date,omitempty"`
	ErrorCode             string             `json:"error_code,omitempty"`
	ErrorText             string             `json:"error_text,omitempty"`
}
type RecordTaxWithheldRequest struct {
	TaxWithheld *RecordTaxWithheldTaxWithheld `json:"tax_withheld,omitempty"`
}
type RecordTaxWithheldTaxWithheld struct {
	Amount          *int64 `json:"amount"`
	ReferenceNumber string `json:"reference_number,omitempty"`
	Date            *int64 `json:"date,omitempty"`
	Description     string `json:"description,omitempty"`
}
type RemoveTaxWithheldRequest struct {
	TaxWithheld *RemoveTaxWithheldTaxWithheld `json:"tax_withheld,omitempty"`
}
type RemoveTaxWithheldTaxWithheld struct {
	Id string `json:"id"`
}
type RefundRequest struct {
	RefundAmount  *int64            `json:"refund_amount,omitempty"`
	CreditNote    *RefundCreditNote `json:"credit_note,omitempty"`
	Comment       string            `json:"comment,omitempty"`
	CustomerNotes string            `json:"customer_notes,omitempty"`
}
type RefundCreditNote struct {
	ReasonCode       creditNote.ReasonCode `json:"reason_code,omitempty"`
	CreateReasonCode string                `json:"create_reason_code,omitempty"`
}
type RecordRefundRequest struct {
	Transaction   *RecordRefundTransaction `json:"transaction,omitempty"`
	CreditNote    *RecordRefundCreditNote  `json:"credit_note,omitempty"`
	Comment       string                   `json:"comment,omitempty"`
	CustomerNotes string                   `json:"customer_notes,omitempty"`
}
type RecordRefundTransaction struct {
	Amount                *int64             `json:"amount,omitempty"`
	PaymentMethod         enum.PaymentMethod `json:"payment_method"`
	ReferenceNumber       string             `json:"reference_number,omitempty"`
	CustomPaymentMethodId string             `json:"custom_payment_method_id,omitempty"`
	Date                  *int64             `json:"date"`
}
type RecordRefundCreditNote struct {
	ReasonCode       creditNote.ReasonCode `json:"reason_code,omitempty"`
	CreateReasonCode string                `json:"create_reason_code,omitempty"`
}
type RemovePaymentRequest struct {
	Transaction *RemovePaymentTransaction `json:"transaction,omitempty"`
}
type RemovePaymentTransaction struct {
	Id string `json:"id"`
}
type RemoveCreditNoteRequest struct {
	CreditNote *RemoveCreditNoteCreditNote `json:"credit_note,omitempty"`
}
type RemoveCreditNoteCreditNote struct {
	Id string `json:"id"`
}
type VoidInvoiceRequest struct {
	Comment        string `json:"comment,omitempty"`
	VoidReasonCode string `json:"void_reason_code,omitempty"`
}
type WriteOffRequest struct {
	Comment string `json:"comment,omitempty"`
}
type DeleteRequest struct {
	Comment string `json:"comment,omitempty"`
}
type UpdateDetailsRequest struct {
	BillingAddress      *UpdateDetailsBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress     *UpdateDetailsShippingAddress     `json:"shipping_address,omitempty"`
	StatementDescriptor *UpdateDetailsStatementDescriptor `json:"statement_descriptor,omitempty"`
	VatNumber           string                            `json:"vat_number,omitempty"`
	VatNumberPrefix     string                            `json:"vat_number_prefix,omitempty"`
	PoNumber            string                            `json:"po_number,omitempty"`
	Comment             string                            `json:"comment,omitempty"`
}
type UpdateDetailsBillingAddress struct {
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
type UpdateDetailsShippingAddress struct {
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
type UpdateDetailsStatementDescriptor struct {
	Descriptor string `json:"descriptor,omitempty"`
}
type ApplyPaymentScheduleSchemeRequest struct {
	SchemeId string `json:"scheme_id"`
	Amount   *int64 `json:"amount,omitempty"`
}

type CreateResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type CreateForChargeItemsAndChargesResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ChargeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ChargeAddonResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type CreateForChargeItemResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type StopDunningResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type PauseDunningResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ResumeDunningResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ImportInvoiceResponse struct {
	Invoice    *Invoice               `json:"invoice,omitempty"`
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type ApplyPaymentsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type SyncUsagesResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type DeleteLineItemsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ApplyCreditsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ListInvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ListResponse struct {
	List       []*ListInvoiceResponse `json:"list,omitempty"`
	NextOffset string                 `json:"next_offset,omitempty"`
}

type InvoicesForCustomerInvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type InvoicesForCustomerResponse struct {
	List       []*InvoicesForCustomerInvoiceResponse `json:"list,omitempty"`
	NextOffset string                                `json:"next_offset,omitempty"`
}

type InvoicesForSubscriptionInvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type InvoicesForSubscriptionResponse struct {
	List       []*InvoicesForSubscriptionInvoiceResponse `json:"list,omitempty"`
	NextOffset string                                    `json:"next_offset,omitempty"`
}

type RetrieveResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type PdfResponse struct {
	Download *download.Download `json:"download,omitempty"`
}

type DownloadEinvoiceResponse struct {
	Downloads []*download.Download `json:"downloads,omitempty"`
}

type ListPaymentReferenceNumbersInvoiceResponse struct {
	PaymentReferenceNumber *paymentreferencenumber.PaymentReferenceNumber `json:"payment_reference_number,omitempty"`
}

type ListPaymentReferenceNumbersResponse struct {
	List       []*ListPaymentReferenceNumbersInvoiceResponse `json:"list,omitempty"`
	NextOffset string                                        `json:"next_offset,omitempty"`
}

type AddChargeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type AddAddonChargeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type AddChargeItemResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type CloseResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type CollectPaymentResponse struct {
	Invoice     *Invoice                 `json:"invoice,omitempty"`
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type RecordPaymentResponse struct {
	Invoice     *Invoice                 `json:"invoice,omitempty"`
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type RecordTaxWithheldResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type RemoveTaxWithheldResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type RefundResponse struct {
	Invoice     *Invoice                 `json:"invoice,omitempty"`
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
	CreditNote  *creditnote.CreditNote   `json:"credit_note,omitempty"`
}

type RecordRefundResponse struct {
	Invoice     *Invoice                 `json:"invoice,omitempty"`
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
	CreditNote  *creditnote.CreditNote   `json:"credit_note,omitempty"`
}

type RemovePaymentResponse struct {
	Invoice     *Invoice                 `json:"invoice,omitempty"`
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type RemoveCreditNoteResponse struct {
	Invoice    *Invoice               `json:"invoice,omitempty"`
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type VoidInvoiceResponse struct {
	Invoice    *Invoice               `json:"invoice,omitempty"`
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type WriteOffResponse struct {
	Invoice    *Invoice               `json:"invoice,omitempty"`
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type DeleteResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type UpdateDetailsResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type ApplyPaymentScheduleSchemeResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type PaymentSchedulesResponse struct {
	PaymentSchedules []*paymentschedule.PaymentSchedule `json:"payment_schedules,omitempty"`
}

type ResendEinvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}

type SendEinvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
}
