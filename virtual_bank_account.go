package chargebee

type Scheme string

const (
	SchemeAchCredit               Scheme = "ach_credit"
	SchemeSepaCredit              Scheme = "sepa_credit"
	SchemeUsAutomatedBankTransfer Scheme = "us_automated_bank_transfer"
	SchemeGbAutomatedBankTransfer Scheme = "gb_automated_bank_transfer"
	SchemeEuAutomatedBankTransfer Scheme = "eu_automated_bank_transfer"
	SchemeJpAutomatedBankTransfer Scheme = "jp_automated_bank_transfer"
	SchemeMxAutomatedBankTransfer Scheme = "mx_automated_bank_transfer"
)

type VirtualBankAccount struct {
	Id               string       `json:"id"`
	CustomerId       string       `json:"customer_id"`
	Email            string       `json:"email"`
	Scheme           Scheme       `json:"scheme"`
	BankName         string       `json:"bank_name"`
	AccountNumber    string       `json:"account_number"`
	RoutingNumber    string       `json:"routing_number"`
	SwiftCode        string       `json:"swift_code"`
	Gateway          enum.Gateway `json:"gateway"`
	GatewayAccountId string       `json:"gateway_account_id"`
	ResourceVersion  int64        `json:"resource_version"`
	UpdatedAt        int64        `json:"updated_at"`
	CreatedAt        int64        `json:"created_at"`
	ReferenceId      string       `json:"reference_id"`
	Deleted          bool         `json:"deleted"`
	Object           string       `json:"object"`
}
type CreateUsingPermanentTokenRequest struct {
	CustomerId  string `json:"customer_id"`
	ReferenceId string `json:"reference_id"`
	Scheme      Scheme `json:"scheme,omitempty"`
}
type CreateRequest struct {
	CustomerId string `json:"customer_id"`
	Email      string `json:"email,omitempty"`
	Scheme     Scheme `json:"scheme,omitempty"`
}
type ListRequest struct {
	Limit      *int32                  `json:"limit,omitempty"`
	Offset     string                  `json:"offset,omitempty"`
	CustomerId *filter.StringFilter    `json:"customer_id,omitempty"`
	UpdatedAt  *filter.TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt  *filter.TimestampFilter `json:"created_at,omitempty"`
}

type CreateUsingPermanentTokenResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
	Customer           *customer.Customer  `json:"customer,omitempty"`
}

type CreateResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
	Customer           *customer.Customer  `json:"customer,omitempty"`
}

type RetrieveResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type ListVirtualBankAccountResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type ListResponse struct {
	List       []*ListVirtualBankAccountResponse `json:"list,omitempty"`
	NextOffset string                            `json:"next_offset,omitempty"`
}

type DeleteResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type DeleteLocalResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}
