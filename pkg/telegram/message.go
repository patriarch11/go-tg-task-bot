package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/dto"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
)

func subjectMessage(message *tgbotapi.Message, subject entity.Subject) tgbotapi.MessageConfig {
	msgText := fmt.Sprintf("Премет: %s\n\nОпис:\n%s", subject.Name, subject.Description)
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	msg.ReplyMarkup = inlineSubjectMarkup(subject)
	return msg
}

func inlineSubjectMarkup(subject entity.Subject) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			dto.SubjectToInlineKbRow(entity.GetTasks, "Показати Завдання", subject),
			dto.SubjectToInlineKbRow(entity.AddTask, "Додати завдання", subject),
		),
		tgbotapi.NewInlineKeyboardRow(
			dto.SubjectToInlineKbRow(entity.DeleteSubject, "Видалити", subject),
			dto.SubjectToInlineKbRow(entity.UpdateSubject, "Редагувати", subject),
		),
	)
}
