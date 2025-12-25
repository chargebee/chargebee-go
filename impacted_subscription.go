package chargebee

import (
	"encoding/json"
)

// just struct
type ImpactedSubscription struct {
	Count           int32                         `json:"count"`
	Download        *ImpactedSubscriptionDownload `json:"download"`
	SubscriptionIds json.RawMessage               `json:"subscription_ids"`
	Object          string                        `json:"object"`
}

// sub resources
type ImpactedSubscriptionDownload struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	MimeType    string `json:"mime_type"`
	Object      string `json:"object"`
}

// operations
// input params
