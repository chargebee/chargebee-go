package csvtaxrule

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/csvtaxrule"
)

func Create(params *csvtaxrule.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/csv_tax_rules"), params)
}
