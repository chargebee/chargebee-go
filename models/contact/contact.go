package contact

type Contact struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Label            string `json:"label"`
	Enabled          bool   `json:"enabled"`
	SendAccountEmail bool   `json:"send_account_email"`
	SendBillingEmail bool   `json:"send_billing_email"`
	Object           string `json:"object"`
}
