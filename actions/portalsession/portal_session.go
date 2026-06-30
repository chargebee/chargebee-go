package portalsession

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/portalsession"
	"net/url"
)

func Create(params *portalsession.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions"), params).WithTelemetryResource("portalSession").WithTelemetryOperation("create").SetIdempotency(true)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/portal_sessions/%v", url.PathEscape(id)), nil).WithTelemetryResource("portalSession").WithTelemetryOperation("retrieve")
}
func Logout(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/logout", url.PathEscape(id)), nil).WithTelemetryResource("portalSession").WithTelemetryOperation("logout").SetIdempotency(true)
}
func Activate(id string, params *portalsession.ActivateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/portal_sessions/%v/activate", url.PathEscape(id)), params).WithTelemetryResource("portalSession").WithTelemetryOperation("activate").SetIdempotency(true)
}
