package customer

import (
	"fmt"

	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/customer"
)

func Create(params *customer.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers"), params)
}
func List(params *customer.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v", chargebee.IDEscape(id)), nil)
}
func Update(id string, params *customer.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v", chargebee.IDEscape(id)), params)
}
func UpdatePaymentMethod(id string, params *customer.UpdatePaymentMethodRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_payment_method", chargebee.IDEscape(id)), params)
}
func UpdateBillingInfo(id string, params *customer.UpdateBillingInfoRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_billing_info", chargebee.IDEscape(id)), params)
}
func ContactsForCustomer(id string, params *customer.ContactsForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/contacts", chargebee.IDEscape(id)), params)
}
func AssignPaymentRole(id string, params *customer.AssignPaymentRoleRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/assign_payment_role", chargebee.IDEscape(id)), params)
}
func AddContact(id string, params *customer.AddContactRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/add_contact", chargebee.IDEscape(id)), params)
}
func UpdateContact(id string, params *customer.UpdateContactRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_contact", chargebee.IDEscape(id)), params)
}
func DeleteContact(id string, params *customer.DeleteContactRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_contact", chargebee.IDEscape(id)), params)
}
func AddPromotionalCredits(id string, params *customer.AddPromotionalCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/add_promotional_credits", chargebee.IDEscape(id)), params)
}
func DeductPromotionalCredits(id string, params *customer.DeductPromotionalCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/deduct_promotional_credits", chargebee.IDEscape(id)), params)
}
func SetPromotionalCredits(id string, params *customer.SetPromotionalCreditsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/set_promotional_credits", chargebee.IDEscape(id)), params)
}
func RecordExcessPayment(id string, params *customer.RecordExcessPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/record_excess_payment", chargebee.IDEscape(id)), params)
}
func CollectPayment(id string, params *customer.CollectPaymentRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/collect_payment", chargebee.IDEscape(id)), params)
}
func Delete(id string, params *customer.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete", chargebee.IDEscape(id)), params)
}
func Move(params *customer.MoveRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/move"), params)
}
func ChangeBillingDate(id string, params *customer.ChangeBillingDateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/change_billing_date", chargebee.IDEscape(id)), params)
}
func Merge(params *customer.MergeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/merge"), params)
}
func ClearPersonalData(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/clear_personal_data", chargebee.IDEscape(id)), nil)
}
func Relationships(id string, params *customer.RelationshipsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/relationships", chargebee.IDEscape(id)), params)
}
func DeleteRelationship(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/delete_relationship", chargebee.IDEscape(id)), nil)
}
func Hierarchy(id string, params *customer.HierarchyRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/customers/%v/hierarchy", chargebee.IDEscape(id)), params)
}
func UpdateHierarchySettings(id string, params *customer.UpdateHierarchySettingsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/update_hierarchy_settings", chargebee.IDEscape(id)), params)
}
