package enum

type OptionProcessingType string

const (
	OptionProcessingTypeBillingUpdate OptionProcessingType = "billing_update"
	OptionProcessingTypeCheckout      OptionProcessingType = "checkout"
	OptionProcessingTypeUrlRedirect   OptionProcessingType = "url_redirect"
	OptionProcessingTypeWebhook       OptionProcessingType = "webhook"
	OptionProcessingTypeEmail         OptionProcessingType = "email"
)
