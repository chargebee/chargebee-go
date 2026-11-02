package retrysegment

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/retrysegment"
	"net/url"
)

func Retrieve(id string, params *retrysegment.RetrieveRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/retry_segments/%v", url.PathEscape(id)), params)
}
