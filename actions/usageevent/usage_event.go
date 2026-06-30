package usageevent

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/usageevent"
)

func Create(params *usageevent.CreateRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/usage_events"), params).WithTelemetryResource("usageEvent").WithTelemetryOperation("create").SetSubDomain("ingest")
}
func BatchIngest(params *usageevent.BatchIngestRequestParams) chargebee.Request {
	return chargebee.SendJsonRequest("POST", fmt.Sprintf("/batch/usage_events"), params).WithTelemetryResource("usageEvent").WithTelemetryOperation("batchIngest").SetSubDomain("ingest")
}
