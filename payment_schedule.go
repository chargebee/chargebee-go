package chargebee

type PaymentScheduleEntityType string

const (
	PaymentScheduleEntityTypeInvoice PaymentScheduleEntityType = "invoice"
)

type PaymentScheduleScheduleEntryStatus string

const (
	PaymentScheduleScheduleEntryStatusPosted     PaymentScheduleScheduleEntryStatus = "posted"
	PaymentScheduleScheduleEntryStatusPaymentDue PaymentScheduleScheduleEntryStatus = "payment_due"
	PaymentScheduleScheduleEntryStatusPaid       PaymentScheduleScheduleEntryStatus = "paid"
)

// just struct
type PaymentSchedule struct {
	Id              string                          `json:"id"`
	SchemeId        string                          `json:"scheme_id"`
	EntityType      PaymentScheduleEntityType       `json:"entity_type"`
	EntityId        string                          `json:"entity_id"`
	Amount          int64                           `json:"amount"`
	CreatedAt       int64                           `json:"created_at"`
	ResourceVersion int64                           `json:"resource_version"`
	UpdatedAt       int64                           `json:"updated_at"`
	CurrencyCode    string                          `json:"currency_code"`
	ScheduleEntries []*PaymentScheduleScheduleEntry `json:"schedule_entries"`
	Object          string                          `json:"object"`
}

// sub resources
type PaymentScheduleScheduleEntry struct {
	Id     string              `json:"id"`
	Date   int64               `json:"date"`
	Amount int64               `json:"amount"`
	Status ScheduleEntryStatus `json:"status"`
	Object string              `json:"object"`
}

// operations
// input params
