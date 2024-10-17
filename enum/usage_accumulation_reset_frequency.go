package enum

type UsageAccumulationResetFrequency string

const (
	UsageAccumulationResetFrequencyNever                        UsageAccumulationResetFrequency = "never"
	UsageAccumulationResetFrequencySubscriptionBillingFrequency UsageAccumulationResetFrequency = "subscription_billing_frequency"
)
