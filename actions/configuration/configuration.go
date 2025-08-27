package configuration

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
)

func List() chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/configurations"), nil)
}
