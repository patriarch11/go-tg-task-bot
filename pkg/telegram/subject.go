package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
	"log"
)

func (b *Bot) startCreatingSubject(message *tgbotapi.Message) error {
	b.subject.Name = message.Text
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введіть опис предмету")
	b.state = WaitForSubjectDescription
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) startUpdateSubject(callbackQuery *tgbotapi.CallbackQuery, callback entity.Callback) error {
	b.subject.ID = callback.ID
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Введіть нове ім'я предмету")
	b.state = WaitForSubjectName
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) updateSubjectName(message *tgbotapi.Message) error {
	b.subject.Name = message.Text
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введіть новий опис предмету")
	b.state = WaitForNewSubjectDescription
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil

}

func (b *Bot) updateSubject(message *tgbotapi.Message) error {
	b.subject.Description = message.Text
	subject, err := b.subjectRepository.Update(&b.subject)
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

func (b *Bot) createSubject(message *tgbotapi.Message) error {
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

func (b *Bot) deleteSubject(callbackQuery *tgbotapi.CallbackQuery, callback entity.Callback) error {
	tasks, err := b.taskRepository.GetAllBySubjectID(callback.ID)
	if err != nil {
		return err
	}
	if err := b.subjectRepository.Delete(callback.ID); err != nil {
		return err
	}
	for _, task := range tasks {
		if err := b.taskRepository.Delete(task.ID); err != nil {
			return err
		}
	}
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Предмет видалено")
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
