package chargebee

type Status string

const (
	StatusOpen             Status = "open"
	StatusAccepted         Status = "accepted"
	StatusDeclined         Status = "declined"
	StatusInvoiced         Status = "invoiced"
	StatusClosed           Status = "closed"
	StatusPendingApproval  Status = "pending_approval"
	StatusApprovalRejected Status = "approval_rejected"
	StatusProposed         Status = "proposed"
	StatusVoided           Status = "voided"
	StatusExpired          Status = "expired"
)

type OperationType string

const (
	OperationTypeCreateSubscriptionForCustomer OperationType = "create_subscription_for_customer"
	OperationTypeChangeSubscription            OperationType = "change_subscription"
	OperationTypeOnetimeInvoice                OperationType = "onetime_invoice"
	OperationTypeRenewSubscription             OperationType = "renew_subscription"
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

type Quote struct {
	Id                         string                 `json:"id"`
	Name                       string                 `json:"name"`
	PoNumber                   string                 `json:"po_number"`
	CustomerId                 string                 `json:"customer_id"`
	SubscriptionId             string                 `json:"subscription_id"`
	InvoiceId                  string                 `json:"invoice_id"`
	Status                     Status                 `json:"status"`
	OperationType              OperationType          `json:"operation_type"`
	VatNumber                  string                 `json:"vat_number"`
	PriceType                  enum.PriceType         `json:"price_type"`
	ValidTill                  int64                  `json:"valid_till"`
	Date                       int64                  `json:"date"`
	TotalPayable               int64                  `json:"total_payable"`
	ChargeOnAcceptance         int64                  `json:"charge_on_acceptance"`
	SubTotal                   int64                  `json:"sub_total"`
	Total                      int64                  `json:"total"`
	CreditsApplied             int64                  `json:"credits_applied"`
	AmountPaid                 int64                  `json:"amount_paid"`
	AmountDue                  int64                  `json:"amount_due"`
	Version                    int32                  `json:"version"`
	ResourceVersion            int64                  `json:"resource_version"`
	UpdatedAt                  int64                  `json:"updated_at"`
	VatNumberPrefix            string                 `json:"vat_number_prefix"`
	LineItems                  []*LineItem            `json:"line_items"`
	LineItemTiers              []*LineItemTier        `json:"line_item_tiers"`
	LineItemDiscounts          []*LineItemDiscount    `json:"line_item_discounts"`
	LineItemTaxes              []*LineItemTax         `json:"line_item_taxes"`
	Discounts                  []*Discount            `json:"discounts"`
	Taxes                      []*Tax                 `json:"taxes"`
	TaxCategory                string                 `json:"tax_category"`
	CurrencyCode               string                 `json:"currency_code"`
	Notes                      json.RawMessage        `json:"notes"`
	ShippingAddress            *ShippingAddress       `json:"shipping_address"`
	BillingAddress             *BillingAddress        `json:"billing_address"`
	ContractTermStart          int64                  `json:"contract_term_start"`
	ContractTermEnd            int64                  `json:"contract_term_end"`
	ContractTermTerminationFee int64                  `json:"contract_term_termination_fee"`
	BusinessEntityId           string                 `json:"business_entity_id"`
	Deleted                    bool                   `json:"deleted"`
	TotalContractValue         int64                  `json:"total_contract_value"`
	TotalDiscount              int64                  `json:"total_discount"`
	CustomField                map[string]interface{} `json:"custom_field"`
	Object                     string                 `json:"object"`
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
type CreateSubForCustomerQuoteRequest struct {
	Name                    string                                      `json:"name,omitempty"`
	Notes                   string                                      `json:"notes,omitempty"`
	ExpiresAt               *int64                                      `json:"expires_at,omitempty"`
	Subscription            *CreateSubForCustomerQuoteSubscription      `json:"subscription,omitempty"`
	BillingCycles           *int32                                      `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubForCustomerQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubForCustomerQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                    `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                      `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                   `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *CreateSubForCustomerQuoteShippingAddress   `json:"shipping_address,omitempty"`
	ContractTerm            *CreateSubForCustomerQuoteContractTerm      `json:"contract_term,omitempty"`
	CouponIds               []string                                    `json:"coupon_ids,omitempty"`
}
type CreateSubForCustomerQuoteSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CreateSubForCustomerQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerQuoteEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CreateSubForCustomerQuoteShippingAddress struct {
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
type CreateSubForCustomerQuoteContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type EditCreateSubForCustomerQuoteRequest struct {
	Notes                   string                                          `json:"notes,omitempty"`
	ExpiresAt               *int64                                          `json:"expires_at,omitempty"`
	Subscription            *EditCreateSubForCustomerQuoteSubscription      `json:"subscription,omitempty"`
	BillingCycles           *int32                                          `json:"billing_cycles,omitempty"`
	Addons                  []*EditCreateSubForCustomerQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*EditCreateSubForCustomerQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	MandatoryAddonsToRemove []string                                        `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                          `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                       `json:"billing_alignment_mode,omitempty"`
	ShippingAddress         *EditCreateSubForCustomerQuoteShippingAddress   `json:"shipping_address,omitempty"`
	ContractTerm            *EditCreateSubForCustomerQuoteContractTerm      `json:"contract_term,omitempty"`
	CouponIds               []string                                        `json:"coupon_ids,omitempty"`
}
type EditCreateSubForCustomerQuoteSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	PlanId                            string                    `json:"plan_id"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditCreateSubForCustomerQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type EditCreateSubForCustomerQuoteEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type EditCreateSubForCustomerQuoteShippingAddress struct {
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
type EditCreateSubForCustomerQuoteContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateSubscriptionQuoteRequest struct {
	Name                    string                                    `json:"name,omitempty"`
	Notes                   string                                    `json:"notes,omitempty"`
	ExpiresAt               *int64                                    `json:"expires_at,omitempty"`
	Subscription            *UpdateSubscriptionQuoteSubscription      `json:"subscription,omitempty"`
	Addons                  []*UpdateSubscriptionQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*UpdateSubscriptionQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                     `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                  `json:"mandatory_addons_to_remove,omitempty"`
	BillingCycles           *int32                                    `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                    `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                    `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                 `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                  `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                     `json:"replace_coupon_list,omitempty"`
	ChangeOption            enum.ChangeOption                         `json:"change_option,omitempty"`
	ChangesScheduledAt      *int64                                    `json:"changes_scheduled_at,omitempty"`
	ForceTermReset          *bool                                     `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                     `json:"reactivate,omitempty"`
	BillingAddress          *UpdateSubscriptionQuoteBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress         *UpdateSubscriptionQuoteShippingAddress   `json:"shipping_address,omitempty"`
	Customer                *UpdateSubscriptionQuoteCustomer          `json:"customer,omitempty"`
	ContractTerm            *UpdateSubscriptionQuoteContractTerm      `json:"contract_term,omitempty"`
}
type UpdateSubscriptionQuoteSubscription struct {
	Id                                string                    `json:"id"`
	PlanId                            string                    `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type UpdateSubscriptionQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type UpdateSubscriptionQuoteEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
}
type UpdateSubscriptionQuoteBillingAddress struct {
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
type UpdateSubscriptionQuoteShippingAddress struct {
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
type UpdateSubscriptionQuoteCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type UpdateSubscriptionQuoteContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type EditUpdateSubscriptionQuoteRequest struct {
	Notes                   string                                        `json:"notes,omitempty"`
	ExpiresAt               *int64                                        `json:"expires_at,omitempty"`
	Subscription            *EditUpdateSubscriptionQuoteSubscription      `json:"subscription,omitempty"`
	Addons                  []*EditUpdateSubscriptionQuoteAddon           `json:"addons,omitempty"`
	EventBasedAddons        []*EditUpdateSubscriptionQuoteEventBasedAddon `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                         `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                      `json:"mandatory_addons_to_remove,omitempty"`
	BillingCycles           *int32                                        `json:"billing_cycles,omitempty"`
	TermsToCharge           *int32                                        `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                        `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                     `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                      `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                         `json:"replace_coupon_list,omitempty"`
	ChangeOption            enum.ChangeOption                             `json:"change_option,omitempty"`
	ChangesScheduledAt      *int64                                        `json:"changes_scheduled_at,omitempty"`
	ForceTermReset          *bool                                         `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                         `json:"reactivate,omitempty"`
	BillingAddress          *EditUpdateSubscriptionQuoteBillingAddress    `json:"billing_address,omitempty"`
	ShippingAddress         *EditUpdateSubscriptionQuoteShippingAddress   `json:"shipping_address,omitempty"`
	Customer                *EditUpdateSubscriptionQuoteCustomer          `json:"customer,omitempty"`
	ContractTerm            *EditUpdateSubscriptionQuoteContractTerm      `json:"contract_term,omitempty"`
}
type EditUpdateSubscriptionQuoteSubscription struct {
	PlanId                            string                    `json:"plan_id,omitempty"`
	PlanQuantity                      *int32                    `json:"plan_quantity,omitempty"`
	PlanUnitPrice                     *int64                    `json:"plan_unit_price,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	PlanQuantityInDecimal             string                    `json:"plan_quantity_in_decimal,omitempty"`
	PlanUnitPriceInDecimal            string                    `json:"plan_unit_price_in_decimal,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditUpdateSubscriptionQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	BillingCycles      *int32 `json:"billing_cycles,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	TrialEnd           *int64 `json:"trial_end,omitempty"`
}
type EditUpdateSubscriptionQuoteEventBasedAddon struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int64        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	QuantityInDecimal   string        `json:"quantity_in_decimal,omitempty"`
	UnitPriceInDecimal  string        `json:"unit_price_in_decimal,omitempty"`
}
type EditUpdateSubscriptionQuoteBillingAddress struct {
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
type EditUpdateSubscriptionQuoteShippingAddress struct {
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
type EditUpdateSubscriptionQuoteCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type EditUpdateSubscriptionQuoteContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CreateForOnetimeChargesRequest struct {
	Name               string                                      `json:"name,omitempty"`
	CustomerId         string                                      `json:"customer_id"`
	PoNumber           string                                      `json:"po_number,omitempty"`
	Notes              string                                      `json:"notes,omitempty"`
	ExpiresAt          *int64                                      `json:"expires_at,omitempty"`
	CurrencyCode       string                                      `json:"currency_code,omitempty"`
	Addons             []*CreateForOnetimeChargesAddon             `json:"addons,omitempty"`
	Charges            []*CreateForOnetimeChargesCharge            `json:"charges,omitempty"`
	Coupon             string                                      `json:"coupon,omitempty"`
	CouponIds          []string                                    `json:"coupon_ids,omitempty"`
	ShippingAddress    *CreateForOnetimeChargesShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields []*CreateForOnetimeChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type CreateForOnetimeChargesAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriod      *int32 `json:"service_period,omitempty"`
}
type CreateForOnetimeChargesCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
}
type CreateForOnetimeChargesShippingAddress struct {
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
type CreateForOnetimeChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type EditOneTimeQuoteRequest struct {
	PoNumber           string                               `json:"po_number,omitempty"`
	Notes              string                               `json:"notes,omitempty"`
	ExpiresAt          *int64                               `json:"expires_at,omitempty"`
	CurrencyCode       string                               `json:"currency_code,omitempty"`
	Addons             []*EditOneTimeQuoteAddon             `json:"addons,omitempty"`
	Charges            []*EditOneTimeQuoteCharge            `json:"charges,omitempty"`
	Coupon             string                               `json:"coupon,omitempty"`
	CouponIds          []string                             `json:"coupon_ids,omitempty"`
	ShippingAddress    *EditOneTimeQuoteShippingAddress     `json:"shipping_address,omitempty"`
	TaxProvidersFields []*EditOneTimeQuoteTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type EditOneTimeQuoteAddon struct {
	Id                 string `json:"id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriod      *int32 `json:"service_period,omitempty"`
}
type EditOneTimeQuoteCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
}
type EditOneTimeQuoteShippingAddress struct {
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
type EditOneTimeQuoteTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type CreateSubItemsForCustomerQuoteRequest struct {
	Name                   string                                            `json:"name,omitempty"`
	Notes                  string                                            `json:"notes,omitempty"`
	ExpiresAt              *int64                                            `json:"expires_at,omitempty"`
	Subscription           *CreateSubItemsForCustomerQuoteSubscription       `json:"subscription,omitempty"`
	BillingCycles          *int32                                            `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*CreateSubItemsForCustomerQuoteSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts              []*CreateSubItemsForCustomerQuoteDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                          `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*CreateSubItemsForCustomerQuoteItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                            `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                         `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *CreateSubItemsForCustomerQuoteShippingAddress    `json:"shipping_address,omitempty"`
	ContractTerm           *CreateSubItemsForCustomerQuoteContractTerm       `json:"contract_term,omitempty"`
	CouponIds              []string                                          `json:"coupon_ids,omitempty"`
	BillingStartOption     enum.BillingStartOption                           `json:"billing_start_option,omitempty"`
	BillingAddress         *CreateSubItemsForCustomerQuoteBillingAddress     `json:"billing_address,omitempty"`
	NetTermDays            *int32                                            `json:"net_term_days,omitempty"`
	Coupons                []*CreateSubItemsForCustomerQuoteCoupon           `json:"coupons,omitempty"`
}
type CreateSubItemsForCustomerQuoteSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type CreateSubItemsForCustomerQuoteSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	StartDate          *int64              `json:"start_date,omitempty"`
	EndDate            *int64              `json:"end_date,omitempty"`
	RampTierId         string              `json:"ramp_tier_id,omitempty"`
}
type CreateSubItemsForCustomerQuoteDiscount struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on,omitempty"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
	Quantity      *int32            `json:"quantity,omitempty"`
	StartDate     *int64            `json:"start_date,omitempty"`
	EndDate       *int64            `json:"end_date,omitempty"`
}
type CreateSubItemsForCustomerQuoteItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
	RampTierId            string           `json:"ramp_tier_id,omitempty"`
}
type CreateSubItemsForCustomerQuoteShippingAddress struct {
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
type CreateSubItemsForCustomerQuoteContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type CreateSubItemsForCustomerQuoteBillingAddress struct {
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
type CreateSubItemsForCustomerQuoteCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsRequest struct {
	Notes                  string                                                `json:"notes,omitempty"`
	ExpiresAt              *int64                                                `json:"expires_at,omitempty"`
	Subscription           *EditCreateSubCustomerQuoteForItemsSubscription       `json:"subscription,omitempty"`
	BillingCycles          *int32                                                `json:"billing_cycles,omitempty"`
	SubscriptionItems      []*EditCreateSubCustomerQuoteForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	Discounts              []*EditCreateSubCustomerQuoteForItemsDiscount         `json:"discounts,omitempty"`
	MandatoryItemsToRemove []string                                              `json:"mandatory_items_to_remove,omitempty"`
	ItemTiers              []*EditCreateSubCustomerQuoteForItemsItemTier         `json:"item_tiers,omitempty"`
	TermsToCharge          *int32                                                `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                             `json:"billing_alignment_mode,omitempty"`
	ShippingAddress        *EditCreateSubCustomerQuoteForItemsShippingAddress    `json:"shipping_address,omitempty"`
	ContractTerm           *EditCreateSubCustomerQuoteForItemsContractTerm       `json:"contract_term,omitempty"`
	CouponIds              []string                                              `json:"coupon_ids,omitempty"`
	BillingStartOption     enum.BillingStartOption                               `json:"billing_start_option,omitempty"`
	BillingAddress         *EditCreateSubCustomerQuoteForItemsBillingAddress     `json:"billing_address,omitempty"`
	NetTermDays            *int32                                                `json:"net_term_days,omitempty"`
	Coupons                []*EditCreateSubCustomerQuoteForItemsCoupon           `json:"coupons,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsSubscription struct {
	Id                                string                    `json:"id,omitempty"`
	PoNumber                          string                    `json:"po_number,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	StartDate          *int64              `json:"start_date,omitempty"`
	EndDate            *int64              `json:"end_date,omitempty"`
	RampTierId         string              `json:"ramp_tier_id,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsDiscount struct {
	ApplyOn       enum.ApplyOn      `json:"apply_on,omitempty"`
	DurationType  enum.DurationType `json:"duration_type"`
	Percentage    *float64          `json:"percentage,omitempty"`
	Amount        *int64            `json:"amount,omitempty"`
	Period        *int32            `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit   `json:"period_unit,omitempty"`
	IncludedInMrr *bool             `json:"included_in_mrr,omitempty"`
	ItemPriceId   string            `json:"item_price_id,omitempty"`
	Quantity      *int32            `json:"quantity,omitempty"`
	StartDate     *int64            `json:"start_date,omitempty"`
	EndDate       *int64            `json:"end_date,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
	RampTierId            string           `json:"ramp_tier_id,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsShippingAddress struct {
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
type EditCreateSubCustomerQuoteForItemsContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type EditCreateSubCustomerQuoteForItemsBillingAddress struct {
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
type EditCreateSubCustomerQuoteForItemsCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type UpdateSubscriptionQuoteForItemsRequest struct {
	Name                   string                                             `json:"name,omitempty"`
	Notes                  string                                             `json:"notes,omitempty"`
	ExpiresAt              *int64                                             `json:"expires_at,omitempty"`
	Subscription           *UpdateSubscriptionQuoteForItemsSubscription       `json:"subscription,omitempty"`
	SubscriptionItems      []*UpdateSubscriptionQuoteForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                           `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                              `json:"replace_items_list,omitempty"`
	Discounts              []*UpdateSubscriptionQuoteForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers              []*UpdateSubscriptionQuoteForItemsItemTier         `json:"item_tiers,omitempty"`
	BillingCycles          *int32                                             `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                             `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                             `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                          `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                           `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                              `json:"replace_coupon_list,omitempty"`
	ChangeOption           enum.ChangeOption                                  `json:"change_option,omitempty"`
	ChangesScheduledAt     *int64                                             `json:"changes_scheduled_at,omitempty"`
	ForceTermReset         *bool                                              `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                              `json:"reactivate,omitempty"`
	BillingAddress         *UpdateSubscriptionQuoteForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress        *UpdateSubscriptionQuoteForItemsShippingAddress    `json:"shipping_address,omitempty"`
	Customer               *UpdateSubscriptionQuoteForItemsCustomer           `json:"customer,omitempty"`
	ContractTerm           *UpdateSubscriptionQuoteForItemsContractTerm       `json:"contract_term,omitempty"`
	NetTermDays            *int32                                             `json:"net_term_days,omitempty"`
	Coupons                []*UpdateSubscriptionQuoteForItemsCoupon           `json:"coupons,omitempty"`
}
type UpdateSubscriptionQuoteForItemsSubscription struct {
	Id                                string                    `json:"id"`
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type UpdateSubscriptionQuoteForItemsSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	StartDate          *int64              `json:"start_date,omitempty"`
	EndDate            *int64              `json:"end_date,omitempty"`
	RampTierId         string              `json:"ramp_tier_id,omitempty"`
}
type UpdateSubscriptionQuoteForItemsDiscount struct {
	ApplyOn       enum.ApplyOn       `json:"apply_on,omitempty"`
	DurationType  enum.DurationType  `json:"duration_type"`
	Percentage    *float64           `json:"percentage,omitempty"`
	Amount        *int64             `json:"amount,omitempty"`
	Period        *int32             `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool              `json:"included_in_mrr,omitempty"`
	ItemPriceId   string             `json:"item_price_id,omitempty"`
	Quantity      *int32             `json:"quantity,omitempty"`
	OperationType enum.OperationType `json:"operation_type"`
	Id            string             `json:"id,omitempty"`
	StartDate     *int64             `json:"start_date,omitempty"`
	EndDate       *int64             `json:"end_date,omitempty"`
}
type UpdateSubscriptionQuoteForItemsItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
	RampTierId            string           `json:"ramp_tier_id,omitempty"`
}
type UpdateSubscriptionQuoteForItemsBillingAddress struct {
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
type UpdateSubscriptionQuoteForItemsShippingAddress struct {
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
type UpdateSubscriptionQuoteForItemsCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type UpdateSubscriptionQuoteForItemsContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type UpdateSubscriptionQuoteForItemsCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsRequest struct {
	Notes                  string                                                 `json:"notes,omitempty"`
	ExpiresAt              *int64                                                 `json:"expires_at,omitempty"`
	SubscriptionItems      []*EditUpdateSubscriptionQuoteForItemsSubscriptionItem `json:"subscription_items,omitempty"`
	MandatoryItemsToRemove []string                                               `json:"mandatory_items_to_remove,omitempty"`
	ReplaceItemsList       *bool                                                  `json:"replace_items_list,omitempty"`
	Subscription           *EditUpdateSubscriptionQuoteForItemsSubscription       `json:"subscription,omitempty"`
	Discounts              []*EditUpdateSubscriptionQuoteForItemsDiscount         `json:"discounts,omitempty"`
	ItemTiers              []*EditUpdateSubscriptionQuoteForItemsItemTier         `json:"item_tiers,omitempty"`
	BillingCycles          *int32                                                 `json:"billing_cycles,omitempty"`
	TermsToCharge          *int32                                                 `json:"terms_to_charge,omitempty"`
	ReactivateFrom         *int64                                                 `json:"reactivate_from,omitempty"`
	BillingAlignmentMode   enum.BillingAlignmentMode                              `json:"billing_alignment_mode,omitempty"`
	CouponIds              []string                                               `json:"coupon_ids,omitempty"`
	ReplaceCouponList      *bool                                                  `json:"replace_coupon_list,omitempty"`
	ChangeOption           enum.ChangeOption                                      `json:"change_option,omitempty"`
	ChangesScheduledAt     *int64                                                 `json:"changes_scheduled_at,omitempty"`
	ForceTermReset         *bool                                                  `json:"force_term_reset,omitempty"`
	Reactivate             *bool                                                  `json:"reactivate,omitempty"`
	BillingAddress         *EditUpdateSubscriptionQuoteForItemsBillingAddress     `json:"billing_address,omitempty"`
	ShippingAddress        *EditUpdateSubscriptionQuoteForItemsShippingAddress    `json:"shipping_address,omitempty"`
	Customer               *EditUpdateSubscriptionQuoteForItemsCustomer           `json:"customer,omitempty"`
	ContractTerm           *EditUpdateSubscriptionQuoteForItemsContractTerm       `json:"contract_term,omitempty"`
	NetTermDays            *int32                                                 `json:"net_term_days,omitempty"`
	Coupons                []*EditUpdateSubscriptionQuoteForItemsCoupon           `json:"coupons,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsSubscriptionItem struct {
	ItemPriceId        string              `json:"item_price_id"`
	Quantity           *int32              `json:"quantity,omitempty"`
	QuantityInDecimal  string              `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64              `json:"unit_price,omitempty"`
	UnitPriceInDecimal string              `json:"unit_price_in_decimal,omitempty"`
	BillingCycles      *int32              `json:"billing_cycles,omitempty"`
	TrialEnd           *int64              `json:"trial_end,omitempty"`
	ServicePeriodDays  *int32              `json:"service_period_days,omitempty"`
	ChargeOnEvent      enum.ChargeOnEvent  `json:"charge_on_event,omitempty"`
	ChargeOnce         *bool               `json:"charge_once,omitempty"`
	ChargeOnOption     enum.ChargeOnOption `json:"charge_on_option,omitempty"`
	ItemType           enum.ItemType       `json:"item_type,omitempty"`
	StartDate          *int64              `json:"start_date,omitempty"`
	EndDate            *int64              `json:"end_date,omitempty"`
	RampTierId         string              `json:"ramp_tier_id,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsSubscription struct {
	SetupFee                          *int64                    `json:"setup_fee,omitempty"`
	StartDate                         *int64                    `json:"start_date,omitempty"`
	TrialEnd                          *int64                    `json:"trial_end,omitempty"`
	Coupon                            string                    `json:"coupon,omitempty"`
	AutoCollection                    enum.AutoCollection       `json:"auto_collection,omitempty"`
	OfflinePaymentMethod              enum.OfflinePaymentMethod `json:"offline_payment_method,omitempty"`
	ContractTermBillingCycleOnRenewal *int32                    `json:"contract_term_billing_cycle_on_renewal,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsDiscount struct {
	ApplyOn       enum.ApplyOn       `json:"apply_on,omitempty"`
	DurationType  enum.DurationType  `json:"duration_type"`
	Percentage    *float64           `json:"percentage,omitempty"`
	Amount        *int64             `json:"amount,omitempty"`
	Period        *int32             `json:"period,omitempty"`
	PeriodUnit    enum.PeriodUnit    `json:"period_unit,omitempty"`
	IncludedInMrr *bool              `json:"included_in_mrr,omitempty"`
	ItemPriceId   string             `json:"item_price_id,omitempty"`
	Quantity      *int32             `json:"quantity,omitempty"`
	OperationType enum.OperationType `json:"operation_type"`
	Id            string             `json:"id,omitempty"`
	StartDate     *int64             `json:"start_date,omitempty"`
	EndDate       *int64             `json:"end_date,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsItemTier struct {
	ItemPriceId           string           `json:"item_price_id,omitempty"`
	StartingUnit          *int32           `json:"starting_unit,omitempty"`
	EndingUnit            *int32           `json:"ending_unit,omitempty"`
	Price                 *int64           `json:"price,omitempty"`
	StartingUnitInDecimal string           `json:"starting_unit_in_decimal,omitempty"`
	EndingUnitInDecimal   string           `json:"ending_unit_in_decimal,omitempty"`
	PriceInDecimal        string           `json:"price_in_decimal,omitempty"`
	PricingType           enum.PricingType `json:"pricing_type,omitempty"`
	PackageSize           *int32           `json:"package_size,omitempty"`
	RampTierId            string           `json:"ramp_tier_id,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsBillingAddress struct {
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
type EditUpdateSubscriptionQuoteForItemsShippingAddress struct {
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
type EditUpdateSubscriptionQuoteForItemsCustomer struct {
	VatNumber        string `json:"vat_number,omitempty"`
	VatNumberPrefix  string `json:"vat_number_prefix,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsContractTerm struct {
	ActionAtTermEnd          contractTerm.ActionAtTermEnd `json:"action_at_term_end,omitempty"`
	CancellationCutoffPeriod *int32                       `json:"cancellation_cutoff_period,omitempty"`
}
type EditUpdateSubscriptionQuoteForItemsCoupon struct {
	Id        string `json:"id,omitempty"`
	StartDate *int64 `json:"start_date,omitempty"`
	EndDate   *int64 `json:"end_date,omitempty"`
}
type CreateForChargeItemsAndChargesRequest struct {
	Name               string                                             `json:"name,omitempty"`
	CustomerId         string                                             `json:"customer_id"`
	PoNumber           string                                             `json:"po_number,omitempty"`
	Notes              string                                             `json:"notes,omitempty"`
	ExpiresAt          *int64                                             `json:"expires_at,omitempty"`
	CurrencyCode       string                                             `json:"currency_code,omitempty"`
	ItemPrices         []*CreateForChargeItemsAndChargesItemPrice         `json:"item_prices,omitempty"`
	ItemTiers          []*CreateForChargeItemsAndChargesItemTier          `json:"item_tiers,omitempty"`
	Charges            []*CreateForChargeItemsAndChargesCharge            `json:"charges,omitempty"`
	Coupon             string                                             `json:"coupon,omitempty"`
	CouponIds          []string                                           `json:"coupon_ids,omitempty"`
	BillingAddress     *CreateForChargeItemsAndChargesBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress    *CreateForChargeItemsAndChargesShippingAddress     `json:"shipping_address,omitempty"`
	Discounts          []*CreateForChargeItemsAndChargesDiscount          `json:"discounts,omitempty"`
	TaxProvidersFields []*CreateForChargeItemsAndChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type CreateForChargeItemsAndChargesItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
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
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
}
type CreateForChargeItemsAndChargesBillingAddress struct {
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
type CreateForChargeItemsAndChargesDiscount struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Quantity    *int32       `json:"quantity,omitempty"`
	Amount      *int64       `json:"amount,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
}
type CreateForChargeItemsAndChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type EditForChargeItemsAndChargesRequest struct {
	PoNumber           string                                           `json:"po_number,omitempty"`
	Notes              string                                           `json:"notes,omitempty"`
	ExpiresAt          *int64                                           `json:"expires_at,omitempty"`
	CurrencyCode       string                                           `json:"currency_code,omitempty"`
	ItemPrices         []*EditForChargeItemsAndChargesItemPrice         `json:"item_prices,omitempty"`
	ItemTiers          []*EditForChargeItemsAndChargesItemTier          `json:"item_tiers,omitempty"`
	Charges            []*EditForChargeItemsAndChargesCharge            `json:"charges,omitempty"`
	Coupon             string                                           `json:"coupon,omitempty"`
	CouponIds          []string                                         `json:"coupon_ids,omitempty"`
	BillingAddress     *EditForChargeItemsAndChargesBillingAddress      `json:"billing_address,omitempty"`
	ShippingAddress    *EditForChargeItemsAndChargesShippingAddress     `json:"shipping_address,omitempty"`
	Discounts          []*EditForChargeItemsAndChargesDiscount          `json:"discounts,omitempty"`
	TaxProvidersFields []*EditForChargeItemsAndChargesTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type EditForChargeItemsAndChargesItemPrice struct {
	ItemPriceId        string `json:"item_price_id,omitempty"`
	Quantity           *int32 `json:"quantity,omitempty"`
	QuantityInDecimal  string `json:"quantity_in_decimal,omitempty"`
	UnitPrice          *int64 `json:"unit_price,omitempty"`
	UnitPriceInDecimal string `json:"unit_price_in_decimal,omitempty"`
	ServicePeriodDays  *int32 `json:"service_period_days,omitempty"`
}
type EditForChargeItemsAndChargesItemTier struct {
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
type EditForChargeItemsAndChargesCharge struct {
	Amount                 *int64               `json:"amount,omitempty"`
	AmountInDecimal        string               `json:"amount_in_decimal,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	ServicePeriod          *int32               `json:"service_period,omitempty"`
}
type EditForChargeItemsAndChargesBillingAddress struct {
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
type EditForChargeItemsAndChargesShippingAddress struct {
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
type EditForChargeItemsAndChargesDiscount struct {
	Percentage  *float64     `json:"percentage,omitempty"`
	Quantity    *int32       `json:"quantity,omitempty"`
	Amount      *int64       `json:"amount,omitempty"`
	ApplyOn     enum.ApplyOn `json:"apply_on"`
	ItemPriceId string       `json:"item_price_id,omitempty"`
}
type EditForChargeItemsAndChargesTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type ListRequest struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	IncludeDeleted *bool                   `json:"include_deleted,omitempty"`
	Id             *filter.StringFilter    `json:"id,omitempty"`
	CustomerId     *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id,omitempty"`
	Status         *filter.EnumFilter      `json:"status,omitempty"`
	Date           *filter.TimestampFilter `json:"date,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type QuoteLineGroupsForQuoteRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type ConvertRequest struct {
	Subscription          *ConvertSubscription `json:"subscription,omitempty"`
	InvoiceDate           *int64               `json:"invoice_date,omitempty"`
	InvoiceImmediately    *bool                `json:"invoice_immediately,omitempty"`
	CreatePendingInvoices *bool                `json:"create_pending_invoices,omitempty"`
	FirstInvoicePending   *bool                `json:"first_invoice_pending,omitempty"`
}
type ConvertSubscription struct {
	Id                string              `json:"id,omitempty"`
	AutoCollection    enum.AutoCollection `json:"auto_collection,omitempty"`
	PoNumber          string              `json:"po_number,omitempty"`
	AutoCloseInvoices *bool               `json:"auto_close_invoices,omitempty"`
}
type UpdateStatusRequest struct {
	Status  Status `json:"status"`
	Comment string `json:"comment,omitempty"`
}
type ExtendExpiryDateRequest struct {
	ValidTill *int64 `json:"valid_till"`
}
type DeleteRequest struct {
	Comment string `json:"comment,omitempty"`
}
type PdfRequest struct {
	ConsolidatedView *bool                `json:"consolidated_view,omitempty"`
	DispositionType  enum.DispositionType `json:"disposition_type,omitempty"`
}

type RetrieveResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       *quotedcharge.QuotedCharge             `json:"quoted_charge,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type CreateSubForCustomerQuoteResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
}

type EditCreateSubForCustomerQuoteResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
}

type UpdateSubscriptionQuoteResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
}

type EditUpdateSubscriptionQuoteResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
}

type CreateForOnetimeChargesResponse struct {
	Quote        *Quote                     `json:"quote,omitempty"`
	QuotedCharge *quotedcharge.QuotedCharge `json:"quoted_charge,omitempty"`
}

type EditOneTimeQuoteResponse struct {
	Quote        *Quote                     `json:"quote,omitempty"`
	QuotedCharge *quotedcharge.QuotedCharge `json:"quoted_charge,omitempty"`
}

type CreateSubItemsForCustomerQuoteResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type EditCreateSubCustomerQuoteForItemsResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type UpdateSubscriptionQuoteForItemsResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type EditUpdateSubscriptionQuoteForItemsResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type CreateForChargeItemsAndChargesResponse struct {
	Quote        *Quote                     `json:"quote,omitempty"`
	QuotedCharge *quotedcharge.QuotedCharge `json:"quoted_charge,omitempty"`
}

type EditForChargeItemsAndChargesResponse struct {
	Quote        *Quote                     `json:"quote,omitempty"`
	QuotedCharge *quotedcharge.QuotedCharge `json:"quoted_charge,omitempty"`
}

type ListQuoteResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type ListResponse struct {
	List       []*ListQuoteResponse `json:"list,omitempty"`
	NextOffset string               `json:"next_offset,omitempty"`
}

type QuoteLineGroupsForQuoteQuoteResponse struct {
	QuoteLineGroup *quotelinegroup.QuoteLineGroup `json:"quote_line_group,omitempty"`
}

type QuoteLineGroupsForQuoteResponse struct {
	List       []*QuoteLineGroupsForQuoteQuoteResponse `json:"list,omitempty"`
	NextOffset string                                  `json:"next_offset,omitempty"`
}

type ConvertResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	Customer           *customer.Customer                     `json:"customer,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       *quotedcharge.QuotedCharge             `json:"quoted_charge,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
	Subscription       *subscription.Subscription             `json:"subscription,omitempty"`
	Invoice            *invoice.Invoice                       `json:"invoice,omitempty"`
	CreditNote         *creditnote.CreditNote                 `json:"credit_note,omitempty"`
	UnbilledCharges    []*unbilledcharge.UnbilledCharge       `json:"unbilled_charges,omitempty"`
}

type UpdateStatusResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       *quotedcharge.QuotedCharge             `json:"quoted_charge,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type ExtendExpiryDateResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       *quotedcharge.QuotedCharge             `json:"quoted_charge,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type DeleteResponse struct {
	Quote              *Quote                                 `json:"quote,omitempty"`
	QuotedSubscription *quotedsubscription.QuotedSubscription `json:"quoted_subscription,omitempty"`
	QuotedCharge       *quotedcharge.QuotedCharge             `json:"quoted_charge,omitempty"`
	QuotedRamp         *quotedramp.QuotedRamp                 `json:"quoted_ramp,omitempty"`
}

type PdfResponse struct {
	Download *download.Download `json:"download,omitempty"`
}
