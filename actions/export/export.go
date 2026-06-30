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

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/exports/%v", url.PathEscape(id)), nil).WithTelemetryResource("export").WithTelemetryOperation("retrieve")
}
func RevenueRecognition(params *export.RevenueRecognitionRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/revenue_recognition"), params).WithTelemetryResource("export").WithTelemetryOperation("revenueRecognition").SetIdempotency(true)
}
func DeferredRevenue(params *export.DeferredRevenueRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/deferred_revenue"), params).WithTelemetryResource("export").WithTelemetryOperation("deferredRevenue").SetIdempotency(true)
}
func Plans(params *export.PlansRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/plans"), params).WithTelemetryResource("export").WithTelemetryOperation("plans").SetIdempotency(true)
}
func Addons(params *export.AddonsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/addons"), params).WithTelemetryResource("export").WithTelemetryOperation("addons").SetIdempotency(true)
}
func Coupons(params *export.CouponsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/coupons"), params).WithTelemetryResource("export").WithTelemetryOperation("coupons").SetIdempotency(true)
}
func Customers(params *export.CustomersRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/customers"), params).WithTelemetryResource("export").WithTelemetryOperation("customers").SetIdempotency(true)
}
func Subscriptions(params *export.SubscriptionsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/subscriptions"), params).WithTelemetryResource("export").WithTelemetryOperation("subscriptions").SetIdempotency(true)
}
func Invoices(params *export.InvoicesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/invoices"), params).WithTelemetryResource("export").WithTelemetryOperation("invoices").SetIdempotency(true)
}
func CreditNotes(params *export.CreditNotesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/credit_notes"), params).WithTelemetryResource("export").WithTelemetryOperation("creditNotes").SetIdempotency(true)
}
func Transactions(params *export.TransactionsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/transactions"), params).WithTelemetryResource("export").WithTelemetryOperation("transactions").SetIdempotency(true)
}
func Orders(params *export.OrdersRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/orders"), params).WithTelemetryResource("export").WithTelemetryOperation("orders").SetIdempotency(true)
}
func ItemFamilies(params *export.ItemFamiliesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/item_families"), params).WithTelemetryResource("export").WithTelemetryOperation("itemFamilies").SetIdempotency(true)
}
func Items(params *export.ItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/items"), params).WithTelemetryResource("export").WithTelemetryOperation("items").SetIdempotency(true)
}
func ItemPrices(params *export.ItemPricesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/item_prices"), params).WithTelemetryResource("export").WithTelemetryOperation("itemPrices").SetIdempotency(true)
}
func AttachedItems(params *export.AttachedItemsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/attached_items"), params).WithTelemetryResource("export").WithTelemetryOperation("attachedItems").SetIdempotency(true)
}
func DifferentialPrices(params *export.DifferentialPricesRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/differential_prices"), params).WithTelemetryResource("export").WithTelemetryOperation("differentialPrices").SetIdempotency(true)
}
func PriceVariants(params *export.PriceVariantsRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/exports/price_variants"), params).WithTelemetryResource("export").WithTelemetryOperation("priceVariants").SetIdempotency(true)
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
