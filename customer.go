package chargebee

type VatNumberStatus string

const (
	VatNumberStatusValid        VatNumberStatus = "valid"
	VatNumberStatusInvalid      VatNumberStatus = "invalid"
	VatNumberStatusNotValidated VatNumberStatus = "not_validated"
	VatNumberStatusUndetermined VatNumberStatus = "undetermined"
)

type BillingDayOfWeek string

const (
	BillingDayOfWeekSunday    BillingDayOfWeek = "sunday"
	BillingDayOfWeekMonday    BillingDayOfWeek = "monday"
	BillingDayOfWeekTuesday   BillingDayOfWeek = "tuesday"
	BillingDayOfWeekWednesday BillingDayOfWeek = "wednesday"
	BillingDayOfWeekThursday  BillingDayOfWeek = "thursday"
	BillingDayOfWeekFriday    BillingDayOfWeek = "friday"
	BillingDayOfWeekSaturday  BillingDayOfWeek = "saturday"
)

type PiiCleared string

const (
	PiiClearedActive            PiiCleared = "active"
	PiiClearedScheduledForClear PiiCleared = "scheduled_for_clear"
	PiiClearedCleared           PiiCleared = "cleared"
)

type CardStatus string

const (
	CardStatusNoCard              CardStatus = "no_card"
	CardStatusValid               CardStatus = "valid"
	CardStatusExpiring            CardStatus = "expiring"
	CardStatusExpired             CardStatus = "expired"
	CardStatusPendingVerification CardStatus = "pending_verification"
	CardStatusInvalid             CardStatus = "invalid"
)

type FraudFlag string

const (
	FraudFlagSafe       FraudFlag = "safe"
	FraudFlagSuspicious FraudFlag = "suspicious"
	FraudFlagFraudulent FraudFlag = "fraudulent"
)

type PaymentMethodStatus string

const (
	PaymentMethodStatusValid               PaymentMethodStatus = "valid"
	PaymentMethodStatusExpiring            PaymentMethodStatus = "expiring"
	PaymentMethodStatusExpired             PaymentMethodStatus = "expired"
	PaymentMethodStatusInvalid             PaymentMethodStatus = "invalid"
	PaymentMethodStatusPendingVerification PaymentMethodStatus = "pending_verification"
)

type ParentAccountAccessPortalEditChildSubscriptions string

const (
	ParentAccountAccessPortalEditChildSubscriptionsYes      ParentAccountAccessPortalEditChildSubscriptions = "yes"
	ParentAccountAccessPortalEditChildSubscriptionsViewOnly ParentAccountAccessPortalEditChildSubscriptions = "view_only"
	ParentAccountAccessPortalEditChildSubscriptionsNo       ParentAccountAccessPortalEditChildSubscriptions = "no"
)

type ParentAccountAccessPortalDownloadChildInvoices string

const (
	ParentAccountAccessPortalDownloadChildInvoicesYes      ParentAccountAccessPortalDownloadChildInvoices = "yes"
	ParentAccountAccessPortalDownloadChildInvoicesViewOnly ParentAccountAccessPortalDownloadChildInvoices = "view_only"
	ParentAccountAccessPortalDownloadChildInvoicesNo       ParentAccountAccessPortalDownloadChildInvoices = "no"
)

type ChildAccountAccessPortalEditSubscriptions string

const (
	ChildAccountAccessPortalEditSubscriptionsYes      ChildAccountAccessPortalEditSubscriptions = "yes"
	ChildAccountAccessPortalEditSubscriptionsViewOnly ChildAccountAccessPortalEditSubscriptions = "view_only"
)

type ChildAccountAccessPortalDownloadInvoices string

const (
	ChildAccountAccessPortalDownloadInvoicesYes      ChildAccountAccessPortalDownloadInvoices = "yes"
	ChildAccountAccessPortalDownloadInvoicesViewOnly ChildAccountAccessPortalDownloadInvoices = "view_only"
	ChildAccountAccessPortalDownloadInvoicesNo       ChildAccountAccessPortalDownloadInvoices = "no"
)

type Customer struct {
	Id                               string                    `json:"id"`
	FirstName                        string                    `json:"first_name"`
	LastName                         string                    `json:"last_name"`
	Email                            string                    `json:"email"`
	Phone                            string                    `json:"phone"`
	Company                          string                    `json:"company"`
	VatNumber                        string                    `json:"vat_number"`
	AutoCollection                   enum.AutoCollection       `json:"auto_collection"`
	OfflinePaymentMethod             enum.OfflinePaymentMethod `json:"offline_payment_method"`
	NetTermDays                      int32                     `json:"net_term_days"`
	VatNumberValidatedTime           int64                     `json:"vat_number_validated_time"`
	VatNumberStatus                  VatNumberStatus           `json:"vat_number_status"`
	AllowDirectDebit                 bool                      `json:"allow_direct_debit"`
	IsLocationValid                  bool                      `json:"is_location_valid"`
	CreatedAt                        int64                     `json:"created_at"`
	CreatedFromIp                    string                    `json:"created_from_ip"`
	ExemptionDetails                 json.RawMessage           `json:"exemption_details"`
	Taxability                       enum.Taxability           `json:"taxability"`
	EntityCode                       enum.EntityCode           `json:"entity_code"`
	ExemptNumber                     string                    `json:"exempt_number"`
	ResourceVersion                  int64                     `json:"resource_version"`
	UpdatedAt                        int64                     `json:"updated_at"`
	Locale                           string                    `json:"locale"`
	BillingDate                      int32                     `json:"billing_date"`
	BillingMonth                     int32                     `json:"billing_month"`
	BillingDateMode                  enum.BillingDateMode      `json:"billing_date_mode"`
	BillingDayOfWeek                 BillingDayOfWeek          `json:"billing_day_of_week"`
	BillingDayOfWeekMode             enum.BillingDayOfWeekMode `json:"billing_day_of_week_mode"`
	PiiCleared                       PiiCleared                `json:"pii_cleared"`
	AutoCloseInvoices                bool                      `json:"auto_close_invoices"`
	Channel                          enum.Channel              `json:"channel"`
	ActiveId                         string                    `json:"active_id"`
	CardStatus                       CardStatus                `json:"card_status"`
	FraudFlag                        FraudFlag                 `json:"fraud_flag"`
	PrimaryPaymentSourceId           string                    `json:"primary_payment_source_id"`
	BackupPaymentSourceId            string                    `json:"backup_payment_source_id"`
	BillingAddress                   *BillingAddress           `json:"billing_address"`
	ReferralUrls                     []*ReferralUrl            `json:"referral_urls"`
	Contacts                         []*Contact                `json:"contacts"`
	PaymentMethod                    *PaymentMethod            `json:"payment_method"`
	InvoiceNotes                     string                    `json:"invoice_notes"`
	BusinessEntityId                 string                    `json:"business_entity_id"`
	PreferredCurrencyCode            string                    `json:"preferred_currency_code"`
	PromotionalCredits               int64                     `json:"promotional_credits"`
	UnbilledCharges                  int64                     `json:"unbilled_charges"`
	RefundableCredits                int64                     `json:"refundable_credits"`
	ExcessPayments                   int64                     `json:"excess_payments"`
	Balances                         []*Balance                `json:"balances"`
	EntityIdentifiers                []*EntityIdentifier       `json:"entity_identifiers"`
	TaxProvidersFields               []*TaxProvidersField      `json:"tax_providers_fields"`
	IsEinvoiceEnabled                bool                      `json:"is_einvoice_enabled"`
	EinvoicingMethod                 enum.EinvoicingMethod     `json:"einvoicing_method"`
	MetaData                         json.RawMessage           `json:"meta_data"`
	Deleted                          bool                      `json:"deleted"`
	RegisteredForGst                 bool                      `json:"registered_for_gst"`
	ConsolidatedInvoicing            bool                      `json:"consolidated_invoicing"`
	CustomerType                     enum.CustomerType         `json:"customer_type"`
	BusinessCustomerWithoutVatNumber bool                      `json:"business_customer_without_vat_number"`
	ClientProfileId                  string                    `json:"client_profile_id"`
	Relationship                     *Relationship             `json:"relationship"`
	UseDefaultHierarchySettings      bool                      `json:"use_default_hierarchy_settings"`
	ParentAccountAccess              *ParentAccountAccess      `json:"parent_account_access"`
	ChildAccountAccess               *ChildAccountAccess       `json:"child_account_access"`
	VatNumberPrefix                  string                    `json:"vat_number_prefix"`
	EntityIdentifierScheme           string                    `json:"entity_identifier_scheme"`
	EntityIdentifierStandard         string                    `json:"entity_identifier_standard"`
	CustomField                      map[string]interface{}    `json:"custom_field"`
	Consents                         map[string]interface{}    `json:"consents"`
	Object                           string                    `json:"object"`
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
type ReferralUrl struct {
	ExternalCustomerId         string              `json:"external_customer_id"`
	ReferralSharingUrl         string              `json:"referral_sharing_url"`
	CreatedAt                  int64               `json:"created_at"`
	UpdatedAt                  int64               `json:"updated_at"`
	ReferralCampaignId         string              `json:"referral_campaign_id"`
	ReferralAccountId          string              `json:"referral_account_id"`
	ReferralExternalCampaignId string              `json:"referral_external_campaign_id"`
	ReferralSystem             enum.ReferralSystem `json:"referral_system"`
	Object                     string              `json:"object"`
}
type Contact struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Label            string `json:"label"`
	Enabled          bool   `json:"enabled"`
	SendAccountEmail bool   `json:"send_account_email"`
	SendBillingEmail bool   `json:"send_billing_email"`
	Object           string `json:"object"`
}
type PaymentMethod struct {
	Type             enum.Type           `json:"type"`
	Gateway          enum.Gateway        `json:"gateway"`
	GatewayAccountId string              `json:"gateway_account_id"`
	Status           PaymentMethodStatus `json:"status"`
	ReferenceId      string              `json:"reference_id"`
	Object           string              `json:"object"`
}
type Balance struct {
	PromotionalCredits  int64  `json:"promotional_credits"`
	ExcessPayments      int64  `json:"excess_payments"`
	RefundableCredits   int64  `json:"refundable_credits"`
	UnbilledCharges     int64  `json:"unbilled_charges"`
	CurrencyCode        string `json:"currency_code"`
	BalanceCurrencyCode string `json:"balance_currency_code"`
	BusinessEntityId    string `json:"business_entity_id"`
	Object              string `json:"object"`
}
type EntityIdentifier struct {
	Id       string `json:"id"`
	Value    string `json:"value"`
	Scheme   string `json:"scheme"`
	Standard string `json:"standard"`
	Object   string `json:"object"`
}
type TaxProvidersField struct {
	ProviderName string `json:"provider_name"`
	FieldId      string `json:"field_id"`
	FieldValue   string `json:"field_value"`
	Object       string `json:"object"`
}
type Relationship struct {
	ParentId       string `json:"parent_id"`
	PaymentOwnerId string `json:"payment_owner_id"`
	InvoiceOwnerId string `json:"invoice_owner_id"`
	Object         string `json:"object"`
}
type ParentAccountAccess struct {
	PortalEditChildSubscriptions ParentAccountAccessPortalEditChildSubscriptions `json:"portal_edit_child_subscriptions"`
	PortalDownloadChildInvoices  ParentAccountAccessPortalDownloadChildInvoices  `json:"portal_download_child_invoices"`
	SendSubscriptionEmails       bool                                            `json:"send_subscription_emails"`
	SendInvoiceEmails            bool                                            `json:"send_invoice_emails"`
	SendPaymentEmails            bool                                            `json:"send_payment_emails"`
	Object                       string                                          `json:"object"`
}
type ChildAccountAccess struct {
	PortalEditSubscriptions ChildAccountAccessPortalEditSubscriptions `json:"portal_edit_subscriptions"`
	PortalDownloadInvoices  ChildAccountAccessPortalDownloadInvoices  `json:"portal_download_invoices"`
	SendSubscriptionEmails  bool                                      `json:"send_subscription_emails"`
	SendInvoiceEmails       bool                                      `json:"send_invoice_emails"`
	SendPaymentEmails       bool                                      `json:"send_payment_emails"`
	Object                  string                                    `json:"object"`
}
type CreateRequest struct {
	Id                               string                       `json:"id,omitempty"`
	FirstName                        string                       `json:"first_name,omitempty"`
	LastName                         string                       `json:"last_name,omitempty"`
	Email                            string                       `json:"email,omitempty"`
	PreferredCurrencyCode            string                       `json:"preferred_currency_code,omitempty"`
	Phone                            string                       `json:"phone,omitempty"`
	Company                          string                       `json:"company,omitempty"`
	AutoCollection                   enum.AutoCollection          `json:"auto_collection,omitempty"`
	NetTermDays                      *int32                       `json:"net_term_days,omitempty"`
	AllowDirectDebit                 *bool                        `json:"allow_direct_debit,omitempty"`
	VatNumber                        string                       `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                       `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                       `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard         string                       `json:"entity_identifier_standard,omitempty"`
	RegisteredForGst                 *bool                        `json:"registered_for_gst,omitempty"`
	IsEinvoiceEnabled                *bool                        `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 enum.EinvoicingMethod        `json:"einvoicing_method,omitempty"`
	Taxability                       enum.Taxability              `json:"taxability,omitempty"`
	ExemptionDetails                 []map[string]interface{}     `json:"exemption_details,omitempty"`
	CustomerType                     enum.CustomerType            `json:"customer_type,omitempty"`
	ClientProfileId                  string                       `json:"client_profile_id,omitempty"`
	TaxjarExemptionCategory          enum.TaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                        `json:"business_customer_without_vat_number,omitempty"`
	Locale                           string                       `json:"locale,omitempty"`
	EntityCode                       enum.EntityCode              `json:"entity_code,omitempty"`
	ExemptNumber                     string                       `json:"exempt_number,omitempty"`
	MetaData                         map[string]interface{}       `json:"meta_data,omitempty"`
	OfflinePaymentMethod             enum.OfflinePaymentMethod    `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices                *bool                        `json:"auto_close_invoices,omitempty"`
	ConsolidatedInvoicing            *bool                        `json:"consolidated_invoicing,omitempty"`
	Card                             *CreateCard                  `json:"card,omitempty"`
	BankAccount                      *CreateBankAccount           `json:"bank_account,omitempty"`
	TokenId                          string                       `json:"token_id,omitempty"`
	PaymentMethod                    *CreatePaymentMethod         `json:"payment_method,omitempty"`
	PaymentIntent                    *CreatePaymentIntent         `json:"payment_intent,omitempty"`
	BillingAddress                   *CreateBillingAddress        `json:"billing_address,omitempty"`
	EntityIdentifiers                []*CreateEntityIdentifier    `json:"entity_identifiers,omitempty"`
	BusinessEntityId                 string                       `json:"business_entity_id,omitempty"`
	TaxProvidersFields               []*CreateTaxProvidersField   `json:"tax_providers_fields,omitempty"`
	CreatedFromIp                    string                       `json:"created_from_ip,omitempty"`
	InvoiceNotes                     string                       `json:"invoice_notes,omitempty"`
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
type CreateBillingAddress struct {
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
type CreateEntityIdentifier struct {
	Id       string `json:"id,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Value    string `json:"value,omitempty"`
	Standard string `json:"standard,omitempty"`
}
type CreateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type ListRequest struct {
	Limit                *int32                  `json:"limit,omitempty"`
	Offset               string                  `json:"offset,omitempty"`
	Relationship         *ListRelationship       `json:"relationship,omitempty"`
	IncludeDeleted       *bool                   `json:"include_deleted,omitempty"`
	Id                   *filter.StringFilter    `json:"id,omitempty"`
	FirstName            *filter.StringFilter    `json:"first_name,omitempty"`
	LastName             *filter.StringFilter    `json:"last_name,omitempty"`
	Email                *filter.StringFilter    `json:"email,omitempty"`
	Company              *filter.StringFilter    `json:"company,omitempty"`
	Phone                *filter.StringFilter    `json:"phone,omitempty"`
	AutoCollection       *filter.EnumFilter      `json:"auto_collection,omitempty"`
	Taxability           *filter.EnumFilter      `json:"taxability,omitempty"`
	CreatedAt            *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt            *filter.TimestampFilter `json:"updated_at,omitempty"`
	OfflinePaymentMethod *filter.EnumFilter      `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices    *filter.BooleanFilter   `json:"auto_close_invoices,omitempty"`
	Channel              *filter.EnumFilter      `json:"channel,omitempty"`
	BusinessEntityId     *filter.StringFilter    `json:"business_entity_id,omitempty"`
	SortBy               *filter.SortFilter      `json:"sort_by,omitempty"`
}
type ListRelationship struct {
	ParentId       *filter.StringFilter `json:"parent_id,omitempty"`
	PaymentOwnerId *filter.StringFilter `json:"payment_owner_id,omitempty"`
	InvoiceOwnerId *filter.StringFilter `json:"invoice_owner_id,omitempty"`
}
type UpdateRequest struct {
	FirstName               string                       `json:"first_name,omitempty"`
	LastName                string                       `json:"last_name,omitempty"`
	Email                   string                       `json:"email,omitempty"`
	PreferredCurrencyCode   string                       `json:"preferred_currency_code,omitempty"`
	Phone                   string                       `json:"phone,omitempty"`
	Company                 string                       `json:"company,omitempty"`
	AutoCollection          enum.AutoCollection          `json:"auto_collection,omitempty"`
	AllowDirectDebit        *bool                        `json:"allow_direct_debit,omitempty"`
	NetTermDays             *int32                       `json:"net_term_days,omitempty"`
	Taxability              enum.Taxability              `json:"taxability,omitempty"`
	ExemptionDetails        []map[string]interface{}     `json:"exemption_details,omitempty"`
	CustomerType            enum.CustomerType            `json:"customer_type,omitempty"`
	ClientProfileId         string                       `json:"client_profile_id,omitempty"`
	TaxjarExemptionCategory enum.TaxjarExemptionCategory `json:"taxjar_exemption_category,omitempty"`
	Locale                  string                       `json:"locale,omitempty"`
	EntityCode              enum.EntityCode              `json:"entity_code,omitempty"`
	ExemptNumber            string                       `json:"exempt_number,omitempty"`
	OfflinePaymentMethod    enum.OfflinePaymentMethod    `json:"offline_payment_method,omitempty"`
	InvoiceNotes            string                       `json:"invoice_notes,omitempty"`
	AutoCloseInvoices       *bool                        `json:"auto_close_invoices,omitempty"`
	MetaData                map[string]interface{}       `json:"meta_data,omitempty"`
	FraudFlag               FraudFlag                    `json:"fraud_flag,omitempty"`
	ConsolidatedInvoicing   *bool                        `json:"consolidated_invoicing,omitempty"`
	TaxProvidersFields      []*UpdateTaxProvidersField   `json:"tax_providers_fields,omitempty"`
}
type UpdateTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type UpdatePaymentMethodRequest struct {
	PaymentMethod *UpdatePaymentMethodPaymentMethod `json:"payment_method,omitempty"`
}
type UpdatePaymentMethodPaymentMethod struct {
	Type                  enum.Type              `json:"type"`
	Gateway               enum.Gateway           `json:"gateway,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	IssuingCountry        string                 `json:"issuing_country,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type UpdateBillingInfoRequest struct {
	BillingAddress                   *UpdateBillingInfoBillingAddress      `json:"billing_address,omitempty"`
	EntityIdentifiers                []*UpdateBillingInfoEntityIdentifier  `json:"entity_identifiers,omitempty"`
	VatNumber                        string                                `json:"vat_number,omitempty"`
	VatNumberPrefix                  string                                `json:"vat_number_prefix,omitempty"`
	EntityIdentifierScheme           string                                `json:"entity_identifier_scheme,omitempty"`
	EntityIdentifierStandard         string                                `json:"entity_identifier_standard,omitempty"`
	RegisteredForGst                 *bool                                 `json:"registered_for_gst,omitempty"`
	BusinessCustomerWithoutVatNumber *bool                                 `json:"business_customer_without_vat_number,omitempty"`
	IsEinvoiceEnabled                *bool                                 `json:"is_einvoice_enabled,omitempty"`
	EinvoicingMethod                 enum.EinvoicingMethod                 `json:"einvoicing_method,omitempty"`
	TaxProvidersFields               []*UpdateBillingInfoTaxProvidersField `json:"tax_providers_fields,omitempty"`
}
type UpdateBillingInfoBillingAddress struct {
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
type UpdateBillingInfoEntityIdentifier struct {
	Id        string         `json:"id,omitempty"`
	Scheme    string         `json:"scheme,omitempty"`
	Value     string         `json:"value,omitempty"`
	Operation enum.Operation `json:"operation,omitempty"`
	Standard  string         `json:"standard,omitempty"`
}
type UpdateBillingInfoTaxProvidersField struct {
	ProviderName string `json:"provider_name,omitempty"`
	FieldId      string `json:"field_id,omitempty"`
	FieldValue   string `json:"field_value,omitempty"`
}
type ContactsForCustomerRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type AssignPaymentRoleRequest struct {
	PaymentSourceId string    `json:"payment_source_id"`
	Role            enum.Role `json:"role"`
}
type AddContactRequest struct {
	Contact *AddContactContact `json:"contact,omitempty"`
}
type AddContactContact struct {
	Id               string `json:"id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email"`
	Phone            string `json:"phone,omitempty"`
	Label            string `json:"label,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
	SendBillingEmail *bool  `json:"send_billing_email,omitempty"`
	SendAccountEmail *bool  `json:"send_account_email,omitempty"`
}
type UpdateContactRequest struct {
	Contact *UpdateContactContact `json:"contact,omitempty"`
}
type UpdateContactContact struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Label            string `json:"label,omitempty"`
	Enabled          *bool  `json:"enabled,omitempty"`
	SendBillingEmail *bool  `json:"send_billing_email,omitempty"`
	SendAccountEmail *bool  `json:"send_account_email,omitempty"`
}
type DeleteContactRequest struct {
	Contact *DeleteContactContact `json:"contact,omitempty"`
}
type DeleteContactContact struct {
	Id string `json:"id"`
}
type AddPromotionalCreditsRequest struct {
	Amount       *int64          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}
type DeductPromotionalCreditsRequest struct {
	Amount       *int64          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}
type SetPromotionalCreditsRequest struct {
	Amount       *int64          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}
type RecordExcessPaymentRequest struct {
	Transaction *RecordExcessPaymentTransaction `json:"transaction,omitempty"`
	Comment     string                          `json:"comment,omitempty"`
}
type RecordExcessPaymentTransaction struct {
	Id                    string             `json:"id,omitempty"`
	Amount                *int64             `json:"amount"`
	CurrencyCode          string             `json:"currency_code,omitempty"`
	Date                  *int64             `json:"date"`
	PaymentMethod         enum.PaymentMethod `json:"payment_method"`
	ReferenceNumber       string             `json:"reference_number,omitempty"`
	CustomPaymentMethodId string             `json:"custom_payment_method_id,omitempty"`
}
type CollectPaymentRequest struct {
	Amount                      *int64                             `json:"amount,omitempty"`
	InvoiceAllocations          []*CollectPaymentInvoiceAllocation `json:"invoice_allocations,omitempty"`
	PaymentSourceId             string                             `json:"payment_source_id,omitempty"`
	TokenId                     string                             `json:"token_id,omitempty"`
	PaymentMethod               *CollectPaymentPaymentMethod       `json:"payment_method,omitempty"`
	Card                        *CollectPaymentCard                `json:"card,omitempty"`
	PaymentIntent               *CollectPaymentPaymentIntent       `json:"payment_intent,omitempty"`
	ReplacePrimaryPaymentSource *bool                              `json:"replace_primary_payment_source,omitempty"`
	RetainPaymentSource         *bool                              `json:"retain_payment_source,omitempty"`
	PaymentInitiator            enum.PaymentInitiator              `json:"payment_initiator,omitempty"`
}
type CollectPaymentInvoiceAllocation struct {
	InvoiceId        string `json:"invoice_id"`
	AllocationAmount *int64 `json:"allocation_amount,omitempty"`
}
type CollectPaymentPaymentMethod struct {
	Type                  enum.Type              `json:"type,omitempty"`
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
	ReferenceId           string                 `json:"reference_id,omitempty"`
	TmpToken              string                 `json:"tmp_token,omitempty"`
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CollectPaymentCard struct {
	GatewayAccountId      string                 `json:"gateway_account_id,omitempty"`
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
	AdditionalInformation map[string]interface{} `json:"additional_information,omitempty"`
}
type CollectPaymentPaymentIntent struct {
	Id                    string                          `json:"id,omitempty"`
	GatewayAccountId      string                          `json:"gateway_account_id,omitempty"`
	GwToken               string                          `json:"gw_token,omitempty"`
	PaymentMethodType     paymentIntent.PaymentMethodType `json:"payment_method_type,omitempty"`
	GwPaymentMethodId     string                          `json:"gw_payment_method_id,omitempty"`
	ReferenceId           string                          `json:"reference_id,omitempty"`
	AdditionalInformation map[string]interface{}          `json:"additional_information,omitempty"`
}
type DeleteRequest struct {
	DeletePaymentMethod *bool `json:"delete_payment_method,omitempty"`
}
type MoveRequest struct {
	IdAtFromSite string `json:"id_at_from_site"`
	FromSite     string `json:"from_site"`
}
type ChangeBillingDateRequest struct {
	BillingDate          *int32                    `json:"billing_date,omitempty"`
	BillingMonth         *int32                    `json:"billing_month,omitempty"`
	BillingDateMode      enum.BillingDateMode      `json:"billing_date_mode,omitempty"`
	BillingDayOfWeek     BillingDayOfWeek          `json:"billing_day_of_week,omitempty"`
	BillingDayOfWeekMode enum.BillingDayOfWeekMode `json:"billing_day_of_week_mode,omitempty"`
}
type MergeRequest struct {
	FromCustomerId string `json:"from_customer_id"`
	ToCustomerId   string `json:"to_customer_id"`
}
type RelationshipsRequest struct {
	ParentId                    string                            `json:"parent_id,omitempty"`
	PaymentOwnerId              string                            `json:"payment_owner_id,omitempty"`
	InvoiceOwnerId              string                            `json:"invoice_owner_id,omitempty"`
	UseDefaultHierarchySettings *bool                             `json:"use_default_hierarchy_settings,omitempty"`
	ParentAccountAccess         *RelationshipsParentAccountAccess `json:"parent_account_access,omitempty"`
	ChildAccountAccess          *RelationshipsChildAccountAccess  `json:"child_account_access,omitempty"`
}
type RelationshipsParentAccountAccess struct {
	PortalEditChildSubscriptions customer.ParentAccountAccessPortalEditChildSubscriptions `json:"portal_edit_child_subscriptions,omitempty"`
	PortalDownloadChildInvoices  customer.ParentAccountAccessPortalDownloadChildInvoices  `json:"portal_download_child_invoices,omitempty"`
	SendSubscriptionEmails       *bool                                                    `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails            *bool                                                    `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails            *bool                                                    `json:"send_invoice_emails,omitempty"`
}
type RelationshipsChildAccountAccess struct {
	PortalEditSubscriptions customer.ChildAccountAccessPortalEditSubscriptions `json:"portal_edit_subscriptions,omitempty"`
	PortalDownloadInvoices  customer.ChildAccountAccessPortalDownloadInvoices  `json:"portal_download_invoices,omitempty"`
	SendSubscriptionEmails  *bool                                              `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails       *bool                                              `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails       *bool                                              `json:"send_invoice_emails,omitempty"`
}
type HierarchyRequest struct {
	HierarchyOperationType enum.HierarchyOperationType `json:"hierarchy_operation_type"`
}
type ListHierarchyDetailRequest struct {
	Limit                  *int32                      `json:"limit,omitempty"`
	Offset                 string                      `json:"offset,omitempty"`
	HierarchyOperationType enum.HierarchyOperationType `json:"hierarchy_operation_type"`
}
type UpdateHierarchySettingsRequest struct {
	UseDefaultHierarchySettings *bool                                       `json:"use_default_hierarchy_settings,omitempty"`
	ParentAccountAccess         *UpdateHierarchySettingsParentAccountAccess `json:"parent_account_access,omitempty"`
	ChildAccountAccess          *UpdateHierarchySettingsChildAccountAccess  `json:"child_account_access,omitempty"`
}
type UpdateHierarchySettingsParentAccountAccess struct {
	PortalEditChildSubscriptions customer.ParentAccountAccessPortalEditChildSubscriptions `json:"portal_edit_child_subscriptions,omitempty"`
	PortalDownloadChildInvoices  customer.ParentAccountAccessPortalDownloadChildInvoices  `json:"portal_download_child_invoices,omitempty"`
	SendSubscriptionEmails       *bool                                                    `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails            *bool                                                    `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails            *bool                                                    `json:"send_invoice_emails,omitempty"`
}
type UpdateHierarchySettingsChildAccountAccess struct {
	PortalEditSubscriptions customer.ChildAccountAccessPortalEditSubscriptions `json:"portal_edit_subscriptions,omitempty"`
	PortalDownloadInvoices  customer.ChildAccountAccessPortalDownloadInvoices  `json:"portal_download_invoices,omitempty"`
	SendSubscriptionEmails  *bool                                              `json:"send_subscription_emails,omitempty"`
	SendPaymentEmails       *bool                                              `json:"send_payment_emails,omitempty"`
	SendInvoiceEmails       *bool                                              `json:"send_invoice_emails,omitempty"`
}

type CreateResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type ListCustomerResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type ListResponse struct {
	List       []*ListCustomerResponse `json:"list,omitempty"`
	NextOffset string                  `json:"next_offset,omitempty"`
}

type RetrieveResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type UpdateResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type UpdatePaymentMethodResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type UpdateBillingInfoResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type ContactsForCustomerCustomerResponse struct {
	Contact *contact.Contact `json:"contact,omitempty"`
}

type ContactsForCustomerResponse struct {
	List       []*ContactsForCustomerCustomerResponse `json:"list,omitempty"`
	NextOffset string                                 `json:"next_offset,omitempty"`
}

type AssignPaymentRoleResponse struct {
	Customer      *Customer                    `json:"customer,omitempty"`
	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type AddContactResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type UpdateContactResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type DeleteContactResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type AddPromotionalCreditsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type DeductPromotionalCreditsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type SetPromotionalCreditsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type RecordExcessPaymentResponse struct {
	Customer    *Customer                `json:"customer,omitempty"`
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type CollectPaymentResponse struct {
	Customer    *Customer                `json:"customer,omitempty"`
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type DeleteResponse struct {
	Customer *Customer  `json:"customer,omitempty"`
	Card     *card.Card `json:"card,omitempty"`
}

type MoveResponse struct {
	ResourceMigration *resourcemigration.ResourceMigration `json:"resource_migration,omitempty"`
}

type ChangeBillingDateResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type MergeResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type ClearPersonalDataResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type RelationshipsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type DeleteRelationshipResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}

type HierarchyResponse struct {
	Hierarchies []*hierarchy.Hierarchy `json:"hierarchies,omitempty"`
}

type ListHierarchyDetailCustomerResponse struct {
	Hierarchy *hierarchy.Hierarchy `json:"hierarchy,omitempty"`
}

type ListHierarchyDetailResponse struct {
	List       []*ListHierarchyDetailCustomerResponse `json:"list,omitempty"`
	NextOffset string                                 `json:"next_offset,omitempty"`
}

type UpdateHierarchySettingsResponse struct {
	Customer *Customer `json:"customer,omitempty"`
}
