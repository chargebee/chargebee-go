package export

import (
	"errors"
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/export"
	exportEnum "github.com/chargebee/chargebee-go/v3/models/export/enum"
	"net/url"
	"time"
)

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/exports/%v", url.PathEscape(id)), nil)
}
func RevenueRecognition(params *export.RevenueRecognitionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/revenue_recognition"), params).SetIdempotency(true)
}
func DeferredRevenue(params *export.DeferredRevenueRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/deferred_revenue"), params).SetIdempotency(true)
}
func Plans(params *export.PlansRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/plans"), params).SetIdempotency(true)
}
func Addons(params *export.AddonsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/addons"), params).SetIdempotency(true)
}
func Coupons(params *export.CouponsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/coupons"), params).SetIdempotency(true)
}
func Customers(params *export.CustomersRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/customers"), params).SetIdempotency(true)
}
func Subscriptions(params *export.SubscriptionsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/subscriptions"), params).SetIdempotency(true)
}
func Invoices(params *export.InvoicesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/invoices"), params).SetIdempotency(true)
}
func CreditNotes(params *export.CreditNotesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/credit_notes"), params).SetIdempotency(true)
}
func Transactions(params *export.TransactionsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/transactions"), params).SetIdempotency(true)
}
func Orders(params *export.OrdersRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/orders"), params).SetIdempotency(true)
}
func ItemFamilies(params *export.ItemFamiliesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/item_families"), params).SetIdempotency(true)
}
func Items(params *export.ItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/items"), params).SetIdempotency(true)
}
func ItemPrices(params *export.ItemPricesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/item_prices"), params).SetIdempotency(true)
}
func AttachedItems(params *export.AttachedItemsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/attached_items"), params).SetIdempotency(true)
}
func DifferentialPrices(params *export.DifferentialPricesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/differential_prices"), params).SetIdempotency(true)
}
func PriceVariants(params *export.PriceVariantsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/exports/price_variants"), params).SetIdempotency(true)
}
func WaitForExportCompletion(exp export.Export) (export.Export, error) {
	return WaitForExportCompletionWithEnv(exp, chargebee.DefaultConfig())
}
func WaitForExportCompletionWithEnv(exp export.Export, env chargebee.Environment) (export.Export, error) {
	count := 0
	for exp.Status == exportEnum.StatusInProcess {
		if count > 30 {
			return exp, errors.New("'Export is taking too long'")
		}
		count++
		time.Sleep(chargebee.ExportWaitInSecs)
		result, err := Retrieve(exp.Id).RequestWithEnv(env)
		if err != nil {
			return exp, err
		}
		exp = *result.Export
	}
	return exp, nil
}
