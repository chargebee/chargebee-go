package csvtaxrule

import (
	"github.com/chargebee/chargebee-go/enum"
	csvTaxRuleEnum "github.com/chargebee/chargebee-go/models/csvtaxrule/enum"
)

type CsvTaxRule struct {
	TaxProfileName string                     `json:"tax_profile_name"`
	Country        string                     `json:"country"`
	State          string                     `json:"state"`
	ZipCode        string                     `json:"zip_code"`
	ZipCodeStart   int32                      `json:"zip_code_start"`
	ZipCodeEnd     int32                      `json:"zip_code_end"`
	Tax1Name       string                     `json:"tax1_name"`
	Tax1Rate       float64                    `json:"tax1_rate"`
	Tax1JurisType  enum.Tax1JurisType         `json:"tax1_juris_type"`
	Tax1JurisName  string                     `json:"tax1_juris_name"`
	Tax1JurisCode  string                     `json:"tax1_juris_code"`
	Tax2Name       string                     `json:"tax2_name"`
	Tax2Rate       float64                    `json:"tax2_rate"`
	Tax2JurisType  enum.Tax2JurisType         `json:"tax2_juris_type"`
	Tax2JurisName  string                     `json:"tax2_juris_name"`
	Tax2JurisCode  string                     `json:"tax2_juris_code"`
	Tax3Name       string                     `json:"tax3_name"`
	Tax3Rate       float64                    `json:"tax3_rate"`
	Tax3JurisType  enum.Tax3JurisType         `json:"tax3_juris_type"`
	Tax3JurisName  string                     `json:"tax3_juris_name"`
	Tax3JurisCode  string                     `json:"tax3_juris_code"`
	Tax4Name       string                     `json:"tax4_name"`
	Tax4Rate       float64                    `json:"tax4_rate"`
	Tax4JurisType  enum.Tax4JurisType         `json:"tax4_juris_type"`
	Tax4JurisName  string                     `json:"tax4_juris_name"`
	Tax4JurisCode  string                     `json:"tax4_juris_code"`
	Status         csvTaxRuleEnum.Status      `json:"status"`
	TimeZone       string                     `json:"time_zone"`
	ValidFrom      int64                      `json:"valid_from"`
	ValidTill      int64                      `json:"valid_till"`
	ServiceType    csvTaxRuleEnum.ServiceType `json:"service_type"`
	RuleWeight     int32                      `json:"rule_weight"`
	Overwrite      bool                       `json:"overwrite"`
	Object         string                     `json:"object"`
}
type CreateRequestParams struct {
	TaxProfileName string                     `json:"tax_profile_name,omitempty"`
	Country        string                     `json:"country,omitempty"`
	State          string                     `json:"state,omitempty"`
	ZipCode        string                     `json:"zip_code,omitempty"`
	ZipCodeStart   *int32                     `json:"zip_code_start,omitempty"`
	ZipCodeEnd     *int32                     `json:"zip_code_end,omitempty"`
	Tax1Name       string                     `json:"tax1_name,omitempty"`
	Tax1Rate       *float64                   `json:"tax1_rate,omitempty"`
	Tax1JurisType  enum.Tax1JurisType         `json:"tax1_juris_type,omitempty"`
	Tax1JurisName  string                     `json:"tax1_juris_name,omitempty"`
	Tax1JurisCode  string                     `json:"tax1_juris_code,omitempty"`
	Tax2Name       string                     `json:"tax2_name,omitempty"`
	Tax2Rate       *float64                   `json:"tax2_rate,omitempty"`
	Tax2JurisType  enum.Tax2JurisType         `json:"tax2_juris_type,omitempty"`
	Tax2JurisName  string                     `json:"tax2_juris_name,omitempty"`
	Tax2JurisCode  string                     `json:"tax2_juris_code,omitempty"`
	Tax3Name       string                     `json:"tax3_name,omitempty"`
	Tax3Rate       *float64                   `json:"tax3_rate,omitempty"`
	Tax3JurisType  enum.Tax3JurisType         `json:"tax3_juris_type,omitempty"`
	Tax3JurisName  string                     `json:"tax3_juris_name,omitempty"`
	Tax3JurisCode  string                     `json:"tax3_juris_code,omitempty"`
	Tax4Name       string                     `json:"tax4_name,omitempty"`
	Tax4Rate       *float64                   `json:"tax4_rate,omitempty"`
	Tax4JurisType  enum.Tax4JurisType         `json:"tax4_juris_type,omitempty"`
	Tax4JurisName  string                     `json:"tax4_juris_name,omitempty"`
	Tax4JurisCode  string                     `json:"tax4_juris_code,omitempty"`
	ServiceType    csvTaxRuleEnum.ServiceType `json:"service_type,omitempty"`
	TimeZone       string                     `json:"time_zone,omitempty"`
	ValidFrom      *int64                     `json:"valid_from,omitempty"`
	ValidTill      *int64                     `json:"valid_till,omitempty"`
	Overwrite      *bool                      `json:"overwrite,omitempty"`
}
