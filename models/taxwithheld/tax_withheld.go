package taxwithheld

import (
	taxWithheldEnum "github.com/chargebee/chargebee-go/models/taxwithheld/enum"
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
	Amount          int32                         `json:"amount"`
	ExchangeRate    float64                       `json:"exchange_rate"`
	Object          string                        `json:"object"`
}
