package paymentschedulescheme

import (
	paymentScheduleSchemeEnum "github.com/chargebee/chargebee-go/v3/models/paymentschedulescheme/enum"
)

type PaymentScheduleScheme struct {
	Id                 string                               `json:"id"`
	Name               string                               `json:"name"`
	Description        string                               `json:"description"`
	NumberOfSchedules  int32                                `json:"number_of_schedules"`
	PeriodUnit         paymentScheduleSchemeEnum.PeriodUnit `json:"period_unit"`
	Period             int32                                `json:"period"`
	CreatedAt          int64                                `json:"created_at"`
	ResourceVersion    int64                                `json:"resource_version"`
	UpdatedAt          int64                                `json:"updated_at"`
	PreferredSchedules []*PreferredSchedule                 `json:"preferred_schedules"`
	Object             string                               `json:"object"`
}
type PreferredSchedule struct {
	Period           int32   `json:"period"`
	AmountPercentage float64 `json:"amount_percentage"`
	Object           string  `json:"object"`
}
type CreateRequestParams struct {
	NumberOfSchedules *int32                               `json:"number_of_schedules"`
	PeriodUnit        paymentScheduleSchemeEnum.PeriodUnit `json:"period_unit"`
	Period            *int32                               `json:"period,omitempty"`
	Description       string                               `json:"description,omitempty"`
}
