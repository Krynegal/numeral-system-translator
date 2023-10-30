package models

type Response struct {
	Result string `json:"result"`
}

func NewResponse(value string) *Response {
	return &Response{
		Result: value,
	}
}
