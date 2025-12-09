package chargebee

type PeriodUnit string

const (
	PeriodUnitDay   PeriodUnit = "day"
	PeriodUnitWeek  PeriodUnit = "week"
	PeriodUnitMonth PeriodUnit = "month"
)

type PaymentScheduleScheme struct {
	Id                 string               `json:"id"`
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	NumberOfSchedules  int32                `json:"number_of_schedules"`
	PeriodUnit         PeriodUnit           `json:"period_unit"`
	Period             int32                `json:"period"`
	CreatedAt          int64                `json:"created_at"`
	ResourceVersion    int64                `json:"resource_version"`
	UpdatedAt          int64                `json:"updated_at"`
	PreferredSchedules []*PreferredSchedule `json:"preferred_schedules"`
	Object             string               `json:"object"`
}
type PreferredSchedule struct {
	Period           int32   `json:"period"`
	AmountPercentage float64 `json:"amount_percentage"`
	Object           string  `json:"object"`
}
type CreateRequest struct {
	NumberOfSchedules *int32                    `json:"number_of_schedules"`
	PeriodUnit        PeriodUnit                `json:"period_unit"`
	Period            *int32                    `json:"period,omitempty"`
	Name              string                    `json:"name"`
	FlexibleSchedules []*CreateFlexibleSchedule `json:"flexible_schedules,omitempty"`
}
type CreateFlexibleSchedule struct {
	Period           *int32   `json:"period,omitempty"`
	AmountPercentage *float64 `json:"amount_percentage,omitempty"`
}

type CreateResponse struct {
	PaymentScheduleScheme *PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type RetrieveResponse struct {
	PaymentScheduleScheme *PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type DeleteResponse struct {
	PaymentScheduleScheme *PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}
