package ledgerentry

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	ledgerEntryEnum "github.com/chargebee/chargebee-go/v3/models/ledgerentry/enum"
)

type LedgerEntry struct {
	Id                     string                      `json:"id"`
	SubscriptionId         string                      `json:"subscription_id"`
	AccountType            ledgerEntryEnum.AccountType `json:"account_type"`
	UnitId                 string                      `json:"unit_id"`
	UnitType               ledgerEntryEnum.UnitType    `json:"unit_type"`
	Amount                 string                      `json:"amount"`
	GrantBlockStartBalance string                      `json:"grant_block_start_balance"`
	GrantBlockEndBalance   string                      `json:"grant_block_end_balance"`
	AccountStartBalance    string                      `json:"account_start_balance"`
	AccountEndBalance      string                      `json:"account_end_balance"`
	Type                   enum.Type                   `json:"type"`
	LedgerOperationId      string                      `json:"ledger_operation_id"`
	GrantBlockId           string                      `json:"grant_block_id"`
	CreatedAt              int64                       `json:"created_at"`
	ModifiedAt             int64                       `json:"modified_at"`
	Object                 string                      `json:"object"`
}
