package usagefile

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/usagefile"
	"net/url"
)

func UploadUrl(params *usagefile.UploadUrlRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/usage_files/upload_url"), params).SetSubDomain("file-ingest")
}
func ProcessingStatus(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/usage_files/%v/processing_status", url.PathEscape(id)), nil).SetSubDomain("file-ingest")
}
