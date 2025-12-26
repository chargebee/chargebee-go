package chargebee

type VirtualBankAccountScheme string

const (
	VirtualBankAccountSchemeAchCredit               VirtualBankAccountScheme = "ach_credit"
	VirtualBankAccountSchemeSepaCredit              VirtualBankAccountScheme = "sepa_credit"
	VirtualBankAccountSchemeUsAutomatedBankTransfer VirtualBankAccountScheme = "us_automated_bank_transfer"
	VirtualBankAccountSchemeGbAutomatedBankTransfer VirtualBankAccountScheme = "gb_automated_bank_transfer"
	VirtualBankAccountSchemeEuAutomatedBankTransfer VirtualBankAccountScheme = "eu_automated_bank_transfer"
	VirtualBankAccountSchemeJpAutomatedBankTransfer VirtualBankAccountScheme = "jp_automated_bank_transfer"
	VirtualBankAccountSchemeMxAutomatedBankTransfer VirtualBankAccountScheme = "mx_automated_bank_transfer"
)

// just struct
type VirtualBankAccount struct {
	Id               string                   `json:"id"`
	CustomerId       string                   `json:"customer_id"`
	Email            string                   `json:"email"`
	Scheme           VirtualBankAccountScheme `json:"scheme"`
	BankName         string                   `json:"bank_name"`
	AccountNumber    string                   `json:"account_number"`
	RoutingNumber    string                   `json:"routing_number"`
	SwiftCode        string                   `json:"swift_code"`
	Gateway          Gateway                  `json:"gateway"`
	GatewayAccountId string                   `json:"gateway_account_id"`
	ResourceVersion  int64                    `json:"resource_version"`
	UpdatedAt        int64                    `json:"updated_at"`
	CreatedAt        int64                    `json:"created_at"`
	ReferenceId      string                   `json:"reference_id"`
	Deleted          bool                     `json:"deleted"`
	Object           string                   `json:"object"`
}

// sub resources
// operations
// input params
type VirtualBankAccountCreateUsingPermanentTokenRequest struct {
	CustomerId  string                   `json:"customer_id"`
	ReferenceId string                   `json:"reference_id"`
	Scheme      VirtualBankAccountScheme `json:"scheme,omitempty"`
	apiRequest  `json:"-" form:"-"`
}

func (r *VirtualBankAccountCreateUsingPermanentTokenRequest) payload() any { return r }

type VirtualBankAccountCreateRequest struct {
	CustomerId string                   `json:"customer_id"`
	Email      string                   `json:"email,omitempty"`
	Scheme     VirtualBankAccountScheme `json:"scheme,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *VirtualBankAccountCreateRequest) payload() any { return r }

type VirtualBankAccountListRequest struct {
	Limit      *int32           `json:"limit,omitempty"`
	Offset     string           `json:"offset,omitempty"`
	CustomerId *StringFilter    `json:"customer_id,omitempty"`
	UpdatedAt  *TimestampFilter `json:"updated_at,omitempty"`
	CreatedAt  *TimestampFilter `json:"created_at,omitempty"`
	apiRequest `json:"-" form:"-"`
}

func (r *VirtualBankAccountListRequest) payload() any { return r }

// operation response
type VirtualBankAccountCreateUsingPermanentTokenResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
	Customer           Customer            `json:"customer,omitempty"`
	apiResponse
}

// operation response
type VirtualBankAccountCreateResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
	Customer           Customer            `json:"customer,omitempty"`
	apiResponse
}

// operation response
type VirtualBankAccountRetrieveResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
	apiResponse
}

// operation sub response
type VirtualBankAccountListVirtualBankAccountResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

// operation response
type VirtualBankAccountListResponse struct {
	List       []*VirtualBankAccountListVirtualBankAccountResponse `json:"list,omitempty"`
	NextOffset string                                              `json:"next_offset,omitempty"`
	apiResponse
}

// operation response
type VirtualBankAccountDeleteResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
	apiResponse
}

// operation response
type VirtualBankAccountDeleteLocalResponse struct {
	VirtualBankAccount *VirtualBankAccount `json:"virtual_bank_account,omitempty"`
	apiResponse
}
