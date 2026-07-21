package meter

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/meter"
)

func List(params *meter.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/meters"), params)
}
