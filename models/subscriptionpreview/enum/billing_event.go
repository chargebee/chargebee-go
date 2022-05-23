package enum

type BillingEvent string

const (
	BillingEventSubscriptionCreated   BillingEvent = "subscription_created"
	BillingEventSubscriptionChanged   BillingEvent = "subscription_changed"
	BillingEventSubscriptionRenewed   BillingEvent = "subscription_renewed"
	BillingEventSubscriptionCancelled BillingEvent = "subscription_cancelled"
	BillingEventSubscriptionResumed   BillingEvent = "subscription_resumed"
	BillingEventSubscriptionPaused    BillingEvent = "subscription_paused"
	BillingEventSubscriptionStarted   BillingEvent = "subscription_started"
	BillingEventSubscriptionActivated BillingEvent = "subscription_activated"
)
