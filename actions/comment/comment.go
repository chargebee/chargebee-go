package comment

import (
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/comment"
)

func Create(params *comment.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/comments"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/comments/%v", id), nil)
}
func List(params *comment.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/comments"), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/comments/%v/delete", id), nil)
}
