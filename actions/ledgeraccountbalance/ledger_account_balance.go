package ledgeraccountbalance

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/ledgeraccountbalance"
)

func ListLedgerAccountBalances(params *ledgeraccountbalance.ListLedgerAccountBalancesRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/ledger_account_balances"), params)
}
