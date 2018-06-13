package download

type Download struct {
	DownloadUrl string `json:"download_url"`
	ValidTill   int64  `json:"valid_till"`
	Object      string `json:"object"`
}
