package businessrule

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/businessrule"
	"net/url"
)

func Create(params *businessrule.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules"), params).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func UpdateDraft(id string, params *businessrule.UpdateDraftRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules/%v/draft", url.PathEscape(id)), params).SetIdempotency(true)
}
func List(params *businessrule.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/business_rules"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/business_rules/%v", url.PathEscape(id)), nil)
}
func RetrieveDraft(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/business_rules/%v/draft", url.PathEscape(id)), nil)
}
func DeleteDraft(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules/%v/delete_draft", url.PathEscape(id)), nil).SetIdempotency(true)
}
func ActivateRule(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules/%v/activate", url.PathEscape(id)), nil).SetIdempotency(true)
}
func DeactivateRule(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules/%v/deactivate", url.PathEscape(id)), nil).SetIdempotency(true)
}
func ReleaseRule(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules/%v/release", url.PathEscape(id)), nil).SetIdempotency(true)
}
func ApplyRules(params *businessrule.ApplyRulesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rules/apply_rules"), params).SetIdempotency(true)
}
