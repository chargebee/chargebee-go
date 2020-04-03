package enum

type EventBasedAddonOnEvent string

const (
	EventBasedAddonOnEventSubscriptionCreation   EventBasedAddonOnEvent = "subscription_creation"
	EventBasedAddonOnEventSubscriptionTrialStart EventBasedAddonOnEvent = "subscription_trial_start"
	EventBasedAddonOnEventPlanActivation         EventBasedAddonOnEvent = "plan_activation"
	EventBasedAddonOnEventSubscriptionActivation EventBasedAddonOnEvent = "subscription_activation"
	EventBasedAddonOnEventContractTermination    EventBasedAddonOnEvent = "contract_termination"
)
