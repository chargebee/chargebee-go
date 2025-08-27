package resourcemigration

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/resourcemigration"
)

func RetrieveLatest(params *resourcemigration.RetrieveLatestRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/resource_migrations/retrieve_latest"), params)
}
