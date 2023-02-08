package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot           *tgbotapi.BotAPI
	adminUserName string
}

func NewBot(bot *tgbotapi.BotAPI, userName string) *Bot {
	return &Bot{bot: bot, adminUserName: userName}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	b.handleUpdates(b.initUpdatesChannel())
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			_ = b.handleCallback(update.CallbackQuery)
			continue
		}
		if update.Message.IsCommand() {
			_ = b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	return updates
}
