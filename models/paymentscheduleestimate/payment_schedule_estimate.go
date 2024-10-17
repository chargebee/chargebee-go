package paymentscheduleestimate

import (
	paymentScheduleEstimateEnum "github.com/chargebee/chargebee-go/v3/models/paymentscheduleestimate/enum"
)

type PaymentScheduleEstimate struct {
	Id              string                                 `json:"id"`
	SchemeId        string                                 `json:"scheme_id"`
	EntityType      paymentScheduleEstimateEnum.EntityType `json:"entity_type"`
	EntityId        string                                 `json:"entity_id"`
	Amount          int64                                  `json:"amount"`
	CurrencyCode    string                                 `json:"currency_code"`
	ScheduleEntries []*ScheduleEntry                       `json:"schedule_entries"`
	Object          string                                 `json:"object"`
}
type ScheduleEntry struct {
	Id     string                                          `json:"id"`
	Date   int64                                           `json:"date"`
	Amount int64                                           `json:"amount"`
	Status paymentScheduleEstimateEnum.ScheduleEntryStatus `json:"status"`
	Object string                                          `json:"object"`
}
