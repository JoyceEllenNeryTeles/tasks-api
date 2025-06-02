package entity

type Task struct {
	Id          int64  `json:"id_taks"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Status      string `json:"status"`
}
