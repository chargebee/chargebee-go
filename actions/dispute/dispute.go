package dispute

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/dispute"
	"net/url"
)

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/disputes/%v", url.PathEscape(id)), nil)
}
func List(params *dispute.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/disputes"), params)
}
