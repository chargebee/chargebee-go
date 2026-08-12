package enum

type PaymentIntentMetadataSource string

const (
	PaymentIntentMetadataSourcePaymentMethodHelper PaymentIntentMetadataSource = "payment_method_helper"
	PaymentIntentMetadataSourceCardComponents      PaymentIntentMetadataSource = "card_components"
	PaymentIntentMetadataSourceCheckout            PaymentIntentMetadataSource = "checkout"
	PaymentIntentMetadataSourceCollectNow          PaymentIntentMetadataSource = "collect_now"
	PaymentIntentMetadataSourcePortal              PaymentIntentMetadataSource = "portal"
	PaymentIntentMetadataSourcePaymentComponents   PaymentIntentMetadataSource = "payment_components"
)
