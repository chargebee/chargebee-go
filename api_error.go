package chargebee

import "encoding/json"

// ErrorType is the list of allowed values for error type.
type ErrorType string

const (
	PaymentError                         ErrorType = "payment"
	InvalidRequestError                  ErrorType = "invalid_request"
	OperationFailedError                 ErrorType = "operation_failed"
	UbbBatchIngestionInvalidRequestError ErrorType = "ubb_batch_ingestion_invalid_request"
)

// Error is the Response returned when a call is unsuccessful
type Error struct {
	HTTPStatusCode int           `json:"http_status_code"`
	Msg            string        `json:"message"`
	Param          string        `json:"param"`
	APIErrorCode   string        `json:"api_error_code"`
	Type           ErrorType     `json:"type"`
	ErrorCode      string        `json:"error_code"`
	Err            error         `json:"_"`
	ErrorCauseID   string        `json:"error_cause_id,omitempty"`
	BatchId        string        `json:"batch_id"`
	FailedEvents   []interface{} `json:"failed_events"`
}

// Error Serializes the error object to JSON and return it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// PaymentErr Occurs when there is Payment related errors.
type paymentErr struct {
	cbErr *Error
}

func (e *paymentErr) Error() string {
	return e.cbErr.Error()
}

// InvalidRequestErr occurs when a request contains invalid parameters.
type invalidRequestErr struct {
	cbErr *Error
}

func (e *invalidRequestErr) Error() string {
	return e.cbErr.Error()
}

// OperationFailedErr occurs when the request parameters were right but the request could not be completed.
type operationFailedErr struct {
	cbErr *Error
}

func (e *operationFailedErr) Error() string {
	return e.cbErr.Error()
}

type ubbBatchIngestionInvalidRequestErr struct {
	cbErr *Error
}

func (e *ubbBatchIngestionInvalidRequestErr) Error() string {
	return e.cbErr.Error()
}
