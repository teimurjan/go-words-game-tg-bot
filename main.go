package main

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
	"github.com/teimurjan/go-words-game/handlers"
	"github.com/teimurjan/go-words-game/router"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates, err := bot.GetUpdatesChan(u)

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
	router := router.CommandRouter{routes}

	for update := range updates {
		isMessageString := reflect.TypeOf(update.Message.Text).Kind() == reflect.String
		if !isMessageString || update.Message.Text == "" {
			continue
		}

		router.GetHandler(update.Message.Text)(&update)
	}
}
