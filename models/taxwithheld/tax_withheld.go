package taxwithheld

import (
	taxWithheldEnum "github.com/chargebee/chargebee-go/v3/models/taxwithheld/enum"
)

type TaxWithheld struct {
	Id string `json:"id"`
	//Deprecated: this field is deprecated
	User            string `json:"user"`
	ReferenceNumber string `json:"reference_number"`
	Description     string `json:"description"`
	//Deprecated: this field is deprecated
	Type taxWithheldEnum.Type `json:"type"`
	//Deprecated: this field is deprecated
	PaymentMethod taxWithheldEnum.PaymentMethod `json:"payment_method"`
	Date          int64                         `json:"date"`
	//Deprecated: this field is deprecated
	CurrencyCode    string `json:"currency_code"`
	Amount          int64  `json:"amount"`
	ResourceVersion int64  `json:"resource_version"`
	UpdatedAt       int64  `json:"updated_at"`
	//Deprecated: this field is deprecated
	ExchangeRate float64 `json:"exchange_rate"`
	Object       string  `json:"object"`
}
