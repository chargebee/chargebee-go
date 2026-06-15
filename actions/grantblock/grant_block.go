package grantblock

import (
	"fmt"
	"github.com/chargebee/chargebee-go/v3"
	"github.com/chargebee/chargebee-go/v3/models/grantblock"
)

func ListGrantBlocks(params *grantblock.ListGrantBlocksRequestParams) chargebee.ListRequest {
	return chargebee.SendList("GET", fmt.Sprintf("/grant_blocks"), params)
}
