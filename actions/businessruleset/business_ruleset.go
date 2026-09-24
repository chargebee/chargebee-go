package businessruleset

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/businessruleset"
	"net/url"
)

func Create(params *businessruleset.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rulesets"), params).SetIdempotency(true)
}
func Update(id string, params *businessruleset.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rulesets/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Delete(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rulesets/%v/delete", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Activate(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rulesets/%v/activate", url.PathEscape(id)), nil).SetIdempotency(true)
}
func Deactivate(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rulesets/%v/deactivate", url.PathEscape(id)), nil).SetIdempotency(true)
}
func AddRules(id string, params *businessruleset.AddRulesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rulesets/%v/add_rules", url.PathEscape(id)), params).SetIdempotency(true)
}
func RemoveRules(id string, params *businessruleset.RemoveRulesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/business_rulesets/%v/remove_rules", url.PathEscape(id)), params).SetIdempotency(true)
}
func ListRules(id string, params *businessruleset.ListRulesRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/business_rulesets/%v/rules", url.PathEscape(id)), params)
}
func List(params *businessruleset.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/business_rulesets"), params)
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/business_rulesets/%v", url.PathEscape(id)), nil)
}
