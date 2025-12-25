package chargebee

type AdvanceInvoiceScheduleScheduleType string

const (
	AdvanceInvoiceScheduleScheduleTypeFixedIntervals AdvanceInvoiceScheduleScheduleType = "fixed_intervals"
	AdvanceInvoiceScheduleScheduleTypeSpecificDates  AdvanceInvoiceScheduleScheduleType = "specific_dates"
)

type AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOn string

const (
	AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOnAfterNumberOfIntervals AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOn = "after_number_of_intervals"
	AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOnSpecificDate           AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOn = "specific_date"
	AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOnSubscriptionEnd        AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOn = "subscription_end"
)

// just struct
type AdvanceInvoiceSchedule struct {
	Id                    string                             `json:"id"`
	ScheduleType          AdvanceInvoiceScheduleScheduleType `json:"schedule_type"`
	FixedIntervalSchedule *FixedIntervalSchedule             `json:"fixed_interval_schedule"`
	SpecificDatesSchedule *SpecificDatesSchedule             `json:"specific_dates_schedule"`
	Object                string                             `json:"object"`
}

// sub resources
type AdvanceInvoiceScheduleFixedIntervalSchedule struct {
	EndScheduleOn       AdvanceInvoiceScheduleFixedIntervalScheduleEndScheduleOn `json:"end_schedule_on"`
	NumberOfOccurrences int32                                                    `json:"number_of_occurrences"`
	DaysBeforeRenewal   int32                                                    `json:"days_before_renewal"`
	EndDate             int64                                                    `json:"end_date"`
	CreatedAt           int64                                                    `json:"created_at"`
	TermsToCharge       int32                                                    `json:"terms_to_charge"`
	Object              string                                                   `json:"object"`
}
type AdvanceInvoiceScheduleSpecificDatesSchedule struct {
	TermsToCharge int32  `json:"terms_to_charge"`
	Date          int64  `json:"date"`
	CreatedAt     int64  `json:"created_at"`
	Object        string `json:"object"`
}

// operations
// input params
