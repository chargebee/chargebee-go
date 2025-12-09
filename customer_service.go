package chargebee

import (
	"fmt"
	"net/url"
)

type CustomerService struct {
	transport *Transport
}

func (s *CustomerService) Create(req *CreateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers"), req).SetIdempotency(true)
}

func (s *CustomerService) List(req *ListRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers"), req)
}

func (s *CustomerService) Retrieve(id string) Request {
	return s.transport.Send("GET", fmt.Sprintf("/customers/%v", url.PathEscape(id)), nil)
}

func (s *CustomerService) Update(id string, req *UpdateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) UpdatePaymentMethod(id string, req *UpdatePaymentMethodRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/update_payment_method", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) UpdateBillingInfo(id string, req *UpdateBillingInfoRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/update_billing_info", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) ContactsForCustomer(id string, req *ContactsForCustomerRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/contacts", url.PathEscape(id)), req)
}

func (s *CustomerService) AssignPaymentRole(id string, req *AssignPaymentRoleRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/assign_payment_role", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) AddContact(id string, req *AddContactRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/add_contact", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) UpdateContact(id string, req *UpdateContactRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/update_contact", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) DeleteContact(id string, req *DeleteContactRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/delete_contact", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) AddPromotionalCredits(id string, req *AddPromotionalCreditsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/add_promotional_credits", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) DeductPromotionalCredits(id string, req *DeductPromotionalCreditsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/deduct_promotional_credits", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) SetPromotionalCredits(id string, req *SetPromotionalCreditsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/set_promotional_credits", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) RecordExcessPayment(id string, req *RecordExcessPaymentRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/record_excess_payment", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) CollectPayment(id string, req *CollectPaymentRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/collect_payment", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) Delete(id string, req *DeleteRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/delete", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) Move(req *MoveRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/move"), req).SetIdempotency(true)
}

func (s *CustomerService) ChangeBillingDate(id string, req *ChangeBillingDateRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/change_billing_date", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) Merge(req *MergeRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/merge"), req).SetIdempotency(true)
}

func (s *CustomerService) ClearPersonalData(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/clear_personal_data", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *CustomerService) Relationships(id string, req *RelationshipsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/relationships", url.PathEscape(id)), req).SetIdempotency(true)
}

func (s *CustomerService) DeleteRelationship(id string) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/delete_relationship", url.PathEscape(id)), nil).SetIdempotency(true)
}

func (s *CustomerService) Hierarchy(id string, req *HierarchyRequest) Request {
	return s.transport.Send("GET", fmt.Sprintf("/customers/%v/hierarchy", url.PathEscape(id)), req)
}

func (s *CustomerService) ListHierarchyDetail(id string, req *ListHierarchyDetailRequest) ListRequest {
	return s.transport.SendList("GET", fmt.Sprintf("/customers/%v/hierarchy_detail", url.PathEscape(id)), req)
}

func (s *CustomerService) UpdateHierarchySettings(id string, req *UpdateHierarchySettingsRequest) Request {
	return s.transport.Send("POST", fmt.Sprintf("/customers/%v/update_hierarchy_settings", url.PathEscape(id)), req).SetIdempotency(true)
}
