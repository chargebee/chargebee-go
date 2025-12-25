package chargebee

type CurrencyForexType string

const (
	CurrencyForexTypeManual CurrencyForexType = "manual"
	CurrencyForexTypeAuto   CurrencyForexType = "auto"
)

// just struct
type Currency struct {
	Id                 string            `json:"id"`
	Enabled            bool              `json:"enabled"`
	ForexType          CurrencyForexType `json:"forex_type"`
	CurrencyCode       string            `json:"currency_code"`
	IsBaseCurrency     bool              `json:"is_base_currency"`
	ManualExchangeRate string            `json:"manual_exchange_rate"`
	Object             string            `json:"object"`
}

// sub resources
// operations
// input params
type CurrencyListRequest struct {
	Limit      *int32 `json:"limit,omitempty"`
	Offset     string `json:"offset,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *CurrencyListRequest) payload() any { return r }

type CurrencyCreateRequest struct {
	CurrencyCode       string            `json:"currency_code"`
	ForexType          CurrencyForexType `json:"forex_type"`
	ManualExchangeRate string            `json:"manual_exchange_rate,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *CurrencyCreateRequest) payload() any { return r }

type CurrencyUpdateRequest struct {
	ForexType          CurrencyForexType `json:"forex_type"`
	ManualExchangeRate string            `json:"manual_exchange_rate,omitempty"`
	apiRequest         `json:"-" form:"-"`
}

func (r *CurrencyUpdateRequest) payload() any { return r }

type CurrencyAddScheduleRequest struct {
	ManualExchangeRate string `json:"manual_exchange_rate"`
	ScheduleAt         *int64 `json:"schedule_at"`
	apiRequest         `json:"-" form:"-"`
}

func (r *CurrencyAddScheduleRequest) payload() any { return r }

// operation sub response
type CurrencyListCurrencyResponse struct {
	Currency *Currency `json:"currency,omitempty"`
}

// operation response
type CurrencyListResponse struct {
	List       []*CurrencyListCurrencyResponse `json:"list,omitempty"`
	NextOffset string                          `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type CurrencyRetrieveResponse struct {
	Currency *Currency `json:"currency,omitempty"`
	apiResponse
}

// operation response
type CurrencyCreateResponse struct {
	Currency *Currency `json:"currency,omitempty"`
	apiResponse
}

// operation response
type CurrencyUpdateResponse struct {
	Currency *Currency `json:"currency,omitempty"`
	apiResponse
}

// operation response
type CurrencyAddScheduleResponse struct {
	ScheduledAt int64     `json:"scheduled_at,omitempty"`
	Currency    *Currency `json:"currency,omitempty"`
	apiResponse
}

// operation response
type CurrencyRemoveScheduleResponse struct {
	ScheduledAt int64     `json:"scheduled_at,omitempty"`
	Currency    *Currency `json:"currency,omitempty"`
	apiResponse
}
