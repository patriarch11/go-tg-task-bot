package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch b.state {
	case WaitForSubjectName:
		return b.addNameSubject(message)
	case WaitForSubjectDescription:
		return b.addDescriptionSubject(message)
	default:
		return nil
	}
}
