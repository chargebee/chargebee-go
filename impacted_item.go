package chargebee

import (
	"encoding/json"
)

// just struct
type ImpactedItem struct {
	Count    int32           `json:"count"`
	Download *Download       `json:"download"`
	Items    json.RawMessage `json:"items"`
	Object   string          `json:"object"`
}

// sub resources
type ImpactedItemDownload struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	MimeType    string `json:"mime_type"`
	Object      string `json:"object"`
}

// operations
// input params
