package enum

type Type string

const (
	TypeRecommended Type = "recommended"
	TypeMandatory   Type = "mandatory"
	TypeOptional    Type = "optional"
)
