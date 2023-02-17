package taxwithheld

import (
	taxWithheldEnum "github.com/chargebee/chargebee-go/v3/models/taxwithheld/enum"
)

type TaxWithheld struct {
	Id              string                        `json:"id"`
	User            string                        `json:"user"`
	ReferenceNumber string                        `json:"reference_number"`
	Description     string                        `json:"description"`
	Type            taxWithheldEnum.Type          `json:"type"`
	PaymentMethod   taxWithheldEnum.PaymentMethod `json:"payment_method"`
	Date            int64                         `json:"date"`
	CurrencyCode    string                        `json:"currency_code"`
	Amount          int64                         `json:"amount"`
	ResourceVersion int64                         `json:"resource_version"`
	UpdatedAt       int64                         `json:"updated_at"`
	ExchangeRate    float64                       `json:"exchange_rate"`
	Object          string                        `json:"object"`
}
