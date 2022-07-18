package impactedsubscription

import (
	"encoding/json"
)

type ImpactedSubscription struct {
	Count           int32           `json:"count"`
	Download        *Download       `json:"download"`
	SubscriptionIds json.RawMessage `json:"subscription_ids"`
	Object          string          `json:"object"`
}
type Download struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	MimeType    string `json:"mime_type"`
	Object      string `json:"object"`
}
