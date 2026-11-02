package citation

import (
	"encoding/json"
	citationEnum "github.com/chargebee/chargebee-go/v3/models/citation/enum"
)

type Citation struct {
	Id              string                       `json:"id"`
	ParamName       string                       `json:"param_name"`
	AuthorName      string                       `json:"author_name"`
	ConfidenceLevel citationEnum.ConfidenceLevel `json:"confidence_level"`
	CreatedAt       int64                        `json:"created_at"`
	Content         json.RawMessage              `json:"content"`
	Object          string                       `json:"object"`
}
