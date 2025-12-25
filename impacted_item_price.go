package chargebee

import (
	"encoding/json"
)

// just struct
type ImpactedItemPrice struct {
	Count      int32                      `json:"count"`
	Download   *ImpactedItemPriceDownload `json:"download"`
	ItemPrices json.RawMessage            `json:"item_prices"`
	Object     string                     `json:"object"`
}

// sub resources
type ImpactedItemPriceDownload struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	MimeType    string `json:"mime_type"`
	Object      string `json:"object"`
}

// operations
// input params
