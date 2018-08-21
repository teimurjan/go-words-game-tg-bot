package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/teimurjan/go-words-game/texts"
)

func NewWelcomeHandler(bot *tgbotapi.BotAPI) func(u *tgbotapi.Update) {
	return func(u *tgbotapi.Update) {
		msg := tgbotapi.NewMessage(u.Message.Chat.ID, texts.WelcomeMessage)
		bot.Send(msg)
	}
}
