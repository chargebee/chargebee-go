package sitemigrationdetail

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/sitemigrationdetail"
)

func List(params *sitemigrationdetail.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/site_migration_details"), params)
}
