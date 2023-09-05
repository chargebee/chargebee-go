package csvtaxrule

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/csvtaxrule"
)

func Create(params *csvtaxrule.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/csv_tax_rules"), params)
}
