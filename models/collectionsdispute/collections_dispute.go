package collectionsdispute

import (
	collectionsDisputeEnum "github.com/chargebee/chargebee-go/v3/models/collectionsdispute/enum"
)

type CollectionsDispute struct {
	Id                 string                        `json:"id"`
	CustomerId         string                        `json:"customer_id"`
	InvoiceId          string                        `json:"invoice_id"`
	LineItemId         string                        `json:"line_item_id"`
	CurrencyCode       string                        `json:"currency_code"`
	BilledAmount       int64                         `json:"billed_amount"`
	ClaimedDeduction   int64                         `json:"claimed_deduction"`
	Status             collectionsDisputeEnum.Status `json:"status"`
	DisputeReasonCode  string                        `json:"dispute_reason_code"`
	ResolvedReasonCode string                        `json:"resolved_reason_code"`
	Reason             string                        `json:"reason"`
	OpenedAt           int64                         `json:"opened_at"`
	ResolvedAt         int64                         `json:"resolved_at"`
	CreatedAt          int64                         `json:"created_at"`
	UpdatedAt          int64                         `json:"updated_at"`
	ResourceVersion    int64                         `json:"resource_version"`
	UndisputedAmount   int64                         `json:"undisputed_amount"`
	Object             string                        `json:"object"`
}
