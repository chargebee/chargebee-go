package session

import (
	"encoding/json"
)

type Session struct {
	Id        string 		  `json:"id"`
	Content   json.RawMessage `json:"content"`
	CreatedAt int64           `json:"created_at"`
	ExpiresAt int64           `json:"expires_at"`
	Object    string          `json:"object"`
}
type CreateRequestParams struct {
	Customer     *CreateCustomerParams     `json:"customer,omitempty"`
	Subscription *CreateSubscriptionParams `json:"subscription,omitempty"`
}
type CreateCustomerParams struct {
	Id string `json:"id,omitempty"`
}
type CreateSubscriptionParams struct {
	Id string `json:"id,omitempty"`
}
type RetrieveRequestParams struct {
	Id string `json:"id,omitempty"`
}
