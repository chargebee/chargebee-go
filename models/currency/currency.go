package currency

import (
	currencyEnum "github.com/chargebee/chargebee-go/v3/models/currency/enum"
)

type Currency struct {
	Id                 string                 `json:"id"`
	Enabled            bool                   `json:"enabled"`
	ForexType          currencyEnum.ForexType `json:"forex_type"`
	CurrencyCode       string                 `json:"currency_code"`
	IsBaseCurrency     bool                   `json:"is_base_currency"`
	ManualExchangeRate string                 `json:"manual_exchange_rate"`
	Object             string                 `json:"object"`
}
type CreateRequestParams struct {
	CurrencyCode       string                 `json:"currency_code"`
	ForexType          currencyEnum.ForexType `json:"forex_type"`
	ManualExchangeRate string                 `json:"manual_exchange_rate,omitempty"`
}
type UpdateRequestParams struct {
	ForexType          currencyEnum.ForexType `json:"forex_type"`
	ManualExchangeRate string                 `json:"manual_exchange_rate,omitempty"`
}
type AddScheduleRequestParams struct {
	ManualExchangeRate string `json:"manual_exchange_rate"`
	ScheduleAt         *int64 `json:"schedule_at"`
}
