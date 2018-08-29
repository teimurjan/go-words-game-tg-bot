package models

type User struct {
	ID     int64 `db:"id" json:"id"`
	ChatID int64 `db:"chat_id" json:"chatId"`
}
