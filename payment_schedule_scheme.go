package chargebee

type PaymentScheduleSchemePeriodUnit string

const (
	PaymentScheduleSchemePeriodUnitDay   PaymentScheduleSchemePeriodUnit = "day"
	PaymentScheduleSchemePeriodUnitWeek  PaymentScheduleSchemePeriodUnit = "week"
	PaymentScheduleSchemePeriodUnitMonth PaymentScheduleSchemePeriodUnit = "month"
)

// just struct
type PaymentScheduleScheme struct {
	Id                string                          `json:"id"`
	Name              string                          `json:"name"`
	Description       string                          `json:"description"`
	NumberOfSchedules int32                           `json:"number_of_schedules"`
	PeriodUnit        PaymentScheduleSchemePeriodUnit `json:"period_unit"`
	Period            int32                           `json:"period"`
	CreatedAt         int64                           `json:"created_at"`
	ResourceVersion   int64                           `json:"resource_version"`
	UpdatedAt         int64                           `json:"updated_at"`
	Object            string                          `json:"object"`
}

// sub resources
type PaymentScheduleSchemePreferredSchedule struct {
	Period           int32   `json:"period"`
	AmountPercentage float64 `json:"amount_percentage"`
	Object           string  `json:"object"`
}

// operations
// input params
type PaymentScheduleSchemeCreateRequest struct {
	NumberOfSchedules *int32                                         `json:"number_of_schedules"`
	PeriodUnit        PaymentScheduleSchemePeriodUnit                `json:"period_unit"`
	Period            *int32                                         `json:"period,omitempty"`
	Name              string                                         `json:"name"`
	FlexibleSchedules []*PaymentScheduleSchemeCreateFlexibleSchedule `json:"flexible_schedules,omitempty"`
	apiRequest        `json:"-" form:"-"`
}

func (r *PaymentScheduleSchemeCreateRequest) payload() any { return r }

// input sub resource params multi
type PaymentScheduleSchemeCreateFlexibleSchedule struct {
	Period           *int32   `json:"period,omitempty"`
	AmountPercentage *float64 `json:"amount_percentage,omitempty"`
}

// operation response
type PaymentScheduleSchemeCreateResponse struct {
	PaymentScheduleScheme *PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
	apiResponse
}

// operation response
type PaymentScheduleSchemeRetrieveResponse struct {
	PaymentScheduleScheme *PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
	apiResponse
}

// operation response
type PaymentScheduleSchemeDeleteResponse struct {
	PaymentScheduleScheme *PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
	apiResponse
}
