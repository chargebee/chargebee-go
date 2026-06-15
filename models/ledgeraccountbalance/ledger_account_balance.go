package ledgeraccountbalance

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	ledgerAccountBalanceEnum "github.com/chargebee/chargebee-go/v3/models/ledgeraccountbalance/enum"
)

type LedgerAccountBalance struct {
	SubscriptionId     string                            `json:"subscription_id"`
	UnitId             string                            `json:"unit_id"`
	UnitType           ledgerAccountBalanceEnum.UnitType `json:"unit_type"`
	ModifiedAt         int64                             `json:"modified_at"`
	ProvisionedBalance *ProvisionedBalance               `json:"provisioned_balance"`
	OverdraftBalance   *OverdraftBalance                 `json:"overdraft_balance"`
	Object             string                            `json:"object"`
}
type ProvisionedBalance struct {
	TotalBalance  string `json:"total_balance"`
	UsableBalance string `json:"usable_balance"`
	HoldAmount    string `json:"hold_amount"`
	Object        string `json:"object"`
}
type OverdraftBalance struct {
	IsUnlimited   bool   `json:"is_unlimited"`
	Limit         string `json:"limit"`
	TotalBalance  string `json:"total_balance"`
	UsableBalance string `json:"usable_balance"`
	UsedAmount    string `json:"used_amount"`
	HoldAmount    string `json:"hold_amount"`
	Object        string `json:"object"`
}
type ListLedgerAccountBalancesRequestParams struct {
	Limit          *int32               `json:"limit,omitempty"`
	Offset         string               `json:"offset,omitempty"`
	SubscriptionId *filter.StringFilter `json:"subscription_id"`
	UnitId         *filter.StringFilter `json:"unit_id,omitempty"`
}
