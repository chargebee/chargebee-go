package omnichannelsubscriptionitem

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"
	"net/url"
)

func ListOmniSubItemScheduleChanges(id string, params *omnichannelsubscriptionitem.ListOmniSubItemScheduleChangesRequestParams) chargebee.ListRequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/omnichannel_subscription_items/%v/scheduled_changes", url.PathEscape(id)), params)
}
