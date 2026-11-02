package privatenote

import (
	privateNoteEnum "github.com/chargebee/chargebee-go/v3/models/privatenote/enum"
)

type PrivateNote struct {
	Id                  string                     `json:"id"`
	EmailConversationId string                     `json:"email_conversation_id"`
	AuthorType          privateNoteEnum.AuthorType `json:"author_type"`
	AuthorId            string                     `json:"author_id"`
	AuthorName          string                     `json:"author_name"`
	ContentHtml         string                     `json:"content_html"`
	ContentText         string                     `json:"content_text"`
	MetaData            string                     `json:"meta_data"`
	CreatedAt           int64                      `json:"created_at"`
	ResourceVersion     int64                      `json:"resource_version"`
	UpdatedAt           int64                      `json:"updated_at"`
	Object              string                     `json:"object"`
}
