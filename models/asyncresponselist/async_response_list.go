package asyncresponselist

import (
	"github.com/chargebee/chargebee-go/v3/models/asyncresponse"
)

type AsyncResponseList struct {
	List   []*asyncresponse.AsyncResponse `json:"list"`
	Object string                         `json:"object"`
}
