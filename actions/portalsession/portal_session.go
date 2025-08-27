package portalsession

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/portalsession"
	"net/url"
)

func Create(params *portalsession.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions"), params).SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/portal_sessions/%v", url.PathEscape(id)), nil)
}
func Logout(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/logout", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Activate(id string, params *portalsession.ActivateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/activate", url.PathEscape(id)), params).SetIdempotency(true)
}
