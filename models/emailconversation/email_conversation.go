package emailconversation

type EmailConversation struct {
	Id                 string `json:"id"`
	CustomerId         string `json:"customer_id"`
	MailboxId          string `json:"mailbox_id"`
	Subject            string `json:"subject"`
	LastMessagePreview string `json:"last_message_preview"`
	MessageCount       int32  `json:"message_count"`
	Status             string `json:"status"`
	LastMessageAt      int64  `json:"last_message_at"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	ResourceVersion    int64  `json:"resource_version"`
	Object             string `json:"object"`
}
