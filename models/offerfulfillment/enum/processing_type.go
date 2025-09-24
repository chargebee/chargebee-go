package enum

type ProcessingType string

const (
	ProcessingTypeBillingUpdate ProcessingType = "billing_update"
	ProcessingTypeCheckout      ProcessingType = "checkout"
	ProcessingTypeUrlRedirect   ProcessingType = "url_redirect"
	ProcessingTypeWebhook       ProcessingType = "webhook"
	ProcessingTypeEmail         ProcessingType = "email"
)
