package customer

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/customer"
	"net/url"
)

func Create(params *customer.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers"), params)
}
func List(params *customer.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v", url.PathEscape(id)), nil)
}
func Update(id string, params *customer.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v", url.PathEscape(id)), params)
}
func UpdatePaymentMethod(id string, params *customer.UpdatePaymentMethodRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_payment_method", url.PathEscape(id)), params)
}
func UpdateBillingInfo(id string, params *customer.UpdateBillingInfoRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_billing_info", url.PathEscape(id)), params)
}
func ContactsForCustomer(id string, params *customer.ContactsForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/contacts", url.PathEscape(id)), params)
}
func AssignPaymentRole(id string, params *customer.AssignPaymentRoleRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/assign_payment_role", url.PathEscape(id)), params)
}
func AddContact(id string, params *customer.AddContactRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/add_contact", url.PathEscape(id)), params)
}
func UpdateContact(id string, params *customer.UpdateContactRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_contact", url.PathEscape(id)), params)
}
func DeleteContact(id string, params *customer.DeleteContactRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_contact", url.PathEscape(id)), params)
}
func AddPromotionalCredits(id string, params *customer.AddPromotionalCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/add_promotional_credits", url.PathEscape(id)), params)
}
func DeductPromotionalCredits(id string, params *customer.DeductPromotionalCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/deduct_promotional_credits", url.PathEscape(id)), params)
}
func SetPromotionalCredits(id string, params *customer.SetPromotionalCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/set_promotional_credits", url.PathEscape(id)), params)
}
func RecordExcessPayment(id string, params *customer.RecordExcessPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/record_excess_payment", url.PathEscape(id)), params)
}
func CollectPayment(id string, params *customer.CollectPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/collect_payment", url.PathEscape(id)), params)
}
func Delete(id string, params *customer.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete", url.PathEscape(id)), params)
}
func Move(params *customer.MoveRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/move"), params)
}
func ChangeBillingDate(id string, params *customer.ChangeBillingDateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/change_billing_date", url.PathEscape(id)), params)
}
func Merge(params *customer.MergeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/merge"), params)
}
func ClearPersonalData(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/clear_personal_data", url.PathEscape(id)), nil)
}
func Relationships(id string, params *customer.RelationshipsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/relationships", url.PathEscape(id)), params)
}
func DeleteRelationship(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_relationship", url.PathEscape(id)), nil)
}
func Hierarchy(id string, params *customer.HierarchyRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/hierarchy", url.PathEscape(id)), params)
}
func UpdateHierarchySettings(id string, params *customer.UpdateHierarchySettingsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_hierarchy_settings", url.PathEscape(id)), params)
}
