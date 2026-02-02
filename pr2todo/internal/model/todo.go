package model

type Todo struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
