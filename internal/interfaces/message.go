package interfaces

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type MessageService interface {
	WrapMessageInMainKeyboard(msg tgbotapi.MessageConfig, username string) tgbotapi.MessageConfig
}
