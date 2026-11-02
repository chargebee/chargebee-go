package collectionspaymentpromise

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/collectionspaymentpromise"
	"net/url"
)

func Retrieve(id string) chargebee.Request {
	return chargebee.Send("GET", fmt.Sprintf("/collections_payment_promises/%v", url.PathEscape(id)), nil)
}
func ListCollectionsPaymentPromises(params *collectionspaymentpromise.ListCollectionsPaymentPromisesRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/collections_payment_promises"), params)
}
func CreateCollectionsPaymentPromise(params *collectionspaymentpromise.CreateCollectionsPaymentPromiseRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/collections_payment_promises"), params)
}
func UpdateCollectionsPaymentPromise(id string, params *collectionspaymentpromise.UpdateCollectionsPaymentPromiseRequestParams) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/collections_payment_promises/%v", url.PathEscape(id)), params).SetIdempotency(true)
}
func Cancel(id string) chargebee.Request {
	return chargebee.Send("POST", fmt.Sprintf("/collections_payment_promises/%v/cancel", url.PathEscape(id)), nil)
}
