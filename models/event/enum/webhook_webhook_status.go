package enum

type WebhookWebhookStatus string

const (
	WebhookWebhookStatusNotConfigured WebhookWebhookStatus = "not_configured"
	WebhookWebhookStatusScheduled     WebhookWebhookStatus = "scheduled"
	WebhookWebhookStatusSucceeded     WebhookWebhookStatus = "succeeded"
	WebhookWebhookStatusReScheduled   WebhookWebhookStatus = "re_scheduled"
	WebhookWebhookStatusFailed        WebhookWebhookStatus = "failed"
	WebhookWebhookStatusSkipped       WebhookWebhookStatus = "skipped"
	WebhookWebhookStatusNotApplicable WebhookWebhookStatus = "not_applicable"
)
