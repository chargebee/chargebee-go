package enum

type Source string

const (
	SourceAdminConsole    Source = "admin_console"
	SourceApi             Source = "api"
	SourceScheduledJob    Source = "scheduled_job"
	SourceHostedPage      Source = "hosted_page"
	SourcePortal          Source = "portal"
	SourceSystem          Source = "system"
	SourceNone            Source = "none"
	SourceJsApi           Source = "js_api"
	SourceMigration       Source = "migration"
	SourceBulkOperation   Source = "bulk_operation"
	SourceExternalService Source = "external_service"
)
