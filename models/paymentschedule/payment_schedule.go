package paymentschedule

import (
	paymentScheduleEnum "github.com/chargebee/chargebee-go/v3/models/paymentschedule/enum"
)

type PaymentSchedule struct {
	Id              string                         `json:"id"`
	SchemeId        string                         `json:"scheme_id"`
	EntityType      paymentScheduleEnum.EntityType `json:"entity_type"`
	EntityId        string                         `json:"entity_id"`
	Amount          int64                          `json:"amount"`
	CreatedAt       int64                          `json:"created_at"`
	ResourceVersion int64                          `json:"resource_version"`
	UpdatedAt       int64                          `json:"updated_at"`
	CurrencyCode    string                         `json:"currency_code"`
	ScheduleEntries []*ScheduleEntry               `json:"schedule_entries"`
	Object          string                         `json:"object"`
}
type ScheduleEntry struct {
	Id     string                                  `json:"id"`
	Date   int64                                   `json:"date"`
	Amount int64                                   `json:"amount"`
	Status paymentScheduleEnum.ScheduleEntryStatus `json:"status"`
	Object string                                  `json:"object"`
}
