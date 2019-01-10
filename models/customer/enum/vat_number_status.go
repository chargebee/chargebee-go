package enum

type VatNumberStatus string

const (
	VatNumberStatusValid        VatNumberStatus = "valid"
	VatNumberStatusInvalid      VatNumberStatus = "invalid"
	VatNumberStatusNotValidated VatNumberStatus = "not_validated"
	VatNumberStatusUndetermined VatNumberStatus = "undetermined"
)
