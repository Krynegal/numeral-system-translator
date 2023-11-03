package models

type Request struct {
	Number *string `json:"number"`
	Base   *int    `json:"base"`
	ToBase *int    `json:"to_base"`
}
