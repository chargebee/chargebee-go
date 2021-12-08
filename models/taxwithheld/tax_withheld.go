package taxwithheld

type TaxWithheld struct {
	Id              string `json:"id"`
	ReferenceNumber string `json:"reference_number"`
	Description     string `json:"description"`
	Date            int64  `json:"date"`
	CurrencyCode    string `json:"currency_code"`
	Amount          int32  `json:"amount"`
	Object          string `json:"object"`
}
