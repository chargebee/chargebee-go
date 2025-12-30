package enum

type RetryEngine string

const (
	RetryEngineChargebee   RetryEngine = "chargebee"
	RetryEngineFlexpay     RetryEngine = "flexpay"
	RetryEngineSuccessplus RetryEngine = "successplus"
)
