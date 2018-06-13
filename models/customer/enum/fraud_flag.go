package enum

type FraudFlag string

const (
	FraudFlagSafe       FraudFlag = "safe"
	FraudFlagSuspicious FraudFlag = "suspicious"
	FraudFlagFraudulent FraudFlag = "fraudulent"
)
