package paymentschedule

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	paymentScheduleEnum "github.com/chargebee/chargebee-go/v3/models/paymentschedule/enum"
	transactionEnum "github.com/chargebee/chargebee-go/v3/models/transaction/enum"
)

type PaymentSchedule struct {
	Id                    string                         `json:"id"`
	SchemeId              string                         `json:"scheme_id"`
	EntityType            paymentScheduleEnum.EntityType `json:"entity_type"`
	EntityId              string                         `json:"entity_id"`
	Amount                int64                          `json:"amount"`
	CreatedAt             int64                          `json:"created_at"`
	ResourceVersion       int64                          `json:"resource_version"`
	UpdatedAt             int64                          `json:"updated_at"`
	CurrencyCode          string                         `json:"currency_code"`
	ScheduleEntries       []*ScheduleEntry               `json:"schedule_entries"`
	ReferenceTransactions []*ReferenceTransaction        `json:"reference_transactions"`
	Object                string                         `json:"object"`
}
type ScheduleEntry struct {
	Id              string                                  `json:"id"`
	Date            int64                                   `json:"date"`
	Amount          int64                                   `json:"amount"`
	ScheduledAmount int64                                   `json:"scheduled_amount"`
	Status          paymentScheduleEnum.ScheduleEntryStatus `json:"status"`
	Object          string                                  `json:"object"`
}
type ReferenceTransaction struct {
	ScheduleEntryId string                 `json:"schedule_entry_id"`
	AppliedAmount   int64                  `json:"applied_amount"`
	TxnId           string                 `json:"txn_id"`
	TxnStatus       transactionEnum.Status `json:"txn_status"`
	TxnDate         int64                  `json:"txn_date"`
	TxnAmount       int64                  `json:"txn_amount"`
	Object          string                 `json:"object"`
}
type ListRequestParams struct {
	Limit     *int32                  `json:"limit,omitempty"`
	Offset    string                  `json:"offset,omitempty"`
	InvoiceId *filter.StringFilter    `json:"invoice_id,omitempty"`
	Id        *filter.StringFilter    `json:"id,omitempty"`
	UpdatedAt *filter.TimestampFilter `json:"updated_at,omitempty"`
}
