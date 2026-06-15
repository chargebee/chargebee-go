package ledgeroperation

import (
	"github.com/chargebee/chargebee-go/v3/filter"
	ledgerOperationEnum "github.com/chargebee/chargebee-go/v3/models/ledgeroperation/enum"
)

type LedgerOperation struct {
	Id                       string                       `json:"id"`
	Type                     ledgerOperationEnum.Type     `json:"type"`
	Amount                   string                       `json:"amount"`
	StartBalance             string                       `json:"start_balance"`
	EndBalance               string                       `json:"end_balance"`
	ProvisionedStartBalance  string                       `json:"provisioned_start_balance"`
	ProvisionedEndBalance    string                       `json:"provisioned_end_balance"`
	OverdraftStartBalance    string                       `json:"overdraft_start_balance"`
	OverdraftEndBalance      string                       `json:"overdraft_end_balance"`
	ParentLedgerOperationId  string                       `json:"parent_ledger_operation_id"`
	LedgerOperationTimestamp int64                        `json:"ledger_operation_timestamp"`
	AutoReleaseTimestamp     int64                        `json:"auto_release_timestamp"`
	Metadata                 string                       `json:"metadata"`
	CreatedAt                int64                        `json:"created_at"`
	ModifiedAt               int64                        `json:"modified_at"`
	SubscriptionId           string                       `json:"subscription_id"`
	UnitId                   string                       `json:"unit_id"`
	UnitType                 ledgerOperationEnum.UnitType `json:"unit_type"`
	Object                   string                       `json:"object"`
}
type ListLedgerOperationsRequestParams struct {
	Limit          *int32                  `json:"limit,omitempty"`
	Offset         string                  `json:"offset,omitempty"`
	SubscriptionId *filter.StringFilter    `json:"subscription_id"`
	UnitId         *filter.StringFilter    `json:"unit_id,omitempty"`
	CreatedAt      *filter.TimestampFilter `json:"created_at,omitempty"`
	Type           *filter.EnumFilter      `json:"type,omitempty"`
	SortBy         *filter.SortFilter      `json:"sort_by,omitempty"`
}
type CaptureRequestParams struct {
	Id                       string                 `json:"id,omitempty"`
	SubscriptionId           string                 `json:"subscription_id"`
	UnitId                   string                 `json:"unit_id"`
	Amount                   string                 `json:"amount"`
	LedgerOperationTimestamp *int64                 `json:"ledger_operation_timestamp"`
	Metadata                 map[string]interface{} `json:"metadata,omitempty"`
}
type AuthorizeRequestParams struct {
	Id                       string                 `json:"id,omitempty"`
	SubscriptionId           string                 `json:"subscription_id"`
	UnitId                   string                 `json:"unit_id"`
	Amount                   string                 `json:"amount"`
	LedgerOperationTimestamp *int64                 `json:"ledger_operation_timestamp"`
	AutoReleaseTimestamp     *int64                 `json:"auto_release_timestamp,omitempty"`
	Metadata                 map[string]interface{} `json:"metadata,omitempty"`
}
type CaptureAuthorizationRequestParams struct {
	AuthorizationId          string                 `json:"authorization_id"`
	Id                       string                 `json:"id,omitempty"`
	Amount                   string                 `json:"amount"`
	LedgerOperationTimestamp *int64                 `json:"ledger_operation_timestamp"`
	Metadata                 map[string]interface{} `json:"metadata,omitempty"`
}
type ReleaseAuthorizationRequestParams struct {
	AuthorizationId          string                 `json:"authorization_id"`
	Id                       string                 `json:"id,omitempty"`
	LedgerOperationTimestamp *int64                 `json:"ledger_operation_timestamp"`
	Metadata                 map[string]interface{} `json:"metadata,omitempty"`
}
