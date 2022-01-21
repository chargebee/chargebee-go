package download

type Download struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	MimeType    string `json:"mime_type"`
	Object      string `json:"object"`
}
