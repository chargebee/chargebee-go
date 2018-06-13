package export

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	exportEnum "github.com/chargebee/chargebee-go/models/export/enum"
)

type Export struct {
	OperationType string              `json:"operation_type"`
	MimeType      exportEnum.MimeType `json:"mime_type"`
	Status        exportEnum.Status   `json:"status"`
	CreatedAt     int64               `json:"created_at"`
	Id            string              `json:"id"`
	Download      *Download           `json:"download"`
	Object        string              `json:"object"`
}
type Download struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	Object      string `json:"object"`
}

type RevenueRecognitionRequestParams struct {
	ReportBy         enum.ReportBy                         `json:"report_by"`
	CurrencyCode     string                                `json:"currency_code,omitempty"`
	ReportFromYear   *int32                                `json:"report_from_year"`
	ReportToYear     *int32                                `json:"report_to_year"`
	ReportToMonth    *int32                                `json:"report_to_month"`
	ReportFromMonth  *int32                                `json:"report_from_month"`
	IncludeDiscounts *bool                                 `json:"include_discounts,omitempty"`
	Invoice          *RevenueRecognitionInvoiceParams      `json:"invoice,omitempty"`
	Subscription     *RevenueRecognitionSubscriptionParams `json:"subscription,omitempty"`
	Customer         *RevenueRecognitionCustomerParams     `json:"customer,omitempty"`
}
type RevenueRecognitionInvoiceParams struct {
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
}
type RevenueRecognitionSubscriptionParams struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type RevenueRecognitionCustomerParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	FirstName      *filter.StringFilter    `json:"first_name,omitempty"`
	LastName       *filter.StringFilter    `json:"last_name,omitempty"`
	Email          *filter.StringFilter    `json:"email,omitempty"`
	Company        *filter.StringFilter    `json:"company,omitempty"`
	Phone          *filter.StringFilter    `json:"phone,omitempty"`
	AutoCollection *filter.EnumFilter      `json:"auto_collection,omitempty"`
	Taxability     *filter.EnumFilter      `json:"taxability,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}

type DeferredRevenueRequestParams struct {
	ReportBy         enum.ReportBy                      `json:"report_by"`
	CurrencyCode     string                             `json:"currency_code,omitempty"`
	ReportFromYear   *int32                             `json:"report_from_year"`
	ReportToYear     *int32                             `json:"report_to_year"`
	ReportToMonth    *int32                             `json:"report_to_month"`
	ReportFromMonth  *int32                             `json:"report_from_month"`
	IncludeDiscounts *bool                              `json:"include_discounts,omitempty"`
	Invoice          *DeferredRevenueInvoiceParams      `json:"invoice,omitempty"`
	Subscription     *DeferredRevenueSubscriptionParams `json:"subscription,omitempty"`
	Customer         *DeferredRevenueCustomerParams     `json:"customer,omitempty"`
}
type DeferredRevenueInvoiceParams struct {
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
}
type DeferredRevenueSubscriptionParams struct {
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
}
type DeferredRevenueCustomerParams struct {
	Id             *filter.StringFilter    `json:"id,omitempty"`
	FirstName      *filter.StringFilter    `json:"first_name,omitempty"`
	LastName       *filter.StringFilter    `json:"last_name,omitempty"`
	Email          *filter.StringFilter    `json:"email,omitempty"`
	Company        *filter.StringFilter    `json:"company,omitempty"`
	Phone          *filter.StringFilter    `json:"phone,omitempty"`
	AutoCollection *filter.EnumFilter      `json:"auto_collection,omitempty"`
	Taxability     *filter.EnumFilter      `json:"taxability,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	UpdatedAt      *filter.TimestampFilter `json:"updated_at,omitempty"`
}
