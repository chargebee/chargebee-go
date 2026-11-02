package emailattachment

type EmailAttachment struct {
	Id                 string `json:"id"`
	EmailMessageId     string `json:"email_message_id"`
	FileName           string `json:"file_name"`
	ContentType        string `json:"content_type"`
	SizeBytes          int64  `json:"size_bytes"`
	ContentDisposition string `json:"content_disposition"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	ResourceVersion    int64  `json:"resource_version"`
	Object             string `json:"object"`
}
