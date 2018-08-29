package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/teimurjan/go-words-game-tg-bot/texts"
)

func NewStartGameHandler(bot *tgbotapi.BotAPI) func(m *tgbotapi.Message) {
	return func(m *tgbotapi.Message) {
		msg := tgbotapi.NewMessage(m.Chat.ID, texts.StartGame)
		bot.Send(msg)
	}
}
