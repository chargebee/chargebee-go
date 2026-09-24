package emaillog

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
)

type EmailLog struct {
	Id               string      `json:"id"`
	TemplateName     string      `json:"template_name"`
	FromAddress      string      `json:"from_address"`
	ToAddress        string      `json:"to_address"`
	Subject          string      `json:"subject"`
	Status           enum.Status `json:"status"`
	SentOn           int64       `json:"sent_on"`
	CustomerId       string      `json:"customer_id"`
	SiteId           string      `json:"site_id"`
	BusinessEntityId string      `json:"business_entity_id"`
	BrandId          string      `json:"brand_id"`
	ErrorMessage     string      `json:"error_message"`
	Object           string      `json:"object"`
}
type EmailLogsForCustomerRequestParams struct {
	Limit            *int32                  `json:"limit,omitempty"`
	Offset           string                  `json:"offset,omitempty"`
	SentOn           *filter.TimestampFilter `json:"sent_on,omitempty"`
	BusinessEntityId *filter.StringFilter    `json:"business_entity_id,omitempty"`
	BrandId          *filter.StringFilter    `json:"brand_id,omitempty"`
}
