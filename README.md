# Chargebee Go Client Library - API V2
This is the beta version of Go Library for integrating with Chargebee. Sign up for a Chargebee account [here](https://www.chargebee.com).

## Installation

Install the latest version of the library with the following commands:

```sh
	go get github.com/chargebee/chargebee-go
```

## Go Requirement 

Use <b>go1.3 or newer</b>. 

## Usage

### To create a customer & subscription
```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go"
  subscriptionAction "github.com/chargebee/chargebee-go/actions/subscription"
  "github.com/chargebee/chargebee-go/models/subscription"
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

### To create a subscription with addons, metadata, coupons :

```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go"
  subscriptionAction "github.com/chargebee/chargebee-go/actions/subscription"
  "github.com/chargebee/chargebee-go/models/subscription"
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

### Use of Filters In List API
To retrieve list of subscriptions :

```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go"
  subscriptionAction "github.com/chargebee/chargebee-go/actions/subscription"
  "github.com/chargebee/chargebee-go/filter"
  "github.com/chargebee/chargebee-go/models/subscription"
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

### To create subscription with custom headers and custom fields:

```go
import (
  "fmt"
  "github.com/chargebee/chargebee-go"
  subscriptionAction "github.com/chargebee/chargebee-go/actions/subscription"
  "github.com/chargebee/chargebee-go/models/subscription"
)

func main() {
  chargebee.Configure("{site_api_key}", "{site}")
  res, err := subscriptionAction.Create(&subscription.CreateRequestParams{
    PlanId: "cbdemo_grow",
  }).Headers("chargebee-request-origin-ip", "192.168.1.2").AddParams("customer[cf_gender]","Female").Request() // Customer level custom field. 
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

## Test suite needs testify's require package to run :

```sh
  go get github.com/stretchr/testify/require
```

## Error Handling

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

## License

See the LICENSE file.

