package chargebee

type MimeType string

const (
	MimeTypePdf MimeType = "pdf"
	MimeTypeZip MimeType = "zip"
)

type Status string

const (
	StatusInProcess Status = "in_process"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Export struct {
	Id            string    `json:"id"`
	OperationType string    `json:"operation_type"`
	MimeType      MimeType  `json:"mime_type"`
	Status        Status    `json:"status"`
	CreatedAt     int64     `json:"created_at"`
	Download      *Download `json:"download"`
	Object        string    `json:"object"`
}
type Download struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	MimeType    string `json:"mime_type"`
	Object      string `json:"object"`
}
type RevenueRecognitionRequest struct {
	Invoice          *RevenueRecognitionInvoice      `json:"invoice,omitempty"`
	Subscription     *RevenueRecognitionSubscription `json:"subscription,omitempty"`
	Customer         *RevenueRecognitionCustomer     `json:"customer,omitempty"`
	Relationship     *RevenueRecognitionRelationship `json:"relationship,omitempty"`
	ReportBy         enum.ReportBy                   `json:"report_by"`
	CurrencyCode     string                          `json:"currency_code,omitempty"`
	ReportFromMonth  *int32                          `json:"report_from_month"`
	ReportFromYear   *int32                          `json:"report_from_year"`
	ReportToMonth    *int32                          `json:"report_to_month"`
	ReportToYear     *int32                          `json:"report_to_year"`
	IncludeDiscounts *bool                           `json:"include_discounts,omitempty"`
	PaymentOwner     *filter.StringFilter            `json:"payment_owner,omitempty"`
	ItemId           *filter.StringFilter            `json:"item_id,omitempty"`
	ItemPriceId      *filter.StringFilter            `json:"item_price_id,omitempty"`
	CancelReasonCode *filter.StringFilter            `json:"cancel_reason_code,omitempty"`
	BusinessEntityId *filter.StringFilter            `json:"business_entity_id,omitempty"`
}
type RevenueRecognitionInvoice struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
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
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	Channel        *filter.EnumFilter      `json:"channel,omitempty"`
}
type RevenueRecognitionSubscription struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
	OfflinePaymentMethod   *filter.EnumFilter      `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices      *filter.BooleanFilter   `json:"auto_close_invoices,omitempty"`
	Channel                *filter.EnumFilter      `json:"channel,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
}
type RevenueRecognitionCustomer struct {
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
}
type RevenueRecognitionRelationship struct {
	ParentId       *filter.StringFilter `json:"parent_id,omitempty"`
	PaymentOwnerId *filter.StringFilter `json:"payment_owner_id,omitempty"`
	InvoiceOwnerId *filter.StringFilter `json:"invoice_owner_id,omitempty"`
}
type DeferredRevenueRequest struct {
	Invoice          *DeferredRevenueInvoice      `json:"invoice,omitempty"`
	Subscription     *DeferredRevenueSubscription `json:"subscription,omitempty"`
	Customer         *DeferredRevenueCustomer     `json:"customer,omitempty"`
	Relationship     *DeferredRevenueRelationship `json:"relationship,omitempty"`
	ReportBy         enum.ReportBy                `json:"report_by"`
	CurrencyCode     string                       `json:"currency_code,omitempty"`
	ReportFromMonth  *int32                       `json:"report_from_month"`
	ReportFromYear   *int32                       `json:"report_from_year"`
	ReportToMonth    *int32                       `json:"report_to_month"`
	ReportToYear     *int32                       `json:"report_to_year"`
	IncludeDiscounts *bool                        `json:"include_discounts,omitempty"`
	PaymentOwner     *filter.StringFilter         `json:"payment_owner,omitempty"`
	ItemId           *filter.StringFilter         `json:"item_id,omitempty"`
	ItemPriceId      *filter.StringFilter         `json:"item_price_id,omitempty"`
	CancelReasonCode *filter.StringFilter         `json:"cancel_reason_code,omitempty"`
	BusinessEntityId *filter.StringFilter         `json:"business_entity_id,omitempty"`
}
type DeferredRevenueInvoice struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
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
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	Channel        *filter.EnumFilter      `json:"channel,omitempty"`
}
type DeferredRevenueSubscription struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
	OfflinePaymentMethod   *filter.EnumFilter      `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices      *filter.BooleanFilter   `json:"auto_close_invoices,omitempty"`
	Channel                *filter.EnumFilter      `json:"channel,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
}
type DeferredRevenueCustomer struct {
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
}
type DeferredRevenueRelationship struct {
	ParentId       *filter.StringFilter `json:"parent_id,omitempty"`
	PaymentOwnerId *filter.StringFilter `json:"payment_owner_id,omitempty"`
	InvoiceOwnerId *filter.StringFilter `json:"invoice_owner_id,omitempty"`
}
type PlansRequest struct {
	Plan         *PlansPlan           `json:"plan,omitempty"`
	CurrencyCode *filter.StringFilter `json:"currency_code,omitempty"`
}
type PlansPlan struct {
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	Name               *filter.StringFilter    `json:"name,omitempty"`
	Price              *filter.NumberFilter    `json:"price,omitempty"`
	Period             *filter.NumberFilter    `json:"period,omitempty"`
	PeriodUnit         *filter.EnumFilter      `json:"period_unit,omitempty"`
	TrialPeriod        *filter.NumberFilter    `json:"trial_period,omitempty"`
	TrialPeriodUnit    *filter.EnumFilter      `json:"trial_period_unit,omitempty"`
	AddonApplicability *filter.EnumFilter      `json:"addon_applicability,omitempty"`
	Giftable           *filter.BooleanFilter   `json:"giftable,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
	Channel            *filter.EnumFilter      `json:"channel,omitempty"`
}
type AddonsRequest struct {
	Addon        *AddonsAddon         `json:"addon,omitempty"`
	CurrencyCode *filter.StringFilter `json:"currency_code,omitempty"`
}
type AddonsAddon struct {
	Id         *filter.StringFilter    `json:"id,omitempty"`
	Name       *filter.StringFilter    `json:"name,omitempty"`
	ChargeType *filter.EnumFilter      `json:"charge_type,omitempty"`
	Price      *filter.NumberFilter    `json:"price,omitempty"`
	Period     *filter.NumberFilter    `json:"period,omitempty"`
	PeriodUnit *filter.EnumFilter      `json:"period_unit,omitempty"`
	Status     *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt  *filter.TimestampFilter `json:"updated_at,omitempty"`
	Channel    *filter.EnumFilter      `json:"channel,omitempty"`
}
type CouponsRequest struct {
	Coupon       *CouponsCoupon       `json:"coupon,omitempty"`
	CurrencyCode *filter.StringFilter `json:"currency_code,omitempty"`
}
type CouponsCoupon struct {
	Id           *filter.StringFilter    `json:"id,omitempty"`
	Name         *filter.StringFilter    `json:"name,omitempty"`
	DiscountType *filter.EnumFilter      `json:"discount_type,omitempty"`
	DurationType *filter.EnumFilter      `json:"duration_type,omitempty"`
	Status       *filter.EnumFilter      `json:"status,omitempty"`
	ApplyOn      *filter.EnumFilter      `json:"apply_on,omitempty"`
	CreatedAt    *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt    *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type CustomersRequest struct {
	Customer         *CustomersCustomer     `json:"customer,omitempty"`
	Relationship     *CustomersRelationship `json:"relationship,omitempty"`
	ExportType       enum.ExportType        `json:"export_type,omitempty"`
	BusinessEntityId *filter.StringFilter   `json:"business_entity_id,omitempty"`
}
type CustomersCustomer struct {
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
}
type CustomersRelationship struct {
	ParentId       *filter.StringFilter `json:"parent_id,omitempty"`
	PaymentOwnerId *filter.StringFilter `json:"payment_owner_id,omitempty"`
	InvoiceOwnerId *filter.StringFilter `json:"invoice_owner_id,omitempty"`
}
type SubscriptionsRequest struct {
	Subscription     *SubscriptionsSubscription `json:"subscription,omitempty"`
	ExportType       enum.ExportType            `json:"export_type,omitempty"`
	ItemId           *filter.StringFilter       `json:"item_id,omitempty"`
	ItemPriceId      *filter.StringFilter       `json:"item_price_id,omitempty"`
	CancelReasonCode *filter.StringFilter       `json:"cancel_reason_code,omitempty"`
}
type SubscriptionsSubscription struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
	OfflinePaymentMethod   *filter.EnumFilter      `json:"offline_payment_method,omitempty"`
	AutoCloseInvoices      *filter.BooleanFilter   `json:"auto_close_invoices,omitempty"`
	Channel                *filter.EnumFilter      `json:"channel,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
}
type InvoicesRequest struct {
	Invoice      *InvoicesInvoice     `json:"invoice,omitempty"`
	PaymentOwner *filter.StringFilter `json:"payment_owner,omitempty"`
}
type InvoicesInvoice struct {
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
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
	Channel        *filter.EnumFilter      `json:"channel,omitempty"`
}
type CreditNotesRequest struct {
	CreditNote *CreditNotesCreditNote `json:"credit_note,omitempty"`
}
type CreditNotesCreditNote struct {
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	CustomerId         *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId     *filter.StringFilter    `json:"subscription_id,omitempty"`
	ReferenceInvoiceId *filter.StringFilter    `json:"reference_invoice_id,omitempty"`
	Type               *filter.EnumFilter      `json:"type,omitempty"`
	ReasonCode         *filter.EnumFilter      `json:"reason_code,omitempty"`
	CreateReasonCode   *filter.StringFilter    `json:"create_reason_code,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	Date               *filter.TimestampFilter `json:"date,omitempty"`
	Total              *filter.NumberFilter    `json:"total,omitempty"`
	PriceType          *filter.EnumFilter      `json:"price_type,omitempty"`
	AmountAllocated    *filter.NumberFilter    `json:"amount_allocated,omitempty"`
	AmountRefunded     *filter.NumberFilter    `json:"amount_refunded,omitempty"`
	AmountAvailable    *filter.NumberFilter    `json:"amount_available,omitempty"`
	VoidedAt           *filter.TimestampFilter `json:"voided_at,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
	Channel            *filter.EnumFilter      `json:"channel,omitempty"`
}
type TransactionsRequest struct {
	Transaction *TransactionsTransaction `json:"transaction,omitempty"`
}
type TransactionsTransaction struct {
	Id               *filter.StringFilter    `json:"id,omitempty"`
	CustomerId       *filter.StringFilter    `json:"customer_id,omitempty"`
	SubscriptionId   *filter.StringFilter    `json:"subscription_id,omitempty"`
	PaymentSourceId  *filter.StringFilter    `json:"payment_source_id,omitempty"`
	PaymentMethod    *filter.EnumFilter      `json:"payment_method,omitempty"`
	Gateway          *filter.EnumFilter      `json:"gateway,omitempty"`
	GatewayAccountId *filter.StringFilter    `json:"gateway_account_id,omitempty"`
	IdAtGateway      *filter.StringFilter    `json:"id_at_gateway,omitempty"`
	ReferenceNumber  *filter.StringFilter    `json:"reference_number,omitempty"`
	Type             *filter.EnumFilter      `json:"type,omitempty"`
	Date             *filter.TimestampFilter `json:"date,omitempty"`
	Amount           *filter.NumberFilter    `json:"amount,omitempty"`
	AmountCapturable *filter.NumberFilter    `json:"amount_capturable,omitempty"`
	Status           *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt        *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type OrdersRequest struct {
	Order *OrdersOrder         `json:"order,omitempty"`
	Total *filter.NumberFilter `json:"total,omitempty"`
}
type OrdersOrder struct {
	Id                      *filter.StringFilter    `json:"id,omitempty"`
	SubscriptionId          *filter.StringFilter    `json:"subscription_id,omitempty"`
	CustomerId              *filter.StringFilter    `json:"customer_id,omitempty"`
	Status                  *filter.EnumFilter      `json:"status,omitempty"`
	PriceType               *filter.EnumFilter      `json:"price_type,omitempty"`
	OrderDate               *filter.TimestampFilter `json:"order_date,omitempty"`
	ShippingDate            *filter.TimestampFilter `json:"shipping_date,omitempty"`
	ShippedAt               *filter.TimestampFilter `json:"shipped_at,omitempty"`
	DeliveredAt             *filter.TimestampFilter `json:"delivered_at,omitempty"`
	CancelledAt             *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	AmountPaid              *filter.NumberFilter    `json:"amount_paid,omitempty"`
	RefundableCredits       *filter.NumberFilter    `json:"refundable_credits,omitempty"`
	RefundableCreditsIssued *filter.NumberFilter    `json:"refundable_credits_issued,omitempty"`
	UpdatedAt               *filter.TimestampFilter `json:"updated_at,omitempty"`
	ResentStatus            *filter.EnumFilter      `json:"resent_status,omitempty"`
	IsResent                *filter.BooleanFilter   `json:"is_resent,omitempty"`
	OriginalOrderId         *filter.StringFilter    `json:"original_order_id,omitempty"`
}
type ItemFamiliesRequest struct {
	ItemFamily                *ItemFamiliesItemFamily `json:"item_family,omitempty"`
	BusinessEntityId          *filter.StringFilter    `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter   `json:"include_site_level_resources,omitempty"`
}
type ItemFamiliesItemFamily struct {
	Id        *filter.StringFilter    `json:"id,omitempty"`
	Name      *filter.StringFilter    `json:"name,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type ItemsRequest struct {
	Item                      *ItemsItem            `json:"item,omitempty"`
	BusinessEntityId          *filter.StringFilter  `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter `json:"include_site_level_resources,omitempty"`
}
type ItemsItem struct {
	Id                 *filter.StringFilter    `json:"id,omitempty"`
	ItemFamilyId       *filter.StringFilter    `json:"item_family_id,omitempty"`
	Type               *filter.EnumFilter      `json:"type,omitempty"`
	Name               *filter.StringFilter    `json:"name,omitempty"`
	ItemApplicability  *filter.EnumFilter      `json:"item_applicability,omitempty"`
	Status             *filter.EnumFilter      `json:"status,omitempty"`
	IsGiftable         *filter.BooleanFilter   `json:"is_giftable,omitempty"`
	UpdatedAt          *filter.TimestampFilter `json:"updated_at,omitempty"`
	EnabledForCheckout *filter.BooleanFilter   `json:"enabled_for_checkout,omitempty"`
	EnabledInPortal    *filter.BooleanFilter   `json:"enabled_in_portal,omitempty"`
	Metered            *filter.BooleanFilter   `json:"metered,omitempty"`
	UsageCalculation   *filter.EnumFilter      `json:"usage_calculation,omitempty"`
	Channel            *filter.EnumFilter      `json:"channel,omitempty"`
}
type ItemPricesRequest struct {
	ItemPrice                 *ItemPricesItemPrice  `json:"item_price,omitempty"`
	ItemFamilyId              *filter.StringFilter  `json:"item_family_id,omitempty"`
	ItemType                  *filter.EnumFilter    `json:"item_type,omitempty"`
	CurrencyCode              *filter.StringFilter  `json:"currency_code,omitempty"`
	BusinessEntityId          *filter.StringFilter  `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter `json:"include_site_level_resources,omitempty"`
}
type ItemPricesItemPrice struct {
	Id              *filter.StringFilter    `json:"id,omitempty"`
	Name            *filter.StringFilter    `json:"name,omitempty"`
	PricingModel    *filter.EnumFilter      `json:"pricing_model,omitempty"`
	ItemId          *filter.StringFilter    `json:"item_id,omitempty"`
	PriceVariantId  *filter.StringFilter    `json:"price_variant_id,omitempty"`
	TrialPeriod     *filter.NumberFilter    `json:"trial_period,omitempty"`
	TrialPeriodUnit *filter.EnumFilter      `json:"trial_period_unit,omitempty"`
	Status          *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt       *filter.TimestampFilter `json:"updated_at,omitempty"`
	PeriodUnit      *filter.EnumFilter      `json:"period_unit,omitempty"`
	Period          *filter.NumberFilter    `json:"period,omitempty"`
	Channel         *filter.EnumFilter      `json:"channel,omitempty"`
}
type AttachedItemsRequest struct {
	AttachedItem *AttachedItemsAttachedItem `json:"attached_item,omitempty"`
	ItemType     *filter.EnumFilter         `json:"item_type,omitempty"`
}
type AttachedItemsAttachedItem struct {
	Id            *filter.StringFilter    `json:"id,omitempty"`
	ItemId        *filter.StringFilter    `json:"item_id,omitempty"`
	Type          *filter.EnumFilter      `json:"type,omitempty"`
	ChargeOnEvent *filter.EnumFilter      `json:"charge_on_event,omitempty"`
	UpdatedAt     *filter.TimestampFilter `json:"updated_at,omitempty"`
	ParentItemId  *filter.StringFilter    `json:"parent_item_id,omitempty"`
}
type DifferentialPricesRequest struct {
	DifferentialPrice *DifferentialPricesDifferentialPrice `json:"differential_price,omitempty"`
	ItemId            *filter.StringFilter                 `json:"item_id,omitempty"`
}
type DifferentialPricesDifferentialPrice struct {
	ItemPriceId  *filter.StringFilter `json:"item_price_id,omitempty"`
	Id           *filter.StringFilter `json:"id,omitempty"`
	ParentItemId *filter.StringFilter `json:"parent_item_id,omitempty"`
}
type PriceVariantsRequest struct {
	PriceVariant              *PriceVariantsPriceVariant `json:"price_variant,omitempty"`
	BusinessEntityId          *filter.StringFilter       `json:"business_entity_id,omitempty"`
	IncludeSiteLevelResources *filter.BooleanFilter      `json:"include_site_level_resources,omitempty"`
}
type PriceVariantsPriceVariant struct {
	Id        *filter.StringFilter    `json:"id,omitempty"`
	Name      *filter.StringFilter    `json:"name,omitempty"`
	Status    *filter.EnumFilter      `json:"status,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt *filter.TimestampFilter `json:"created_at,omitempty"`
}

type RetrieveResponse struct {
	Export *Export `json:"export,omitempty"`
}

type RevenueRecognitionResponse struct {
	Export *Export `json:"export,omitempty"`
}

type DeferredRevenueResponse struct {
	Export *Export `json:"export,omitempty"`
}

type PlansResponse struct {
	Export *Export `json:"export,omitempty"`
}

type AddonsResponse struct {
	Export *Export `json:"export,omitempty"`
}

type CouponsResponse struct {
	Export *Export `json:"export,omitempty"`
}

type CustomersResponse struct {
	Export *Export `json:"export,omitempty"`
}

type SubscriptionsResponse struct {
	Export *Export `json:"export,omitempty"`
}

type InvoicesResponse struct {
	Export *Export `json:"export,omitempty"`
}

type CreditNotesResponse struct {
	Export *Export `json:"export,omitempty"`
}

type TransactionsResponse struct {
	Export *Export `json:"export,omitempty"`
}

type OrdersResponse struct {
	Export *Export `json:"export,omitempty"`
}

type ItemFamiliesResponse struct {
	Export *Export `json:"export,omitempty"`
}

type ItemsResponse struct {
	Export *Export `json:"export,omitempty"`
}

type ItemPricesResponse struct {
	Export *Export `json:"export,omitempty"`
}

type AttachedItemsResponse struct {
	Export *Export `json:"export,omitempty"`
}

type DifferentialPricesResponse struct {
	Export *Export `json:"export,omitempty"`
}

type PriceVariantsResponse struct {
	Export *Export `json:"export,omitempty"`
}
