package models

type Label struct {
	ID    int64  `json:"id" db:"id"`
	Key   string `json:"key" db:"key"`
	Value string `json:"value" db:"value"`
}
