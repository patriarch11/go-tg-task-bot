package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) addNameSubject(message *tgbotapi.Message) error {
	b.subject.Name = message.Text
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введіть опис предмету")
	b.state = WaitForSubjectDescription
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) addDescriptionSubject(message *tgbotapi.Message) error {
	b.subject.Description = message.Text
	subject, err := b.subjectRepository.Create(&b.subject)
	if err != nil {
		log.Fatal(err)
	}
	msg := subjectMessage(message, *subject)
	b.state = None
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
