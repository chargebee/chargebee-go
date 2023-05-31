# Go Client Library for Chargebee API
***

This is the source code for the Go client library for [Chargebee APIs](https://apidocs.chargebee.com/docs/api?lang=go).

**Go version**: v3 and v2 of the library require Go v1.3 or higher.

## Library versions
***

The versioning scheme of this library is inspired by [SemVer](https://semver.org/) and the format is `v{MAJOR}.{MINOR}.{PATCH}`. For example, `v3.0.0` and `v2.5.1` are valid library versions.

The following table provides some details for each major version:

| Library major version | Status   | Compatible API versions                                                                             | **Branch**        |
|----------------------------|----------|-----------------------------------------------------------------------------------------------------|---------------|
| v3                         | Active   | [v2](https://apidocs.chargebee.com/docs/api/v2?lang=go) and [v1](https://apidocs.chargebee.com/docs/api/v1?lang=go) | `master`      |
| v2                         | Active   | [v2](https://apidocs.chargebee.com/docs/api/v2?lang=go) and [v1](https://apidocs.chargebee.com/docs/api/v1?lang=go) | `chargebee-v2`|


A couple of terms used in the above table are explained below:
- **Status**: The current development status for the library version. An **Active** major version is currently being maintained and continues to get backward-compatible changes.
- **Branch**: The branch in this repository containing the source code for the latest release of the library version. Every version of the library has been [tagged](https://github.com/chargebee/chargebee-go/tags). You can check out the source code for any version using its tag.

🔴 **Attention**: The support for v2 will eventually be discontinued on **December 31st 2023** and will no longer receive any further updates. We strongly recommend [upgrading to v3](https://github.com/chargebee/chargebee-go/wiki/Migration-guide-for-v3) as soon as possible.

**Note:** See the [changelog](CHANGELOG.md) for a history of changes.

## Install the library
***

Install the latest version of the library with the following commands:

### Install v3
``` shell
go get github.com/chargebee/chargebee-go/v3
```

### Install v2
``` shell
go get github.com/chargebee/chargebee-go
```


## Use the library
***

Some examples for using the library are listed below.

### Create a customer and subscription

```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go/v3"
  subscriptionAction "github.com/chargebee/chargebee-go/v3/actions/subscription"
  "github.com/chargebee/chargebee-go/v3/models/subscription"
)

func main() {
  chargebee.Configure("{site_api_key}", "{site}")
  res, err := subscriptionAction.Create(&subscription.CreateRequestParams{
    PlanId:         "cbdemo_grow",
    BillingCycles:  chargebee.Int32(3),
    AutoCollection: enum.AutoCollectionOff,
    Customer: &subscription.CreateCustomerParams{
      Email:          "john@user.com",
      FirstName:      "John",
      LastName:       "Doe",
      Locale:         "fr-CA",
      Phone:          "+1-949-999-9999",
      AutoCollection: enum.AutoCollectionOff,
    }}).Request()
  if err != nil {
    panic(err)
  }else{
     Subscription := res.Subscription
     Customer := res.Customer
     Invoice := res.Invoice
  }
}
```

### Create a subscription with addons, metadata, and coupons

```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go/v3"
  subscriptionAction "github.com/chargebee/chargebee-go/v3/actions/subscription"
  "github.com/chargebee/chargebee-go/v3/models/subscription"
)

func main() {
  chargebee.Configure("{site_api_key}", "{site}")
  res, err := subscriptionAction.Create(&subscription.CreateRequestParams{
    PlanId:         "cbdemo_grow",
    BillingCycles:  chargebee.Int32(3),
    AutoCollection: enum.AutoCollectionOff,
    Customer: &subscription.CreateCustomerParams{
      Email:          "john@user.com",
      FirstName:      "John",
      LastName:       "Doe",
      Locale:         "fr-CA",
      Phone:          "+1-949-999-9999",
      AutoCollection: enum.AutoCollectionOff,
    },
    BillingAddress: &subscription.CreateBillingAddressParams{
      FirstName: "John",
      LastName:  "Doe",
      Line1:     "PO Box 9999",
      City:      "Walnut",
      State:     "California",
      Zip:       "91789",
      Country:   "US",
    },
    MetaData: map[string]interface{}{
      "features": map[string]interface{}{
        "usage-limit":        "5GB",
        "speed-within-quota": "2MBbps",
        "post-usage-quota":   "512kbps",
      },
    },
    Addons: []*subscription.CreateAddonParams{
      {
        Id: "cbdemo_conciergesupport",
      },
      {
        Id:       "cbdemo_additionaluser",
        Quantity: chargebee.Int32(2),
      },
    },
    CouponIds: []string{"cbdemo_earlybird"},
  }).Request()
  if err != nil {
    panic(err)
  }else{
  Subscription := res.Subscription
  Customer := res.Customer
  Card := res.Card
  Invoice := res.Invoice
  UnbilledCharges := res.UnbilledCharges
  }
}
```

### Create a subscription with custom headers and custom fields and custom context

```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go/v3"
  subscriptionAction "github.com/chargebee/chargebee-go/v3/actions/subscription"
  "github.com/chargebee/chargebee-go/v3/models/subscription"
)

func main() {
  chargebee.Configure("{site_api_key}", "{site}")
  res, err := subscriptionAction.Create(&subscription.CreateRequestParams{
    PlanId: "cbdemo_grow",
  }).Headers("chargebee-request-origin-ip", "192.168.1.2").Contexts(ctx).AddParams("cf_gender","Female").Request() // Customer level custom field. 
  if err != nil {
    panic(err)
  }else{
  Subscription := res.Subscription
  Customer := res.Customer
  Card := res.Card
  Invoice := res.Invoice
  UnbilledCharges := res.UnbilledCharges
  }
}
```

### Retrieve a filtered list of subscriptions

```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go/v3"
  subscriptionAction "github.com/chargebee/chargebee-go/v3/actions/subscription"
  "github.com/chargebee/chargebee-go/v3/filter"
  "github.com/chargebee/chargebee-go/v3/models/subscription"
)

func main() {
  chargebee.Configure("{site_api_key}", "{site}")
  res, err := subscriptionAction.List(&subscription.ListRequestParams{
    Limit: chargebee.Int32(5),
    Id: &filter.StringFilter{
      In: []string{"cbdemo_john-sub", "cbdemo_ricky-sub"},
    },
    PlanId: &filter.StringFilter{
      IsNot: "basic",
    },
    Status: &filter.EnumFilter{
      Is: subscriptionEnum.StatusActive,
    },
    SortBy: &filter.SortFilter{
      Asc: "created_at",
    },
  }).ListRequest()
  if err != nil {
    panic(err)
  }else{
  for i := range res.List {
    Subscription := res.List[i].Subscription
    Customer := res.List[i].Customer
    Card := res.List[i].Card
  }
  }
}
```

### Create an idempotent request

[Idempotency keys](https://apidocs.chargebee.com/docs/api/idempotency?prod_cat_ver=2) are passed along with request headers to allow a safe retry of POST requests. 

```go
import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	customerAction "github.com/chargebee/chargebee-go/v3/actions/customer"
	"github.com/chargebee/chargebee-go/v3/models/customer"
)

func main() {
    chargebee.Configure("{site_api_key}", "{site}")
	res, err := customerAction.Create(&customer.CreateRequestParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
	})
	.SetIdempotencyKey("ghggh") // Replace <<UUID>> with a unique string
	.Request()
	if err != nil {
		fmt.Println(err)
	} else {
		Customer := res.Customer
		fmt.Println(Customer)
	}
  headerValue := res.GetResponseHeaders() // Retrieves response headers
	fmt.Println(headerValue)
  idempotencyReplayedValue := res.IsIdempotencyReplayed()// Retrieves idempotency replayed header value 
  fmt.Println(idempotencyReplayedValue)
}
```
`IsIdempotencyReplayed()` method can be accessed to differentiate between original and replayed requests.

## Use the test suite
***
Use [Testify's `require`](https://github.com/stretchr/testify/#require-package) package to run the test suite

```shell
go get github.com/stretchr/testify/require
```

## Handle errors
***

```go
_,err := //Go Library call 

if err != nil {
  if goErr,ok := err.(*chargebee.Error); ok {

    //Identify the type of Error 
    switch goErr.Type {
      
    case chargebee.PaymentError:
      // First check for card parameters entered by the user.
        // We recommend you to validate the input at the client side itself to catch simple mistakes.
        if goErr.Param == "card[number]" {
          // Ask your user to recheck the card number. A better way is to use 
          // Stripe's https://github.com/stripe/jquery.payment for validating it in the client side itself.  
          //}else if(goErr.Param == &lt;other card params&gt;){ 
            //Similarly check for other card parameters entered by the user.
            //....
        } else {
            // Verfication or processing failures.
            // Provide a standard message to your user to recheck his card details or provide a different card.
            // Like  'Sorry,there was a problem when processing your card, please check the details and try again'. 
        }

      case chargebee.InvalidRequestError:
        // For coupons you could decide to provide specific messages by using 
        // the 'api_error_code' attribute in the ex.
        if goErr.Param == "coupon" {
          if goErr.APIErrorCode == "resource_not_found" {
            // Inform user to recheck his coupon code.
          } else if goErr.APIErrorCode == "resource_limit_exhausted" {
            // Inform user that the coupon code has expired.
          } else if goErr.APIErrorCode == "invalid_request" {
            // Inform user that the coupon code is not applicable for his plan(/addons).
          } else {
            // Inform user to recheck his coupon code.
          }
        } else {
          // Since you would have validated all other parameters on your side itself, 
          // this could probably be a bug in your code. Provide a generic message to your users.
        }

    case chargebee.OperationFailedError:
      // Indicates that the request parameters were right but the request couldn't be completed.
        // The reasons might be "api_request_limit_exceeded" or could be due to an issue in ChargeBee side.
        // These should occur very rarely and mostly be of temporary nature. 
        // You could ask your user to retry after some time.
      default :
        // These are unhandled exceptions (Could be due to a bug in your code or very rarely in client library).
          // The errors from ChargeBee such as authentication failures will come here.
            // You could ask users contact your support.     
    }
  }
}
```

## Contribution
***
You may contribute patches to any of the **Active** versions of this library. To do so, raise a PR against the [respective branch](#library-versions).

If you find something amiss, you are welcome to create an [issue](https://github.com/chargebee/chargebee-go/issues).

## API documentation
***

The API documentation for the Go library can be found in our [API reference](https://apidocs.chargebee.com/docs/api?lang=go).

## License
***

See the [LICENSE](LICENSE).

---

