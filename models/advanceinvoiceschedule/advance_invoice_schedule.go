package advanceinvoiceschedule

import (
	"github.com/chargebee/chargebee-go/enum"
	advanceInvoiceScheduleEnum "github.com/chargebee/chargebee-go/models/advanceinvoiceschedule/enum"
)

type AdvanceInvoiceSchedule struct {
	Id                    string                                  `json:"id"`
	ScheduleType          advanceInvoiceScheduleEnum.ScheduleType `json:"schedule_type"`
	FixedIntervalSchedule *FixedIntervalSchedule                  `json:"fixed_interval_schedule"`
	SpecificDatesSchedule *SpecificDatesSchedule                  `json:"specific_dates_schedule"`
	Object                string                                  `json:"object"`
}
type FixedIntervalSchedule struct {
	EndScheduleOn       enum.EndScheduleOn `json:"end_schedule_on"`
	NumberOfOccurrences int32              `json:"number_of_occurrences"`
	DaysBeforeRenewal   int32              `json:"days_before_renewal"`
	EndDate             int64              `json:"end_date"`
	CreatedAt           int64              `json:"created_at"`
	TermsToCharge       int32              `json:"terms_to_charge"`
	Object              string             `json:"object"`
}
type SpecificDatesSchedule struct {
	TermsToCharge int32  `json:"terms_to_charge"`
	Date          int64  `json:"date"`
	CreatedAt     int64  `json:"created_at"`
	Object        string `json:"object"`
}
