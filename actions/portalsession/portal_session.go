package portalsession

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/portalsession"
)

func Create(params *portalsession.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/portal_sessions/%v", chargebee.IDEscape(id)), nil)
}
func Logout(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/logout", chargebee.IDEscape(id)), nil)
}
func Activate(id string, params *portalsession.ActivateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/activate", chargebee.IDEscape(id)), params)
}
