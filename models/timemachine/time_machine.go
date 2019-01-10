package timemachine

import (
	timeMachineEnum "github.com/chargebee/chargebee-go/models/timemachine/enum"
)

type TimeMachine struct {
	Name             string                           `json:"name"`
	TimeTravelStatus timeMachineEnum.TimeTravelStatus `json:"time_travel_status"`
	GenesisTime      int64                            `json:"genesis_time"`
	DestinationTime  int64                            `json:"destination_time"`
	FailureCode      string                           `json:"failure_code"`
	FailureReason    string                           `json:"failure_reason"`
	ErrorJson        string                           `json:"error_json"`
	Object           string                           `json:"object"`
}
type StartAfreshRequestParams struct {
	GenesisTime *int64 `json:"genesis_time,omitempty"`
}
type TravelForwardRequestParams struct {
	DestinationTime *int64 `json:"destination_time,omitempty"`
}
