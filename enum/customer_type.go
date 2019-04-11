package enum

type CustomerType string

const (
	CustomerTypeResidential   CustomerType = "residential"
	CustomerTypeBusiness      CustomerType = "business"
	CustomerTypeSeniorCitizen CustomerType = "senior_citizen"
	CustomerTypeIndustrial    CustomerType = "industrial"
)
