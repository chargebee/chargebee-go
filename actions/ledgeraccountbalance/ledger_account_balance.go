package ledgeraccountbalance

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/ledgeraccountbalance"
)

func ListLedgerAccountBalances(params *ledgeraccountbalance.ListLedgerAccountBalancesRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/ledger_account_balances"), params)
}
func UpdateOverdraftLimit(params *ledgeraccountbalance.UpdateOverdraftLimitRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/ledger_account_balances/update_overdraft_limit"), params)
}
