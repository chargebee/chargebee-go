package enum

type OnEvent string

const (
	OnEventSubscriptionCreation   OnEvent = "subscription_creation"
	OnEventSubscriptionTrialStart OnEvent = "subscription_trial_start"
	OnEventPlanActivation         OnEvent = "plan_activation"
	OnEventSubscriptionActivation OnEvent = "subscription_activation"
	OnEventContractTermination    OnEvent = "contract_termination"
)
