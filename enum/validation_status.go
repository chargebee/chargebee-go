package enum

type ValidationStatus string

const (
	ValidationStatusNotValidated   ValidationStatus = "not_validated"
	ValidationStatusValid          ValidationStatus = "valid"
	ValidationStatusPartiallyValid ValidationStatus = "partially_valid"
	ValidationStatusInvalid        ValidationStatus = "invalid"
)
