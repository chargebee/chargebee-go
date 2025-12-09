package chargebee

type TimeTravelStatus string

const (
	TimeTravelStatusNotEnabled TimeTravelStatus = "not_enabled"
	TimeTravelStatusInProgress TimeTravelStatus = "in_progress"
	TimeTravelStatusSucceeded  TimeTravelStatus = "succeeded"
	TimeTravelStatusFailed     TimeTravelStatus = "failed"
)

type TimeMachine struct {
	Name             string           `json:"name"`
	TimeTravelStatus TimeTravelStatus `json:"time_travel_status"`
	GenesisTime      int64            `json:"genesis_time"`
	DestinationTime  int64            `json:"destination_time"`
	FailureCode      string           `json:"failure_code"`
	FailureReason    string           `json:"failure_reason"`
	ErrorJson        string           `json:"error_json"`
	Object           string           `json:"object"`
}
type StartAfreshRequest struct {
	GenesisTime *int64 `json:"genesis_time,omitempty"`
}
type TravelForwardRequest struct {
	DestinationTime *int64 `json:"destination_time,omitempty"`
}

type RetrieveResponse struct {
	TimeMachine *TimeMachine `json:"time_machine,omitempty"`
}

type StartAfreshResponse struct {
	TimeMachine *TimeMachine `json:"time_machine,omitempty"`
}

type TravelForwardResponse struct {
	TimeMachine *TimeMachine `json:"time_machine,omitempty"`
}
