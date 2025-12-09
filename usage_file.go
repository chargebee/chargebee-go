package chargebee

type Status string

const (
	StatusQueued     Status = "queued"
	StatusImported   Status = "imported"
	StatusProcessing Status = "processing"
	StatusProcessed  Status = "processed"
	StatusFailed     Status = "failed"
)

type UsageFile struct {
	Id                    string        `json:"id"`
	Name                  string        `json:"name"`
	MimeType              string        `json:"mime_type"`
	ErrorCode             string        `json:"error_code"`
	ErrorReason           string        `json:"error_reason"`
	Status                Status        `json:"status"`
	TotalRecordsCount     int64         `json:"total_records_count"`
	ProcessedRecordsCount int64         `json:"processed_records_count"`
	FailedRecordsCount    int64         `json:"failed_records_count"`
	FileSizeInBytes       int64         `json:"file_size_in_bytes"`
	ProcessingStartedAt   int64         `json:"processing_started_at"`
	ProcessingCompletedAt int64         `json:"processing_completed_at"`
	UploadedBy            string        `json:"uploaded_by"`
	UploadedAt            int64         `json:"uploaded_at"`
	ErrorFilePath         string        `json:"error_file_path"`
	ErrorFileUrl          string        `json:"error_file_url"`
	UploadDetails         *UploadDetail `json:"upload_details"`
	Object                string        `json:"object"`
}
type UploadDetail struct {
	Url       string `json:"url"`
	ExpiresAt int64  `json:"expires_at"`
	Object    string `json:"object"`
}
type UploadUrlRequest struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}

type UploadUrlResponse struct {
	UsageFile *UsageFile `json:"usage_file,omitempty"`
}

type ProcessingStatusResponse struct {
	UsageFile *UsageFile `json:"usage_file,omitempty"`
}
