package address

import (
	"github.com/chargebee/chargebee-go/enum"
)

type Address struct {
	Label            string                `json:"label"`
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Addr             string                `json:"addr"`
	ExtendedAddr     string                `json:"extended_addr"`
	ExtendedAddr2    string                `json:"extended_addr2"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	SubscriptionId   string                `json:"subscription_id"`
	Object           string                `json:"object"`
}
type RetrieveRequestParams struct {
	SubscriptionId string `json:"subscription_id"`
	Label          string `json:"label"`
}
type UpdateRequestParams struct {
	SubscriptionId   string                `json:"subscription_id"`
	Label            string                `json:"label"`
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Addr             string                `json:"addr,omitempty"`
	ExtendedAddr     string                `json:"extended_addr,omitempty"`
	ExtendedAddr2    string                `json:"extended_addr2,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
