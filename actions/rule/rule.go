package rule

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"net/url"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/rules/%v", url.PathEscape(id)), nil)
}
