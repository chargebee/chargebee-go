package enum

type WebhookStatus string

const (
	WebhookStatusNotConfigured WebhookStatus = "not_configured"
	WebhookStatusScheduled     WebhookStatus = "scheduled"
	WebhookStatusSucceeded     WebhookStatus = "succeeded"
	WebhookStatusReScheduled   WebhookStatus = "re_scheduled"
	WebhookStatusFailed        WebhookStatus = "failed"
	WebhookStatusSkipped       WebhookStatus = "skipped"
	WebhookStatusNotApplicable WebhookStatus = "not_applicable"
)
