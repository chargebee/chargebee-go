package enum

type ChargeEvent string

const (
	ChargeEventImmediate            ChargeEvent = "immediate"
	ChargeEventSubscriptionCreation ChargeEvent = "subscription_creation"
	ChargeEventTrialStart           ChargeEvent = "trial_start"
	ChargeEventSubscriptionChange   ChargeEvent = "subscription_change"
	ChargeEventSubscriptionRenewal  ChargeEvent = "subscription_renewal"
	ChargeEventSubscriptionCancel   ChargeEvent = "subscription_cancel"
)
