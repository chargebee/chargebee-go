package impactedcustomer

type ImpactedCustomer struct {
	ActionType string    `json:"action_type"`
	Download   *Download `json:"download"`
	Object     string    `json:"object"`
}
type Download struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	MimeType    string `json:"mime_type"`
	Object      string `json:"object"`
}
