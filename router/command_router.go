package router

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandRoute struct {
	Command string
	Handler func(m *tgbotapi.Message)
}

type CommandRouter struct {
	Routes []*CommandRoute
}

func (router *CommandRouter) GetHandler(command string) func(m *tgbotapi.Message) {
	for _, route := range router.Routes {
		if route.Command == command {
			return route.Handler
		}
	}
	return nil
}
