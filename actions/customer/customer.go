package customer

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/customer"
	"net/url"
)

func Create(params *customer.CreateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers"), params).WithTelemetryResource("customer").WithTelemetryOperation("create").SetIdempotency(true)
}
func List(params *customer.ListRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers"), params).WithTelemetryResource("customer").WithTelemetryOperation("list")
}
func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v", url.PathEscape(id)), nil).WithTelemetryResource("customer").WithTelemetryOperation("retrieve")
}
func Update(id string, params *customer.UpdateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("update").SetIdempotency(true)
}
func UpdatePaymentMethod(id string, params *customer.UpdatePaymentMethodRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_payment_method", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("updatePaymentMethod").SetIdempotency(true)
}
func UpdateBillingInfo(id string, params *customer.UpdateBillingInfoRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_billing_info", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("updateBillingInfo").SetIdempotency(true)
}
func ContactsForCustomer(id string, params *customer.ContactsForCustomerRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/contacts", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("contactsForCustomer")
}
func AssignPaymentRole(id string, params *customer.AssignPaymentRoleRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/assign_payment_role", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("assignPaymentRole").SetIdempotency(true)
}
func AddContact(id string, params *customer.AddContactRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/add_contact", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("addContact").SetIdempotency(true)
}
func UpdateContact(id string, params *customer.UpdateContactRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_contact", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("updateContact").SetIdempotency(true)
}
func DeleteContact(id string, params *customer.DeleteContactRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_contact", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("deleteContact").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func AddPromotionalCredits(id string, params *customer.AddPromotionalCreditsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/add_promotional_credits", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("addPromotionalCredits").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func DeductPromotionalCredits(id string, params *customer.DeductPromotionalCreditsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/deduct_promotional_credits", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("deductPromotionalCredits").SetIdempotency(true)
}

// Deprecated: This function is deprecated.
func SetPromotionalCredits(id string, params *customer.SetPromotionalCreditsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/set_promotional_credits", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("setPromotionalCredits").SetIdempotency(true)
}
func RecordExcessPayment(id string, params *customer.RecordExcessPaymentRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/record_excess_payment", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("recordExcessPayment").SetIdempotency(true)
}
func CollectPayment(id string, params *customer.CollectPaymentRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/collect_payment", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("collectPayment").SetIdempotency(true)
}
func Delete(id string, params *customer.DeleteRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("delete").SetIdempotency(true)
}
func Move(params *customer.MoveRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/move"), params).WithTelemetryResource("customer").WithTelemetryOperation("move").SetIdempotency(true)
}
func ChangeBillingDate(id string, params *customer.ChangeBillingDateRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/change_billing_date", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("changeBillingDate").SetIdempotency(true)
}
func Merge(params *customer.MergeRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/merge"), params).WithTelemetryResource("customer").WithTelemetryOperation("merge").SetIdempotency(true)
}
func ClearPersonalData(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/clear_personal_data", url.PathEscape(id)), nil).WithTelemetryResource("customer").WithTelemetryOperation("clearPersonalData").SetIdempotency(true)
}
func Relationships(id string, params *customer.RelationshipsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/relationships", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("relationships").SetIdempotency(true)
}
func DeleteRelationship(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_relationship", url.PathEscape(id)), nil).WithTelemetryResource("customer").WithTelemetryOperation("deleteRelationship").SetIdempotency(true)
}
func Hierarchy(id string, params *customer.HierarchyRequestParams) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/hierarchy", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("hierarchy")
}
func ListHierarchyDetail(id string, params *customer.ListHierarchyDetailRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/hierarchy_detail", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("listHierarchyDetail")
}
func UpdateHierarchySettings(id string, params *customer.UpdateHierarchySettingsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_hierarchy_settings", url.PathEscape(id)), params).WithTelemetryResource("customer").WithTelemetryOperation("updateHierarchySettings").SetIdempotency(true)
}
