package models

type Word struct {
	ID    int64  `db:"id" json:"id"`
	Value string `db:"value" json:"value"`
}
