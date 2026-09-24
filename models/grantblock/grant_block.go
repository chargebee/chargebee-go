package grantblock

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/v3/enum"
	"github.com/chargebee/chargebee-go/v3/filter"
	grantBlockEnum "github.com/chargebee/chargebee-go/v3/models/grantblock/enum"
)

type GrantBlock struct {
	Id             string                     `json:"id"`
	SubscriptionId string                     `json:"subscription_id"`
	UnitId         string                     `json:"unit_id"`
	UnitType       grantBlockEnum.UnitType    `json:"unit_type"`
	AccountType    grantBlockEnum.AccountType `json:"account_type"`
	//Deprecated: this field is deprecated
	GrantedAmount string `json:"granted_amount"`
	EffectiveFrom int64  `json:"effective_from"`
	ExpiresAt     int64  `json:"expires_at"`
	//Deprecated: this field is deprecated
	Balance string `json:"balance"`
	//Deprecated: this field is deprecated
	HoldAmount string `json:"hold_amount"`
	//Deprecated: this field is deprecated
	UsedAmount string `json:"used_amount"`
	//Deprecated: this field is deprecated
	ExpiredAmount string `json:"expired_amount"`
	//Deprecated: this field is deprecated
	RolledOverAmount string `json:"rolled_over_amount"`
	//Deprecated: this field is deprecated
	VoidedAmount            string                     `json:"voided_amount"`
	OriginGrantBlockId      string                     `json:"origin_grant_block_id"`
	Status                  enum.Status                `json:"status"`
	GrantSource             grantBlockEnum.GrantSource `json:"grant_source"`
	CreatedAt               int64                      `json:"created_at"`
	ModifiedAt              int64                      `json:"modified_at"`
	ResourceVersion         int64                      `json:"resource_version"`
	ProvisionedBlockBalance *ProvisionedBlockBalance   `json:"provisioned_block_balance"`
	OverdraftBlockBalance   *OverdraftBlockBalance     `json:"overdraft_block_balance"`
	Metadata                json.RawMessage            `json:"metadata"`
	Object                  string                     `json:"object"`
}
type ProvisionedBlockBalance struct {
	GrantedAmount    string `json:"granted_amount"`
	TotalBalance     string `json:"total_balance"`
	UsableBalance    string `json:"usable_balance"`
	HoldAmount       string `json:"hold_amount"`
	UsedAmount       string `json:"used_amount"`
	ExpiredAmount    string `json:"expired_amount"`
	RolledOverAmount string `json:"rolled_over_amount"`
	VoidedAmount     string `json:"voided_amount"`
	Object           string `json:"object"`
}
type OverdraftBlockBalance struct {
	IsUnlimited   bool   `json:"is_unlimited"`
	Limit         string `json:"limit"`
	TotalBalance  string `json:"total_balance"`
	UsableBalance string `json:"usable_balance"`
	UsedAmount    string `json:"used_amount"`
	Object        string `json:"object"`
}
type ListGrantBlocksRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id"`
	UnitId         *filter.StringFilter    `json:"unit_id,omitempty"`
	AccountType    *filter.EnumFilter      `json:"account_type,omitempty"`
	EffectiveFrom  *filter.TimestampFilter `json:"effective_from,omitempty"`
	ExpiresAt      *filter.TimestampFilter `json:"expires_at,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
