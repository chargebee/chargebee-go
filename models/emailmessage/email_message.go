package emailmessage

import (
	emailMessageEnum "github.com/chargebee/chargebee-go/v3/models/emailmessage/enum"
)

type EmailMessage struct {
	Id                  string                      `json:"id"`
	EmailConversationId string                      `json:"email_conversation_id"`
	MessageType         string                      `json:"message_type"`
	AuthorType          emailMessageEnum.AuthorType `json:"author_type"`
	SendStatus          emailMessageEnum.SendStatus `json:"send_status"`
	ContentHtml         string                      `json:"content_html"`
	ContentText         string                      `json:"content_text"`
	CreatedAt           int64                       `json:"created_at"`
	UpdatedAt           int64                       `json:"updated_at"`
	ResourceVersion     int64                       `json:"resource_version"`
	MailboxId           string                      `json:"mailbox_id"`
	AuthorId            string                      `json:"author_id"`
	AuthorName          string                      `json:"author_name"`
	FromEmail           string                      `json:"from_email"`
	ToEmails            []string                    `json:"to_emails"`
	CcEmails            []string                    `json:"cc_emails"`
	Attachments         []*Attachment               `json:"attachments"`
	Object              string                      `json:"object"`
}
type Attachment struct {
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
