package installmentconfig

import (
	installmentConfigEnum "github.com/chargebee/chargebee-go/v3/models/installmentconfig/enum"
)

type InstallmentConfig struct {
	Id                   string                           `json:"id"`
	Description          string                           `json:"description"`
	NumberOfInstallments int32                            `json:"number_of_installments"`
	PeriodUnit           installmentConfigEnum.PeriodUnit `json:"period_unit"`
	Period               int32                            `json:"period"`
	PreferredDay         int32                            `json:"preferred_day"`
	CreatedAt            int64                            `json:"created_at"`
	ResourceVersion      int64                            `json:"resource_version"`
	UpdatedAt            int64                            `json:"updated_at"`
	Installments         []*Installment                   `json:"installments"`
	Object               string                           `json:"object"`
}
type Installment struct {
	Period           int32   `json:"period"`
	AmountPercentage float64 `json:"amount_percentage"`
	Object           string  `json:"object"`
}
type CreateRequestParams struct {
	NumberOfInstallments *int32                           `json:"number_of_installments"`
	PeriodUnit           installmentConfigEnum.PeriodUnit `json:"period_unit"`
	Period               *int32                           `json:"period,omitempty"`
	PreferredDay         *int32                           `json:"preferred_day,omitempty"`
	Description          string                           `json:"description,omitempty"`
	Installments         []*CreateInstallmentParams       `json:"installments,omitempty"`
}
type CreateInstallmentParams struct {
	Period           *int32   `json:"period,omitempty"`
	AmountPercentage *float64 `json:"amount_percentage,omitempty"`
}
