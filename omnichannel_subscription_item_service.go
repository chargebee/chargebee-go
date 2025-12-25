package chargebee

import (
	"fmt"
	"net/url"
)

type OmnichannelSubscriptionItemService struct {
	config *ClientConfig
}

func (s *OmnichannelSubscriptionItemService) ListOmniSubItemScheduleChanges(id string, req *OmnichannelSubscriptionItemListOmniSubItemScheduleChangesRequest) (*OmnichannelSubscriptionItemListOmniSubItemScheduleChangesResponse, error) {
	req.method = "GET"
	req.path = fmt.Sprintf("/omnichannel_subscription_items/%v/scheduled_changes", url.PathEscape(id))
	req.isListRequest = true
	return send[*OmnichannelSubscriptionItemListOmniSubItemScheduleChangesResponse](req, s.config)
}
