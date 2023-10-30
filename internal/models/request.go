package models

type Request struct {
	Value string `json:"value"`
	Base  int    `json:"base"`
}

func NewRequest(value string, base int) *Request {
	return &Request{
		Value: value,
		Base:  base,
	}
}
