package appliedrule

import (
	"encoding/json"
)

type AppliedRule struct {
	Id               string          `json:"id"`
	Version          int32           `json:"version"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	EvaluationResult bool            `json:"evaluation_result"`
	ErrorMessage     string          `json:"error_message"`
	Actions          json.RawMessage `json:"actions"`
	Object           string          `json:"object"`
}
