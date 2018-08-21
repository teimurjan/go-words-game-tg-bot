package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandRoute struct {
	Command string
	Handler func(update *tgbotapi.Update)
}

type Router interface {
	GetHandler(command string) func(update *tgbotapi.Update)
}

type CommandRouter struct {
	Routes []*CommandRoute
}

func (router *CommandRouter) GetHandler(command string) func(update *tgbotapi.Update) {
	for _, route := range router.Routes {
		if route.Command == command {
			return route.Handler
		}
	}
	return nil
}
