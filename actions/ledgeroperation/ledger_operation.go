package ledgeroperation

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/ledgeroperation"
	"net/url"
)

func RetrieveLedgerOperation(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/ledger_operations/%v", url.PathEscape(id)), nil)
}
func ListLedgerOperations(params *ledgeroperation.ListLedgerOperationsRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/ledger_operations"), params)
}
func Capture(params *ledgeroperation.CaptureRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/ledger_operations/capture"), params)
}
func Authorize(params *ledgeroperation.AuthorizeRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/ledger_operations/authorize"), params)
}
func CaptureAuthorization(params *ledgeroperation.CaptureAuthorizationRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/ledger_operations/capture_authorization"), params)
}
func ReleaseAuthorization(params *ledgeroperation.ReleaseAuthorizationRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/ledger_operations/release_authorization"), params)
}
