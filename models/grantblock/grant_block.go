package grantblock

import (
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
	grantBlockEnum "github.com/chargebee/chargebee-go/v3/models/grantblock/enum"
)

type GrantBlock struct {
	Id                 string                     `json:"id"`
	GrantedAmount      string                     `json:"granted_amount"`
	EffectiveFrom      int64                      `json:"effective_from"`
	ExpiresAt          int64                      `json:"expires_at"`
	Balance            string                     `json:"balance"`
	HoldAmount         string                     `json:"hold_amount"`
	UsedAmount         string                     `json:"used_amount"`
	ExpiredAmount      string                     `json:"expired_amount"`
	RolledOverAmount   string                     `json:"rolled_over_amount"`
	VoidedAmount       string                     `json:"voided_amount"`
	OriginGrantBlockId string                     `json:"origin_grant_block_id"`
	Status             enum.Status                `json:"status"`
	Metadata           string                     `json:"metadata"`
	GrantSource        grantBlockEnum.GrantSource `json:"grant_source"`
	CreatedAt          int64                      `json:"created_at"`
	AccountType        grantBlockEnum.AccountType `json:"account_type"`
	UnitId             string                     `json:"unit_id"`
	UnitType           grantBlockEnum.UnitType    `json:"unit_type"`
	Object             string                     `json:"object"`
}
type ListGrantBlocksRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id"`
	UnitId         *filter.StringFilter    `json:"unit_id,omitempty"`
	EffectiveFrom  *filter.TimestampFilter `json:"effective_from,omitempty"`
	ExpiresAt      *filter.TimestampFilter `json:"expires_at,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
