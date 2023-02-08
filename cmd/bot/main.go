package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/pkg/telegram"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6015916633:AAETNXJFMPeWvbI2935QXYntnakUhgU3zP0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	b := telegram.NewBot(bot, "white0501")
	b.Start()
}
