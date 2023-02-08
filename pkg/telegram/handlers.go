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
	if !b.isAdmin(message) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "permission denied")
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
		return nil
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Введіть назву предмету")
	b.state = WaitForSubjectName
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleShowSubjectsCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "предмети")
	msg.ReplyMarkup = numericKeyboard
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

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

func (b *Bot) handleCallback(callbackQuery *tgbotapi.CallbackQuery) error {
	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := b.bot.Request(callback); err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, callbackQuery.Data)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) isAdmin(message *tgbotapi.Message) bool {
	return b.adminUserName == message.Chat.UserName
}
