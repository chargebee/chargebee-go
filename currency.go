package chargebee

type ForexType string

const (
	ForexTypeManual ForexType = "manual"
	ForexTypeAuto   ForexType = "auto"
)

type Currency struct {
	Id                 string    `json:"id"`
	Enabled            bool      `json:"enabled"`
	ForexType          ForexType `json:"forex_type"`
	CurrencyCode       string    `json:"currency_code"`
	IsBaseCurrency     bool      `json:"is_base_currency"`
	ManualExchangeRate string    `json:"manual_exchange_rate"`
	Object             string    `json:"object"`
}
type ListRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}
type CreateRequest struct {
	CurrencyCode       string    `json:"currency_code"`
	ForexType          ForexType `json:"forex_type"`
	ManualExchangeRate string    `json:"manual_exchange_rate,omitempty"`
}
type UpdateRequest struct {
	ForexType          ForexType `json:"forex_type"`
	ManualExchangeRate string    `json:"manual_exchange_rate,omitempty"`
}
type AddScheduleRequest struct {
	ManualExchangeRate string `json:"manual_exchange_rate"`
	ScheduleAt         *int64 `json:"schedule_at"`
}

type ListCurrencyResponse struct {
	Currency *Currency `json:"currency,omitempty"`
}

type ListResponse struct {
	List       []*ListCurrencyResponse `json:"list,omitempty"`
	NextOffset string                  `json:"next_offset,omitempty"`
}

type RetrieveResponse struct {
	Currency *Currency `json:"currency,omitempty"`
}

type CreateResponse struct {
	Currency *Currency `json:"currency,omitempty"`
}

type UpdateResponse struct {
	Currency *Currency `json:"currency,omitempty"`
}

type AddScheduleResponse struct {
	ScheduledAt int64     `json:"scheduled_at,omitempty"`
	Currency    *Currency `json:"currency,omitempty"`
}

type RemoveScheduleResponse struct {
	ScheduledAt int64     `json:"scheduled_at,omitempty"`
	Currency    *Currency `json:"currency,omitempty"`
}
