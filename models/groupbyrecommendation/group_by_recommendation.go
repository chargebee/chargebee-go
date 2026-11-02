package groupbyrecommendation

type GroupByRecommendation struct {
	AttributeName string `json:"attribute_name"`
	Rank          int32  `json:"rank"`
	Reason        string `json:"reason"`
	Object        string `json:"object"`
}
