package chargebee

import (
	"fmt"
	"net/url"
)

type OmnichannelSubscriptionItemService struct {
	transport *Transport
}

func (s *OmnichannelSubscriptionItemService) ListOmniSubItemScheduleChanges(id string, req *ListOmniSubItemScheduleChangesRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/omnichannel_subscription_items/%v/scheduled_changes", url.PathEscape(id)), req)
}
