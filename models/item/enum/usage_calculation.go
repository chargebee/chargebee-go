package enum

type UsageCalculation string

const (
	UsageCalculationSumOfUsages UsageCalculation = "sum_of_usages"
	UsageCalculationLastUsage   UsageCalculation = "last_usage"
)
