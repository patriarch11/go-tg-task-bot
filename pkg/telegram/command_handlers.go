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
		if !b.isAdmin(message.Chat.UserName) {
			return b.notAdminResponse(message.Chat.ID)
		}
		return b.handleAddSubjectCommand(message)
	case commandShowSubjects:
		return b.handleShowSubjectsCommand(message)
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
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введіть назву предмету")
	b.state = WaitForSubjectName
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleShowSubjectsCommand(message *tgbotapi.Message) error {
	subjects, err := b.subjectRepository.GetAll()
	if err != nil {
		return err
	}
	for _, subject := range subjects {
		msg := subjectMessage(message, *subject)
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
