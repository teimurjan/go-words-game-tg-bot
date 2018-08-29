package models

type Game struct {
	ID            int64 `db:"id" json:"id"`
	IsUserTurn    bool  `db:"is_user_turn" json:"isUserTurn"`
	LastBotWordID int64 `db:"last_bot_word_id" json:"lastBotWordId"`
	UserID        int64 `db:"user_id" json:"userId"`
}
