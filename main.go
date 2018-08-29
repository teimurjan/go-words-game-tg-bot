package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
	"github.com/teimurjan/go-words-game-tg-bot/handlers"
	"github.com/teimurjan/go-words-game-tg-bot/router"
)

func makeCommandRouter(bot *tgbotapi.BotAPI) *router.CommandRouter {
	routes := []*router.CommandRoute{
		&router.CommandRoute{
			"/start",
			handlers.NewWelcomeHandler(bot),
		},
		&router.CommandRoute{
			"/start_game",
			handlers.NewStartGameHandler(bot),
		},
	}
	return &router.CommandRouter{routes}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates, err := bot.GetUpdatesChan(u)

	commandRouter := makeCommandRouter(bot)

	for update := range updates {
		if update.Message != nil {
			continue
		}

		if update.Message.IsCommand() {
			commandRouter.GetHandler(update.Message.Text)(update.Message)
		} else if update.Message.Text != "" {
			// TODO: implement text routing
		}
	}
}
