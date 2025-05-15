package usagefile

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/usagefile"
	"net/url"
)

func Upload(params *usagefile.UploadRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/usage_files/upload"), params).SetSubDomain("file-ingest")
}
func Status(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/usage_files/%v/status", url.PathEscape(id)), nil).SetSubDomain("file-ingest")
}
