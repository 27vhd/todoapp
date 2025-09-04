package models

type Todo struct {
	ID   uint32 `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
