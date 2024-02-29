package gatewayerrordetail

type GatewayErrorDetail struct {
	RequestId             string `json:"request_id"`
	ErrorCategory         string `json:"error_category"`
	ErrorCode             string `json:"error_code"`
	ErrorMessage          string `json:"error_message"`
	DeclineCode           string `json:"decline_code"`
	DeclineMessage        string `json:"decline_message"`
	NetworkErrorCode      string `json:"network_error_code"`
	NetworkErrorMessage   string `json:"network_error_message"`
	ErrorField            string `json:"error_field"`
	RecommendationCode    string `json:"recommendation_code"`
	RecommendationMessage string `json:"recommendation_message"`
	ProcessorErrorCode    string `json:"processor_error_code"`
	ProcessorErrorMessage string `json:"processor_error_message"`
	Object                string `json:"object"`
}
