package enum

type ChargeOnEvent string

const (
	ChargeOnEventSubscriptionCreation   ChargeOnEvent = "subscription_creation"
	ChargeOnEventSubscriptionTrialStart ChargeOnEvent = "subscription_trial_start"
	ChargeOnEventPlanActivation         ChargeOnEvent = "plan_activation"
	ChargeOnEventSubscriptionActivation ChargeOnEvent = "subscription_activation"
	ChargeOnEventContractTermination    ChargeOnEvent = "contract_termination"
	ChargeOnEventOnDemand               ChargeOnEvent = "on_demand"
)
