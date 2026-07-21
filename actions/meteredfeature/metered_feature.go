package meteredfeature

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/meteredfeature"
	"net/url"
)

func Create(params *meteredfeature.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/metered_features"), params).SetIdempotency(true)
}
func Archive(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/metered_features/%v/archive_command", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Reactivate(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/metered_features/%v/reactivate_command", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/metered_features/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
