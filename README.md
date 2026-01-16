# Chargebee Go Client Library

> [!NOTE]
> [![Join Discord](https://img.shields.io/badge/Discord-Early%20Access-blue?logo=discord&logoColor=white)](https://discord.gg/S3SXDzXHAg)
>
> We are trialing a Discord server for developers building with Chargebee. Limited spots are open on a first-come basis. Join [here](https://discord.gg/S3SXDzXHAg) if interested.

> [!CAUTION]
> This new major version (v4) of the Chargebee GO SDK contains BREAKING CHANGES and is currently in beta. It is a complete rewrite and addresses feedback raised by the community. If interested, please try it out and share your feedback with us on the Discord channel.
>
> A [migration guide](https://github.com/chargebee/chargebee-go/wiki/Go-SDK-v4-migration-guide) is available on the wiki for users upgrading from v3.


This is the official Go library for integrating with Chargebee
- 📘 For a complete reference of available APIs, check out our [API Documentation](https://apidocs.chargebee.com/docs/api/?lang=go).  
- 🧪 To explore and test API capabilities interactively, head over to our [API Explorer](https://api-explorer.chargebee.com).

**Go version**: v4 requires Go 1.21 or higher. v3 requires Go v1.3 or higher.

## Library versions
***

The versioning scheme of this library is inspired by [SemVer](https://semver.org/) and the format is `v{MAJOR}.{MINOR}.{PATCH}`. For example, `v3.0.0` and `v2.5.1` are valid library versions.

The following table provides some details for each major version:

| Library major version | Status   | Compatible API versions                                                                             | **Branch**        |
|----------------------------|----------|-----------------------------------------------------------------------------------------------------|---------------|
| v4                         | Active   | [v2](https://apidocs.chargebee.com/docs/api/v2?lang=go) and [v1](https://apidocs.chargebee.com/docs/api/v1?lang=go) | `v4`      |
| v3                         | Active   | [v2](https://apidocs.chargebee.com/docs/api/v2?lang=go) and [v1](https://apidocs.chargebee.com/docs/api/v1?lang=go) | `master`      |


A couple of terms used in the above table are explained below:
- **Status**: The current development status for the library version. An **Active** major version is currently being maintained and continues to get backward-compatible changes.
- **Branch**: The branch in this repository containing the source code for the latest release of the library version. Every version of the library has been [tagged](https://github.com/chargebee/chargebee-go/tags). You can check out the source code for any version using its tag.


**Note:** See the [changelog](CHANGELOG.md) for a history of changes.

## Install the library
***

Install the latest version of the library with the following commands:

``` shell
go get github.com/chargebee/chargebee-go/v4
```

## Use the library
***

Some examples for using the library are listed below.

### Create a subscription with items

```go
package main

import (
	"fmt"
	"os"

	"github.com/chargebee/chargebee-go/v4"
)

func main() {
	config := &chargebee.ClientConfig{
		SiteName: os.Getenv("CHARGEBEE_SITE"),
		ApiKey:   os.Getenv("CHARGEBEE_API_KEY"),
	}
	client := chargebee.NewClient(config)
	
	req := &chargebee.SubscriptionCreateWithItemsRequest{
		Id: "__test__8asz8Ru9WhHOJO",
		SubscriptionItems: []*chargebee.SubscriptionCreateWithItemsSubscriptionItem{
			{
				ItemPriceId: "day-pass-USD",
				UnitPrice:   chargebee.Int64(100),
			},
			{
				ItemPriceId:   "basic-USD",
				BillingCycles: chargebee.Int32(2),
				Quantity:      chargebee.Int32(1),
			},
		},
	}
	if response, err := client.Subscription.CreateWithItems("__test__8asz8Ru9WhHOJO", req); err == nil {
		fmt.Println(string(response.Subscription.Id))
	}
}
```

### Create a subscription with metadata, custom headers and fields

```go
package main

import (
	"fmt"
	"os"

	"github.com/chargebee/chargebee-go/v4"
)

func main() {
	config := &chargebee.ClientConfig{
		SiteName: os.Getenv("CHARGEBEE_SITE"),
		ApiKey:   os.Getenv("CHARGEBEE_API_KEY"),
	}
	client := chargebee.NewClient(config)
	
	req := &chargebee.SubscriptionCreateWithItemsRequest{
		Id: "__test__8asz8Ru9WhHOJO",
		SubscriptionItems: []*chargebee.SubscriptionCreateWithItemsSubscriptionItem{
			{
				ItemPriceId: "day-pass-USD",
				UnitPrice:   chargebee.Int64(100),
			},
			{
				ItemPriceId:   "basic-USD",
				BillingCycles: chargebee.Int32(2),
				Quantity:      chargebee.Int32(1),
			},
		},
    MetaData: map[string]any{
      "features": map[string]any{
        "usage-limit":        "5GB",
        "speed-within-quota": "2MBbps",
        "post-usage-quota":   "512kbps",
      },
    },
	}

  req.AddHeader("chargebee-request-origin-ip", "192.168.1.2")
  req.AddCustomField("cf_gender", "Female")
	
  if response, err := client.Subscription.CreateWithItems("__test__8asz8Ru9WhHOJO", req); err == nil {
		fmt.Println(string(response.Subscription.Id))
	}
}
```

### Retrieve a filtered list of subscriptions

```go
package main

import (
	"fmt"
	"os"

	"github.com/chargebee/chargebee-go/v4"
)

func main() {
	config := &chargebee.ClientConfig{
		SiteName: os.Getenv("CHARGEBEE_SITE"),
		ApiKey:   os.Getenv("CHARGEBEE_API_KEY"),
	}
	client := chargebee.NewClient(config)

	res, err := client.Subscription.List(&chargebee.SubscriptionListRequest{
		Limit: chargebee.Int32(5),
		Id: &chargebee.StringFilter{
			In: []string{"cbdemo_john-sub", "cbdemo_ricky-sub"},
		},
		PlanId: &chargebee.StringFilter{
			IsNot: "basic",
		},
		Status: &chargebee.EnumFilter{
			Is: chargebee.SubscriptionStatusActive,
		},
		SortBy: &chargebee.SortFilter{
			Asc: "created_at",
		},
	})

	if err != nil {
		panic(err)
	} else {
		for i := range res.List {
			subscription := res.List[i].Subscription
			customer := res.List[i].Customer
			card := res.List[i].Card
			
			fmt.Printf("Subscription: %s, Customer: %s\n", subscription.Id, customer.Id)
			if card.Status != "" {
				fmt.Printf("Card: %s\n", card.MaskedNumber)
			}
		}
	}
}
```

### Create an idempotent request

[Idempotency keys](https://apidocs.chargebee.com/docs/api/idempotency?prod_cat_ver=2) are passed along with request headers to allow a safe retry of POST requests. 

```go
package main

import (
	"fmt"
	"os"

	"github.com/chargebee/chargebee-go/v4"
	"github.com/google/uuid"
)

func main() {
	config := &chargebee.ClientConfig{
		SiteName: os.Getenv("CHARGEBEE_SITE"),
		ApiKey:   os.Getenv("CHARGEBEE_API_KEY"),
	}
	client := chargebee.NewClient(config)

	req := &chargebee.CustomerCreateRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
	}
	
	// Set Idempotency Key
	req.SetIdempotencyKey(uuid.NewString()) // Replace with a unique string

	if res, err := client.Customer.Create(req); err == nil {
		fmt.Println(res.Customer)
		fmt.Println(res.IsIdempotencyReplayed())
	}

	// Try the same request again with the same idempotency key
	// This throw a HTTP 422 error since the same key is being reused
	if res, err := client.Customer.Create(req); err == nil {
		fmt.Println(res.IsIdempotencyReplayed())
	} else {
		fmt.Println(err)
	}
}
```

`IsIdempotencyReplayed()` method can be accessed to differentiate between original and replayed requests.

### Handle webhooks

Use the `webhook` package to parse and route webhook payloads from Chargebee.

*High-level*: route events with callbacks using `WebhookHandler`:

```go
package main

import (
	"log"
	"net/http"

	"github.com/chargebee/chargebee-go/v4/webhook"
)

func main() {
	handler := &webhook.WebhookHandler{
		// Optional: protect endpoint (e.g., Basic Auth)
		RequestValidator: webhook.BasicAuthValidator(func(user, pass string) bool {
			return user == "admin" && pass == "secret"
		}),
		OnError: webhook.BasicAuthErrorHandler, // Optional: standard auth error responses

		// Register only the events you care about
		OnSubscriptionCreated: func(e webhook.SubscriptionCreatedEvent) error {
			log.Printf("Subscription created event %s", e.Id)
			return nil
		},
		OnPaymentSucceeded: func(e webhook.PaymentSucceededEvent) error {
			log.Printf("Payment succeeded for customer: %v", e.Content.Customer)
			return nil
		},
	}

	http.Handle("/chargebee/webhooks", handler.HTTPHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

*Low-level*: parse just the event type and unmarshal yourself:

```go
import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chargebee/chargebee-go/v4"
	"github.com/chargebee/chargebee-go/v4/webhook"
)

func cbWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	evtType, err := webhook.ParseEventType(body) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch evtType {
	case chargebee.EventTypeSubscriptionCreated:
		var e webhook.SubscriptionCreatedEvent
		if err := json.Unmarshal(body, &e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// handle e
	default:
		// ignore or log
	}
	w.WriteHeader(http.StatusOK)
}
```

#### Unhandled events

By default, if an incoming webhook’s event type is unknown or you have not registered a corresponding handler on `WebhookHandler`, the SDK treats it as an error. When using `HTTPHandler()`, this results in a 500 response unless you provide a custom `OnError` handler.

If you prefer to acknowledge unknown/unregistered events (return 200) and just log them, set `OnUnhandledEvent` to a function that returns `nil`:

```go
import (
	"log"
	"github.com/chargebee/chargebee-go/v4"
	"github.com/chargebee/chargebee-go/v4/webhook"
)

// Create a webhook handler with the OnUnhandledEvent callback
handler := &webhook.WebhookHandler{
	OnUnhandledEvent: func(t chargebee.EventType, body []byte) error {
		log.Printf("Ignoring unhandled event: %s", t)
		return nil // swallow as OK
	},
}
```

## Handling API errors
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

### Retrying API requests

Chargebee's SDK includes built-in retry logic to handle temporary network issues and server-side errors. This feature is **disabled by default** but can be **enabled when needed**.

#### Key features include:

- **Automatic retries for specific HTTP status codes**: Retries are automatically triggered for status codes `500`, `502`, `503`, and `504`.
- **Exponential backoff**: Retry delays increase exponentially to prevent overwhelming the server.
- **Rate limit management**: If a `429 Too Many Requests` response is received with a `Retry-After` header, the SDK waits for the specified duration before retrying.
  > *Note: Exponential backoff and max retries do not apply in this case.*
- **Customizable retry behavior**: Retry logic can be configured using the `retryConfig` parameter in the environment configuration.

#### Example: Customizing Retry Logic

You can enable and configure the retry logic by passing a `retryConfig` object when initializing the Chargebee environment:

```go
package main

import (
	"os"

	"github.com/chargebee/chargebee-go/v4"
)

func main() {
	config := &chargebee.ClientConfig{
		SiteName: os.Getenv("CHARGEBEE_SITE"),
		ApiKey:   os.Getenv("CHARGEBEE_API_KEY"),
		RetryConfig: &chargebee.RetryConfig{
			Enabled:    true,
			MaxRetries: 3,
			DelayMs:    500,
			RetryOn:    []int{500, 503}, // Retry on specific status codes
		},
	}
	
	client := chargebee.NewClient(config)
	
	// ... your Chargebee API operations below ...
}
```

### Using a custom HTTP Client

The default http client can be overridden if required with a client that satisfies the `http.Client` interface. To do this, simply override the `HTTPClient` field in `ClientConfig`:

```go
package main

import (
	"crypto/tls"
	"os"
	"net/http"
	"time"

	"github.com/chargebee/chargebee-go/v4"
)

func main() {
	config := chargebee.NewClientConfig(os.Getenv("CHARGEBEE_SITE"), os.Getenv("CHARGEBEE_API_KEY"))
	config.HTTPClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 100,
			TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12},
		},
		Timeout: 10 * time.Second,
	}
	client := chargebee.NewClient(config)
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
