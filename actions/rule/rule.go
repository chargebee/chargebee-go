package rule

import (
	"fmt"
	"net/url"

	"github.com/chargebee/chargebee-go/v3"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/rules/%v", url.PathEscape(id)), nil)
}
