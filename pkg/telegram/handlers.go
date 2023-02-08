package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart        = "start"
	commandAddSubject   = "add_subject"
	commandShowSubjects = "show"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandAddSubject:
		return b.handleAddSubjectCommand(message)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Упс! Я не знаю такої команди...")
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf("Привіт %s %s, тут ти можеш подивитись актуальні дз!",
			message.Chat.FirstName, message.Chat.LastName,
		))
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddSubjectCommand(message *tgbotapi.Message) error {
	if !b.isAdmin(message) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "permission denied")
		_, err := b.bot.Send(msg)
		return err
	}
	_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "додавання предмету"))
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}

func (b *Bot) isAdmin(message *tgbotapi.Message) bool {
	return b.adminUserName == message.Chat.UserName
}
