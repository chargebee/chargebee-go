package emaillog

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/emaillog"
	"net/url"
)

func EmailLogsForCustomer(id string, params *emaillog.EmailLogsForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/email_logs", url.PathEscape(id)), params)
}
